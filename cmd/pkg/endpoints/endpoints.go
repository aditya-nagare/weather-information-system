package endpoints

import (
	handler "github.com/aditya-nagare/weather-information-system/cmd/pkg/handlers"

	"github.com/gorilla/mux"
)

//NewRoute Creates API Endpoints
func NewRoute(router *mux.Router, handler *handler.WeatherHandlersImpl) {

	router.HandleFunc("/weather", handler.FetchData).Methods("POST")

	router.HandleFunc("/weather/{cityName}", handler.GetData).Methods("GET")

}
