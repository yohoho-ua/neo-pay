package main

import (
	"log"
	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/CityOfZion/neo-go-sdk/neo/models"
	"fmt"
	"os"
	"encoding/json"
	"strconv"
)

const (
	assetTypeNEO = "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
	assetTypeGAS = "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7"

	//amount of blocks that will be tracked after customer has got payment address. Max wait time, default - 24 hours (or 5760 blocks, 15 sec/block)
	//maxAddressLife = (60 * 60  * 24) / 15
	maxAddressLife = (60 *10) / 15 //for test
)



//from config.json
type Configuration struct {
	NodeURI string
}

func Check(customer *Customer) bool{
	configuration := initConfig()
	client := neo.NewClient(configuration.NodeURI)
	ok := client.Ping()
	if !ok {
		log.Fatal("Unable to connect to NEO node")
	}

	//get current block hash to start tracking
	bestBlockHash, err := client.GetBestBlockHash()
	if err != nil {
		log.Fatal(err)
	}

	//get corresponding block object
	currentBlock, err := client.GetBlockByHash(bestBlockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentBlock.Index)
	//currentIndex := currentBlock.Index

	//for test
	currentIndex := int64(1817210)


	//for test
	price := int64(26)

	customer.startBlock=currentIndex

	//fmt.Println(customer)

	trackPayment(client, currentIndex, customer, price)

	fmt.Println(customer)
	return customer.statusPaid
}
func trackPayment(client neo.Client, currentIndex int64, customer *Customer,  price int64) {
	for {
		if (currentIndex - customer.startBlock)>= maxAddressLife  {
			fmt.Println("Payment not found")
			break
		}
		currentBlock, err := client.GetBlockByIndex(currentIndex)
		if err != nil {
			log.Fatal(err)
		}
		//all transactions for current (last) block
		var transactions []models.Transaction = currentBlock.Transactions


		//detect address
		for _, element := range transactions {
			//fmt.Println(element)
			checkVouts(element, customer.AssignedAddress, customer)
		}
		fmt.Printf("current index : %v, balance: %v \n", currentIndex, customer.balance)

		if customer.balance >= price {
			fmt.Println("sucsess")
			customer.statusPaid = true
		} else {
			currentIndex++
		}
	}
}

func checkVouts(transaction models.Transaction, address string, customer *Customer) {
	for _, vout := range transaction.Vout {
		if vout.Address == address && vout.Asset == assetTypeNEO {
			paidAmount, err := strconv.ParseInt(vout.Value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			customer.balance = customer.balance + paidAmount
		}
	}

}

func initConfig() *Configuration {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &configuration
}



