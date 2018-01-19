package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	//"encoding/gob"
	//"github.com/CityOfZion/neo-go-sdk/neo"
	//"github.com/CityOfZion/neo-go-sdk/neo/models"
	//"fmt"
	"github.com/gorilla/sessions"
	//"os"
)



var (
	key = []byte("super-duper-secret-key-yohoho")
	store = sessions.NewCookieStore(key)
	CurrentCustomer Customer
)

//var customers [] Customer

func AddressHandler(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")

	// Check if user is authenticated
	if address, ok := session.Values["address"].(string); !ok || address == "" {
		configuration, err :=NewConfiguraion()
		if err != nil {
			log.Fatal(err)
		}
		CurrentCustomer = CreateCustomer(configuration)
		session.Values["address"] = CurrentCustomer.AssignedAddress
		session.Save(req, w)
		//http.Error(w, "Forbidden", http.StatusForbidden)
		//return
	}

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

	//gob.Register(Customer{})
	router := newRouter()
	//log.Fatal(http.ListenAndServe(":8080", router))
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("Server x_x : %v", err)
	}
}



