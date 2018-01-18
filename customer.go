package main

type Customer struct {
	AssignedAddress string 	`json:"address,omitempty"`
	Balance         int64  		`json:"balance,omitempty"`
	StartBlock      int64
	StatusPaid      bool 		`json:"status,omitempty"`
}

//for better testing and mocking
type NewAddressGetter func() string

func CreateCustomer(addressGetter NewAddressGetter) Customer {
	_assignedAddress := addressGetter();
	_startBlock := GetCurrentBlockIndex()
	return Customer{AssignedAddress: _assignedAddress, Balance: 0, StartBlock: _startBlock, StatusPaid:false}

}
//todo: Create new Address via rpc call
func GetNewAddress() string {
	//return "AeQeWwHki197HDhaZJFKLeUN5tzi32gyZr"
	return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"
}