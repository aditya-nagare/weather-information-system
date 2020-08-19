package repositories

import (
	"context"

	"github.com/aditya-nagare/weather-information-system/cmd/models"

	"github.com/jinzhu/gorm"
)

//WeatherRepository implements all methods in WeatherRepository
type WeatherRepository interface {
	FetchData(context.Context, models.CityData) error
	GetData(context.Context, string) (*models.CityData, error)
}

//WeatherRepositoryImpl **
type WeatherRepositoryImpl struct {
	db *gorm.DB
}

//NewWeatherRepositoryImpl dependencies of Database
func NewWeatherRepositoryImpl(db *gorm.DB) WeatherRepository {
	return &WeatherRepositoryImpl{db: db}
}

//FetchData Function to fetch & store the Data from the Weather API
func (repositoryImpl WeatherRepositoryImpl) FetchData(ctx context.Context, city models.CityData) (err error) {
	db := repositoryImpl.db
	err = db.Table("city_data").Create(&city).Error
	if err != nil {
		return
	}
	return nil
}

//GetData Handler Function to get the data for the city from DB
func (repositoryImpl WeatherRepositoryImpl) GetData(ctx context.Context, cityName string) (*models.CityData, error) {
	db := repositoryImpl.db
	city := models.CityData{}
	err := db.Where("name=?", cityName).First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}
