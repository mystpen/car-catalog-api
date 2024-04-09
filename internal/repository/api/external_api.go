package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mystpen/car-catalog-api/config"
	"github.com/mystpen/car-catalog-api/internal/model"
)
var (
	ErrBadRequest = errors.New("incorrect request")
	ErrNoResponce = errors.New("no responce from API")
)


type ApiClient struct {
	config *config.Config
}

func NewApiClient(config *config.Config) *ApiClient{
	return &ApiClient{config: config}
}

func (ac *ApiClient) GetCarInfo(regNum string) (*model.CarInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/info?regNum=%s", ac.config.ExternalAPIURL, regNum))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		if resp.StatusCode == http.StatusBadRequest{
			return nil, ErrBadRequest
		} else {
			return nil, ErrNoResponce
		}
	}

	var carInfo model.CarInfo
	if err := json.NewDecoder(resp.Body).Decode(&carInfo); err != nil {
		return nil, err
	}

	return &carInfo, nil
}
