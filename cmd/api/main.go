package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/mystpen/car-catalog-api/config"
	"github.com/mystpen/car-catalog-api/internal/delivery/http"
	"github.com/mystpen/car-catalog-api/internal/repository/postgresql"
	"github.com/mystpen/car-catalog-api/internal/service"
	"github.com/mystpen/car-catalog-api/pkg/logger"
)

func main() {
	logger.New(os.Stdout, logger.LevelInfo)

	cfg, err := config.Load()
	if err != nil {
		logger.PrintError(err, nil)
	}

	// Connect to DB
	db, err := openDB(*cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()

	// Database migrations
	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", migrationDriver)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.PrintFatal(err, nil)
	}

	// prepare repo
	carRepo := postgresql.NewCarsRepository(db)
	peopleRepo := postgresql.NewPeopleRepository(db)

	// service layer
	carCatalogService := service.NewCarCatalogService(carRepo, peopleRepo)

	// handler
	handler := http.NewHandler(carCatalogService)

	srv := NewServer(
		handler,
		cfg)

	err = srv.Start()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
