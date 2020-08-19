package models

import (
	"encoding/json"
	"time"
)

//WeatherInfo Struct with all the Weather Data
type WeatherInfo struct {
	Coord    Coord     `json:"coord"`
	Weather  []Weather `json:"weather"`
	Base     string    `json:"base"`
	Main     Main      `json:"main"`
	Wind     Wind      `json:"wind"`
	Clouds   Clouds    `json:"clouds"`
	Dt       int       `json:"dt"`
	Sys      Sys       `json:"sys"`
	Timezone int       `json:"timezone"`
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Cod      int       `json:"cod"`
}

//Coord Struct
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

//Weather Struct
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

//Main Struct
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

//Wind Struct
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

//Clouds Struct
type Clouds struct {
	All int `json:"all"`
}

//Sys Struct
type Sys struct {
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

//CityData Struct with all the Data for Specific City
type CityData struct {
	ID        uint            `gorm:"primary_key;column:id" json:"city_id"`
	Name      string          `gorm:"column:name" json:"city_name"`
	Lon       float64         `gorm:"column:lon" json:"lon"`
	Lat       float64         `gorm:"column:lat" json:"lat"`
	Temp      float64         `gorm:"column:temp" json:"average_temp"`
	Pressure  int             `gorm:"column:pressure" json:"average_ground_level_pressure"`
	SeaLevel  int             `gorm:"column:sea_level" json:"average_sea_level_pressure"`
	MoreInfo  json.RawMessage `gorm:"column:more_info" json:"raw_weather_data"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *time.Time      `sql:"index" json:"-"`
}

//FetchResponse Struct for response for fetch
type FetchResponse struct {
	Message string `json:"message"`
}
