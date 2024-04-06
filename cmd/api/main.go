package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/mystpen/car-catalog-api/config"
	"github.com/mystpen/car-catalog-api/internal/delivery/http"
	"github.com/mystpen/car-catalog-api/internal/repository/postgresql"
	"github.com/mystpen/car-catalog-api/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		// logger.ErrLog.Fatal(err)
	}

	// Connect to DB
	db, err := openDB(*cfg)
	if err != nil {
		// logger.ErrLog.Fatal(err)
	}
	defer db.Close()

	// Database migrations
	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		// logger.ErrLog.Fatal(err)
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", migrationDriver)
	if err != nil {
		// logger.ErrLog.Fatal(err)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		// logger.ErrLog.Fatal(err)
	}

	// prepare repo
	carRepo := postgresql.NewCarsRepository(db)
	peopleRepo := postgresql.NewPeopleRepository(db)

	// service layer
	carCatalogService := service.NewCarCatalogService(carRepo, peopleRepo)

	// handler
	http.NewHandler(carCatalogService)

	// srv := httpserver.New(handler)
}
