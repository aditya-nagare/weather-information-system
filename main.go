package main

import (
	"log"
	"net/http"

	"github.com/aditya-nagare/weather-information-system/cmd/database"

	endpoints "github.com/aditya-nagare/weather-information-system/cmd/pkg/endpoints"
	handler "github.com/aditya-nagare/weather-information-system/cmd/pkg/handlers"
	repository "github.com/aditya-nagare/weather-information-system/cmd/pkg/repositories"
	service "github.com/aditya-nagare/weather-information-system/cmd/pkg/service"

	"github.com/gorilla/mux"
)

const (
	//HTTPPort is Port Number
	HTTPPort = "9001"
)

func main() {
	router := mux.NewRouter()

	db := database.NewDBConnection()

	weatherRepository := repository.NewWeatherRepositoryImpl(db)
	weatherService := service.NewWeatherServiceImpl(weatherRepository)
	weatherHandler := handler.NewWeatherHandlerImpl(weatherService)
	endpoints.NewRoute(router, weatherHandler)

	log.Println("Staring Server on http://localhost:" + HTTPPort)
	log.Fatal(http.ListenAndServe(":"+HTTPPort, router))
}
