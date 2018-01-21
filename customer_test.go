package main

import (
	"testing"
	//"fmt"
)
var TestConfiguration = Configuration{NodeURI:"http://localhost:10332"}

func TestGetNewAddress(t *testing.T) {
	expectedLength := len("AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4")

	actualAddress, err := GetNewAddress(&TestConfiguration)
	if err != nil {
		t.Errorf("GetNewAddress returned error: %v", err)
	}
	actualLength := len(actualAddress)
	if actualLength != expectedLength {
		t.Errorf("GetNewAddress returned unexpected NEO Address: got %v want %v", actualLength, expectedLength)
	}
}

func TestCreateCustomer(t *testing.T) {

	var expectedBalance int64 = 0
	expectedStartBlock := GetCurrentBlockIndex(&TestConfiguration);
	expectedStatusPaid := false
	//expectedAddress := "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"

	actualCustomer := CreateCustomer(&Configuration{NodeURI:"http://localhost:10332"})


	if actualCustomer.Deposit != expectedBalance {
		t.Errorf("CreateCustomer returned unexpected customer object balance : got %v want %v", actualCustomer.Deposit, expectedBalance)
	}
	if actualCustomer.StartBlock != expectedStartBlock {
		t.Errorf("CreateCustomer returned unexpected customer object startBlock : got %+v want %+v", actualCustomer.StartBlock, expectedStartBlock)
	}
	if actualCustomer.StatusPaid != expectedStatusPaid {
		t.Errorf("CreateCustomer returned unexpected customer object statusPaid : got %v want %v", actualCustomer.StatusPaid, expectedStatusPaid)
	}

	//if actualCustomer.AssignedAddress != expectedAddress {
	//	t.Errorf("CreateCustomer returned unexpected customer object AssignedAddress : got %v want %v", actualCustomer.AssignedAddress, expectedAddress)
	//}
}

//func mock_GetNewAddress() (string, error) {
//	return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4", nil
//}