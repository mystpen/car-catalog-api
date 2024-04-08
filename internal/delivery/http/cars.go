package http

import (
	"errors"
	"net/http"

	"github.com/mystpen/car-catalog-api/internal/model"
	"github.com/mystpen/car-catalog-api/internal/repository/postgresql"
	"github.com/mystpen/car-catalog-api/pkg/errorres"
	"github.com/mystpen/car-catalog-api/pkg/jsonutil"
	"github.com/mystpen/car-catalog-api/pkg/logger"
	"github.com/mystpen/car-catalog-api/pkg/validator"
)

type CarCatalogService interface {
	Get(id int64) (*model.CarInfo, error)
	GetAll(model.Filters) ([]*model.CarInfo, error)
	Update(cars *model.CarInfo) error
	Delete(id int64) error
	InsertRegNums([]string, *[]model.CarInfo) error
}

func (h *Handler) listCarsHandler(w http.ResponseWriter, r *http.Request) {
	var filters model.Filters
	qs := r.URL.Query()
	v := validator.New()

	filters.Mark = readString(qs, "mark", "")
	filters.Model = readString(qs, "model", "")
	filters.Year = readInt(qs, "year", 0, v)

	logger.PrintDebug("", map[string]any{
		"method": r.Method,
		"url": r.URL.String(),
		"filters": filters,
	})

	cars, err := h.service.GetAll(filters)

	logger.PrintDebug("", map[string]any{
		"url": r.URL.String(),
		"number of records": len(cars),
		"cars list": cars,
	})
	// Send a JSON response containing the car info.
	err = jsonutil.WriteJSON(w, http.StatusOK, jsonutil.Envelope{"cars": cars}, nil)
	if err != nil {
		errorres.ServerErrorResponse(w, r, err)
	}
}

func (h *Handler) addCarInfoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RegNums []string `json:"regNums"`
	}

	err := jsonutil.ReadJSON(w, r, &input)
	if err != nil {
		errorres.BadRequestResponse(w, r, err)
		return
	}
	logger.PrintDebug("", map[string]any{
		"method": r.Method,
		"url": r.URL.String(),
		"input": input.RegNums,
	})

	// validate
	var cars *[]model.CarInfo
	err = h.service.InsertRegNums(input.RegNums, cars)
	if err != nil {
		errorres.ServerErrorResponse(w, r, err)
		return
	}

	err = jsonutil.WriteJSON(w, http.StatusOK, jsonutil.Envelope{"cars": cars}, nil)
	if err != nil {
		errorres.ServerErrorResponse(w, r, err)
	}
}

func (h *Handler) updateCarInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		errorres.NotFoundResponse(w, r)
		return
	}

	// Fetch the existing car info from the database
	car, err := h.service.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, postgresql.ErrRecordNotFound):
			errorres.NotFoundResponse(w, r)
		default:
			errorres.ServerErrorResponse(w, r, err)
		}
		return
	}

	// Declare an input struct to hold the expected data from the client.
	var input struct {
		RegNum *string `json:"regNum"`
		Mark   *string `json:"mark"`
		Model  *string `json:"model"`
		Year   *int    `json:"year"`
	}

	logger.PrintDebug("", map[string]any{
		"method": r.Method,
		"url": r.URL.String(),
		"id": id,
		"input": input,
	})

	err = jsonutil.ReadJSON(w, r, &input)
	if err != nil {
		errorres.BadRequestResponse(w, r, err)
		return
	}
	// Copy the values from the request body
	if input.RegNum != nil {
		car.RegNum = *input.RegNum
	}
	if input.Year != nil {
		car.Year = *input.Year
	}
	if input.Model != nil {
		car.Model = *input.Model
	}
	if input.Mark != nil {
		car.Mark = *input.Mark
	}

	// validate

	err = h.service.Update(car)
	if err != nil{
		errorres.ServerErrorResponse(w, r, err)
		return
	}

	logger.PrintDebug("Updated", map[string]any{
		"car": car,
	})

	err = jsonutil.WriteJSON(w, http.StatusOK, jsonutil.Envelope{"car": car}, nil)
	if err != nil {
		errorres.ServerErrorResponse(w, r, err)
	}
}

func (h *Handler) deleteCarInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		errorres.NotFoundResponse(w, r)
		return
	}

	logger.PrintDebug("", map[string]any{
		"method": r.Method,
		"url": r.URL.String(),
		"id": id,
	})

	err = h.service.Delete(id)
	if err != nil{
		switch {
		case errors.Is(err, postgresql.ErrRecordNotFound):
			errorres.NotFoundResponse(w, r)
		default:
			errorres.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = jsonutil.WriteJSON(w, http.StatusOK, jsonutil.Envelope{"message": "car info successfully deleted"}, nil)
	if err != nil {
		errorres.ServerErrorResponse(w, r, err)
	}
}