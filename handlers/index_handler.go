package handlers

import (
	"github.com/gorilla/mux"
)

func IndexRouting() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cars", GetAllCars).Methods("GET")
	r.HandleFunc("/cars", AddCar).Methods("POST")
	r.HandleFunc("/cars/{registration}/rent", RentCar).Methods("POST")
	r.HandleFunc("/cars/{registration}/return", RentCar).Methods("POST")
	return r
}
