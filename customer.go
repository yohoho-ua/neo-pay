package main

type Customer struct {
	AssignedAddress string
	balance         int64
	startBlock      int64
	statusPaid      bool
}

func CreateCustomer() Customer {
	newAddress := getNewAddress();
	return Customer{AssignedAddress: newAddress, balance: 0, statusPaid:false}

}
//todo: Create new Address via rpc call
func getNewAddress() string {
	return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"
}