package database

import (
	"fmt"

	//Blank Import for MySQL Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	//DBUSER Database Username
	DBUSER = "root"

	//DBPASS Database Password
	DBPASS = "password"

	//DBNAME Database Name
	DBNAME = "weather-information-system"
)

//NewDBConnection for Database Connection
func NewDBConnection() *gorm.DB {

	connectionStr := fmt.Sprintf(DBUSER + ":" + DBPASS + "@/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("mysql", connectionStr)

	if err != nil {
		panic("Database Connection Error!")
	}
	return db
}
