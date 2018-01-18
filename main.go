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
	"os"
)

//from config.json
type Configuration struct {
	NodeURI string
}


var (
	key = []byte("super-duper-secret-key-yohoho")
	store = sessions.NewCookieStore(key)
	CurrentCustomer Customer
)

//var customers [] Customer

func AddressHandler(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")

	// Check if user is authenticated
	if address, ok := session.Values["address"].(string); !ok || address=="" {
	CurrentCustomer= CreateCustomer(GetNewAddress)
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
	err:= http.ListenAndServe(":8080", router)
	if err != nil {
		log.Printf("Server x_x : %v", err)
	}
}


func InitConfig() (*Configuration, error) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, err
}



