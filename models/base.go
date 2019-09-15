package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB
func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbUri := "database.db"
	fmt.Println(dbUri)

	conn, err := gorm.Open("sqlite3", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.AutoMigrate(&Account{},&Score{},&RefreshToken{},&Destination{})
}

func GetDB() *gorm.DB {
	return db
}
