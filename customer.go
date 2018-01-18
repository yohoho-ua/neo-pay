package main

import (
	"github.com/labstack/gommon/log"
	"net/url"
	"net/http"
	"encoding/json"
	//"fmt"
	"io/ioutil"
)

type Customer struct {
	AssignedAddress string        `json:"address,omitempty"`
	Balance         int64                `json:"balance,omitempty"`
	StartBlock      int64
	StatusPaid      bool                `json:"status,omitempty"`
}

//for better testing and mocking
type NewAddressGetter func() string

func CreateCustomer(addressGetter NewAddressGetter) Customer {
	_assignedAddress := addressGetter();
	_startBlock := GetCurrentBlockIndex()
	return Customer{AssignedAddress: _assignedAddress, Balance: 0, StartBlock: _startBlock, StatusPaid:false}

}

func GetNewAddress() string {
	//return "AeQeWwHki197HDhaZJFKLeUN5tzi32gyZr"
	//return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"

	//build url
	configuration := InitConfig()
	u, err := url.Parse(configuration.NodeURI)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("jsonrpc", "2")
	q.Set("method", "getnewaddress")
	q.Set("id", "1")
	q.Set("params", "[]")
	u.RawQuery = q.Encode()
	responseBlob, _ := http.Get(u.String())

	buf, _ := ioutil.ReadAll(responseBlob.Body)


	type Response struct {
		Result string `json:"result"`
	}

	var response Response

	error := json.Unmarshal(buf, &response)
	if error != nil {
		log.Fatal(error)
	}

	return response.Result
}

