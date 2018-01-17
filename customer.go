package main

type Customer struct {
	AssignedAddress string 	`json:"address,omitempty"`
	balance         int64  		`json:"balance,omitempty"`
	startBlock      int64
	statusPaid      bool 		`json:"status,omitempty"`
}

//for better testing and mocking
type NewAddressGetter func() string

func CreateCustomer(addressGetter NewAddressGetter) Customer {
	newAddress := addressGetter();
	return Customer{AssignedAddress: newAddress, balance: 0, statusPaid:false}

}
//todo: Create new Address via rpc call
func GetNewAddress() string {
	return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"
}