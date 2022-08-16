package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int
	Name  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "Toyota", "Corolla", 100000},
	{2, "Kia", "Camry", 200000},
	{3, "Honda", "Civic", 500000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}
func returnCarsByBrand(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)  // make a map [make:Toyota]
	carM := vars["make"] // value
	//fmt.Println(vars, carM)
	cars := &[]Vehicle{}
	for _, car := range vehicles {
		if car.Name == carM {
			*cars = append(*cars, car)
		}
	}
	json.NewEncoder(w).Encode(cars)
}
func returnCarsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	carid, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert")
	}
	for _, car := range vehicles {
		if car.Id == carid {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
		}
	}

}
func updateCar(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	carid, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert")
	}
	var updateCar Vehicle
	json.NewDecoder(r.Body).Decode(&updateCar)
	for k, v := range vehicles {
		if v.Id == carid {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
			vehicles = append(vehicles, updateCar)
		}
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}
func createCar(w http.ResponseWriter, r *http.Request) {

	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar) //set r.body request data equal to newcar
	vehicles = append(vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}
func removeCarByIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unble to Convert")
	}
	for k, v := range vehicles {
		if v.Id == carId {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", returnCarsByBrand).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCarsById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", removeCarByIndex).Methods("DELETE")
	http.ListenAndServe(":8081", router)
}
