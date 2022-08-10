package models

type Car struct {
	CarModel     string  `json:"model" validate:"required" `
	Registration string  `gorm:"not null;unique"`
	Mileage      float32 `json:"mileage" validate:"required"`
	IsRented     bool    `json:"is_rented" `
}
