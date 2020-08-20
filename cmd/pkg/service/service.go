package services

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/aditya-nagare/weather-information-system/cmd/models"
	repository "github.com/aditya-nagare/weather-information-system/cmd/pkg/repositories"
)

// WeatherService describes the service.
type WeatherService interface {
	FetchData(context.Context, *sync.WaitGroup) error
	GetData(context.Context, string) (*models.CityData, error)
}

//WeatherServiceImpl **
type WeatherServiceImpl struct {
	repo repository.WeatherRepository
}

//NewWeatherServiceImpl inject dependencies user repository
func NewWeatherServiceImpl(repo repository.WeatherRepository) WeatherService {
	return &WeatherServiceImpl{repo: repo}
}

const (
	//APIKEY for Weather API
	APIKEY = "fb49499e61f6086a2e22ba57030965bf"
)

//FetchData Function to fetch & process the Data from the Weather API
func (serviceImpl WeatherServiceImpl) FetchData(ctx context.Context, wg *sync.WaitGroup) error {

	wg.Add(10)

	var routine = len(Cities)
	citiesCh := make(chan string)

	for i := 0; i < routine; i++ {
		go func(cityCh <-chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				cityname, ok := <-cityCh
				if !ok {
					return
				}

				response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + cityname + "&appid=" + APIKEY)
				if err != nil {
					return
				}

				cityInfo := models.WeatherInfo{}

				decoder := json.NewDecoder(response.Body)
				err = decoder.Decode(&cityInfo)
				if err != nil {
					return
				}

				rawData, err := json.Marshal(cityInfo)
				if err != nil {
					return
				}

				city := models.CityData{}
				city.Name = cityname
				city.Lat = cityInfo.Coord.Lat
				city.Lon = cityInfo.Coord.Lon
				city.Temp = cityInfo.Main.Temp
				city.Pressure = cityInfo.Main.Pressure
				city.SeaLevel = cityInfo.Main.SeaLevel
				city.MoreInfo = rawData

				err = serviceImpl.repo.FetchData(ctx, city)
				if err != nil {
					return
				}
			}
		}(citiesCh, wg)
	}

	for _, city := range Cities {
		citiesCh <- city
	}

	wg.Wait()

	close(citiesCh)
	return nil
}

//GetData Handler Function to get the data for the city from DB
func (serviceImpl WeatherServiceImpl) GetData(ctx context.Context, cityName string) (resp *models.CityData, err error) {
	resp, err = serviceImpl.repo.GetData(ctx, cityName)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
