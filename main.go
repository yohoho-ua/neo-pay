package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/CityOfZion/neo-go-sdk/neo"
	//"github.com/CityOfZion/neo-go-sdk/neo/models"
	//"fmt"
)




//var customers [] Customer
var CurrentCustomer Customer

func AddressHandler(w http.ResponseWriter, req *http.Request) {

	CurrentCustomer = CreateCustomer(GetNewAddress)
	//fmt.Println(CurrentCustomer)
	json.NewEncoder(w).Encode(CurrentCustomer)
	//w.Write(customer.AssignedAddress)
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {

	//26 is test price-value, later should be gotten from front
	CheckStatus(&CurrentCustomer, 26)
	json.NewEncoder(w).Encode(CurrentCustomer)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/address", AddressHandler).Methods("GET")
	r.HandleFunc("/status", StatusHandler).Methods("GET")
	return r
}

func main() {

	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
