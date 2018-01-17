package main

import (
	"testing"
)



func TestGetNewAddress(t *testing.T) {
	expectedLength := len("AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4")

	actualLength := len(GetNewAddress())

	if actualLength != expectedLength {
		t.Errorf("GetNewAddress returned unexpected NEO Address: got %v want %v", actualLength, expectedLength)
	}
}


func TestCreateCustomer(t *testing.T) {

	var expectedBalance int64 = 0
	var expectedStartBlock int64 = 0
	expectedStatusPaid := false
	expectedAddress := "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"


	actualCustomer := CreateCustomer(mock_GetNewAddress)//mocked


	if actualCustomer.balance != expectedBalance {
		t.Errorf("CreateCustomer returned unexpected customer object balance : got %v want %v", actualCustomer.balance, expectedBalance)
	}
	if actualCustomer.startBlock != expectedBalance {
		t.Errorf("CreateCustomer returned unexpected customer object startBlock : got %v want %v", actualCustomer.balance, expectedStartBlock)
	}
	if actualCustomer.statusPaid != expectedStatusPaid {
		t.Errorf("CreateCustomer returned unexpected customer object statusPaid : got %v want %v", actualCustomer.statusPaid, expectedStatusPaid)
	}

	if actualCustomer.AssignedAddress != expectedAddress {
		t.Errorf("CreateCustomer returned unexpected customer object AssignedAddress : got %v want %v", actualCustomer.AssignedAddress, expectedAddress)
	}
}

func mock_GetNewAddress() string {
	return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"
}