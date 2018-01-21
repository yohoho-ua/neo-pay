package main

import (
	"log"
	"github.com/CityOfZion/neo-go-sdk/neo/models"
	"fmt"
	"strconv"
	//"time"
)

const (
	assetTypeNEO = "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
	assetTypeGAS = "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7"

	//amount of blocks that will be tracked after customer has got payment address. Max wait time, default - 24 hours (or 5760 blocks, 15 sec/block)
	//maxAddressLife = (60 * 60  * 24) / 15
	maxAddressLife = (60 * 60 * 24 * 10) //for test
)

func CheckStatus(customer *Customer, configuration *Configuration) {

	//customer.StartBlock = 1817208 //for test

	currentIndex := customer.StartBlock
	var transactions []models.Transaction
	fmt.Println(customer)
	for {

		////check if customer is still valid
		////or if payment was not made for too long (constant maxAddressLife)
		//if !isAddressStillValid(currentIndex, customer.StartBlock) || customer.StartBlock ==-1 {
		//	fmt.Printf("Payment for address: %v was not found\n", customer.AssignedAddress)
		//	return
		//}

		//wait for new blocks (default 180 sec)
		if currentIndex >= GetCurrentBlockIndex(configuration) {
			fmt.Printf("waiting for new blocks\n")
			//time.Sleep(time.Duration(TestConfiguration.WaitTimeSec) * time.Second)
			return
		}

		currentBlock := GetBlockByIndex(currentIndex)

		//all transactions for current (last) block
		transactions = currentBlock.Transactions


		//detect payment address in vouts of current transaction
		for _, element := range transactions {
			checkCurrentBlockTransactions(element, customer)
		}
		fmt.Printf("current index : %v, balance: %v \n", currentIndex, customer.Deposit)

		//check customer Balance enough
		//if customer.Deposit >= price {
		//	fmt.Println("sucsess")
		//	customer.StatusPaid = true
		//	return
		//}
		currentIndex++
		customer.StartBlock = currentIndex
	}
}
func isAddressStillValid(currentBlockIndex int64, startBlockIndex int64) bool {
	return (currentBlockIndex - startBlockIndex) < maxAddressLife
}

func checkCurrentBlockTransactions(transaction models.Transaction, customer *Customer) {
	for _, vout := range transaction.Vout {
		if vout.Address == customer.AssignedAddress && vout.Asset == assetTypeNEO {
			paidAmount, err := strconv.ParseInt(vout.Value, 10, 64)
			if err != nil {
				log.Println("vout.Value parse error, NaN?")
				return
			}
			//if payment was found - increase customer Balance
			customer.Deposit = customer.Deposit + paidAmount
		}
	}

}





