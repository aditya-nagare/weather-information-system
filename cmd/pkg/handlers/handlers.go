package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/aditya-nagare/weather-information-system/cmd/models"

	service "github.com/aditya-nagare/weather-information-system/cmd/pkg/service"

	"github.com/gorilla/mux"
)

//WeatherHandlersImpl for handler Functions
type WeatherHandlersImpl struct {
	svc service.WeatherService
}

//NewWeatherHandlerImpl inits dependencies for Handlers
func NewWeatherHandlerImpl(service service.WeatherService) *WeatherHandlersImpl {
	return &WeatherHandlersImpl{svc: service}
}

//FetchData Handler Function to fetch the Data from the Weather API
func (handlersImpl WeatherHandlersImpl) FetchData(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var wg sync.WaitGroup
	go handlersImpl.svc.FetchData(ctx, &wg)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var resp models.FetchResponse
	resp.Message = "Fetching Weather Data in background."

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
}

//GetData Handler Function to get the data for the city from DB
func (handlersImpl WeatherHandlersImpl) GetData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := req.Context()

	vars := mux.Vars(req)
	cityName := vars["cityName"]
	resp, err := handlersImpl.svc.GetData(ctx, cityName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
