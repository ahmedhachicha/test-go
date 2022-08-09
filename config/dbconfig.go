package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func DbConfiguration() {
	err := godotenv.Load(".env")

	var x = fmt.Sprintf("%s:@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	DB, err = gorm.Open(mysql.Open(x), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database")

}
