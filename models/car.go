package models

type Car struct {
	Model        string  `json:"model" validate:"required" `
	Registration string  `json:"registration" validate:"required"`
	Mileage      float32 `json:"mileage" validate:"required"`
	IsAvailable  bool    `json:"isAvailable" `
}
