package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"test-go/models"
	"test-go/repositories"
)

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllCars())
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&car)
	carDb, err := repositories.AddCar(car)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(carDb)
	}

}
func RentCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("***********", params["registration"])
	registration := params["registration"]
	w.Header().Set("Content-Type", "application-json")
	repositories.RentCar(registration)
	//	w.WriteHeader(http.StatusOK)
	//	json.NewEncoder(w).Encode(carDb)
}
func ReturnCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("***********", params["registration"])
	registration := params["registration"]
	w.Header().Set("Content-Type", "application-json")
	repositories.RentCar(registration)
	//	w.WriteHeader(http.StatusOK)
	//	json.NewEncoder(w).Encode(carDb)
}
