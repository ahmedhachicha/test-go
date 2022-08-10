package repositories

import (
	"fmt"
	"test-go/config"
	"test-go/models"
)

func GetAllCars() []models.Car {
	var cars []models.Car
	config.DB.Find(&cars)
	return cars
}

func AddCar(car models.Car) (models.Car, error) {
	err := config.DB.Create(&car).Error
	return car, err
}

func FindByRegistration(registration string) models.Car {
	var car models.Car
	config.DB.Where("registration = ?", registration).First(&car)
	return car
}

func RentCar(registration string) models.Car {
	var car models.Car
	x := config.DB.Where("registration = ?", registration).First(&car)
	fmt.Println(x)
	if car.IsRented == true {
		fmt.Println("car already rented")
	}
	if car.IsRented == false {
		config.DB.Model(car).Where("registration = ?", registration).Update("is_rented", true)
	}
	return car
}

func ReturnCar(registration string) models.Car {
	var car models.Car
	x := config.DB.Where("registration = ?", registration).First(&car)
	fmt.Println(x)
	if car.IsRented == true {
		config.DB.Model(car).Where("registration = ?", registration).Update("is_rented", true)
	}
	if car.IsRented == false {
		config.DB.Model(car).Where("registration = ?", registration).Update("is_rented", true)
	}
	return car
}
