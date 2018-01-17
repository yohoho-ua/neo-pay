package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"fmt"
)




var customers [] Customer

func DepositHandler(w http.ResponseWriter, req *http.Request) {

	customer := CreateCustomer()
	Check(customer)
	w.Write(customer.AssignedAddress)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/neo", DepositHandler).Methods("GET")
	return r
}

func main() {
	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
