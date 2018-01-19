package main

import (
	"github.com/labstack/gommon/log"
	"net/url"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Customer struct {
	AssignedAddress string        `json:"address,omitempty"`
	Balance         int64                `json:"balance,omitempty"`
	StartBlock      int64
	StatusPaid      bool                `json:"status,omitempty"`
}

//for better testing and mocking
type NewAddressGetter func(Configuration) (string, error)

func CreateCustomer(configuration *Configuration) Customer {
	_assignedAddress, err := GetNewAddress(configuration)
	if err != nil {
		fmt.Printf("Customer will be returned without address. %v\n", err)
		_assignedAddress = ""
	}
	_startBlock := GetCurrentBlockIndex(configuration)
	return Customer{AssignedAddress: _assignedAddress, Balance: 0, StartBlock: _startBlock, StatusPaid:false}

}

func GetNewAddress(configuration *Configuration) (string, error) {
	//return "AeQeWwHki197HDhaZJFKLeUN5tzi32gyZr"
	//return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"

	//build url
	u, err := url.Parse(configuration.NodeURI)
	if err != nil {
		log.Printf("Node URI not parsed. %v\n", err)
		return "", err
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
		Error map[string]string `json:"error"`
	}

	var response Response

	//check is wallet open and decrypted
	mapConfig := make(map[string]interface{})
	json.Unmarshal(buf, &mapConfig)
	if mapConfig["error"] != nil {
	fmt.Printf("%+v\n", mapConfig["error"])
	fmt.Println("Wallet is closed, open please")
	return "", nil
	}

	err = json.Unmarshal(buf, &response)
	if err != nil {
		return "", err
	}
	return response.Result, nil
}

