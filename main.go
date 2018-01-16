package main

import (
	"log"
	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/CityOfZion/neo-go-sdk/neo/models"
	"fmt"
)

//from config.json
type Configuration struct {
	AccountAddress string
	Host           string
}

func main() {
	nodeURI := "http://localhost:10332"
	//   nodeURI := "http://test1.cityofzion.io:8880"
	client := neo.NewClient(nodeURI)

	ok := client.Ping()
	if !ok {
		log.Fatal("Unable to connect to NEO node")
	}


	//1.Generate new address and send to client
	myAddress := GetNewAddress();
	SendNewAddress(myAddress)


	//bestBlockHash, err := client.GetBestBlockHash()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//currentBlock, err := client.GetBlockByHash(bestBlockHash)
	//if err != nil {
	//	log.Fatal(err)
	//}

	currentBlock, err := client.GetBlockByIndex(1820000)
	if err != nil {
		log.Fatal(err)
	}

	//trans, err := client.GetTransaction("0e9550edcf0a9675334a7dd3c32b621bbb4f5d9b2e65d42b3c33cd124752cf50")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//all transactions for current (last) block
	var transactions []models.Transaction = currentBlock.Transactions

	//all vouts where myAddress was detected
	var vouts [] models.Vout

	//detect address
	for _, element := range transactions {
		//fmt.Println(element)
		checkVouts(element, myAddress, &vouts)
	}

	for _, vout := range vouts {
		fmt.Println(vout)
	}
	//log.Printf("Transaction: %v", trans.Vout[0].Address)
	//log.Printf("currentBlock : %v", currentBlock)

}
func checkVouts(transaction models.Transaction, address string, vouts *[]models.Vout)  {
	for _, vout := range transaction.Vout {
		//fmt.Println(vout.Address)
		if vout.Address == address {
			*vouts = append(*vouts, vout)
		}
	}

}
func SendNewAddress(address string){
	//todo send pay address to Front
}

