package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"fmt"
)

var (
	CurrentCustomer Customer
	configuration   = NewConfiguraion()
	key             = []byte("super-secret-key-yohoho.ua")
	store           = sessions.NewCookieStore(key)
)

//var customers [] Customer

func NewAddressHandler(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "neo-pay-cookie")

	CurrentCustomer = CreateCustomer(configuration)

	session.Values["address"] = CurrentCustomer.AssignedAddress
	session.Save(req, w)

	fmt.Println(CurrentCustomer)
	json.NewEncoder(w).Encode(CurrentCustomer)
	//w.Write(customer.AssignedAddress)
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {

	//if user wants to generate new address
	paramNew := req.URL.Query().Get("new")
	if paramNew == "true" {
		//invalidate current customer/address
		CurrentCustomer.StartBlock = -1
	}

	//if current customer is empty or invalid
	if CurrentCustomer.AssignedAddress == "" || CurrentCustomer.StartBlock ==-1 {
		CurrentCustomer = CreateCustomer(configuration)
	}

	if CurrentCustomer.StartBlock < GetCurrentBlockIndex(configuration) {
		CheckStatus(&CurrentCustomer, configuration)
	}
	json.NewEncoder(w).Encode(CurrentCustomer)

}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/address", NewAddressHandler).Methods("GET")
	r.HandleFunc("/status", StatusHandler).Methods("GET")
	return r
}

func main() {
	router := newRouter()
	//log.Fatal(http.ListenAndServe(":8080", router))
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("Server x_x : %v", err)
	}
}
