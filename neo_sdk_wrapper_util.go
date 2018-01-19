package main

import (
	"log"
	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/CityOfZion/neo-go-sdk/neo/models"
	//"fmt"
	//"os"
	//"encoding/json"
	//"strconv"
)



var Client neo.Client

func GetCurrentBlockIndex(configuration *Configuration) int64 {
	if Client.NodeURI == "" {
		initClient(configuration)
	}
	currentBlockIndex, err := Client.GetBlockCount()
	if err != nil {
		log.Println(err)
		return -1
	}
	return currentBlockIndex
}

func GetBlockByIndex(index int64) *models.Block {
	block, err := Client.GetBlockByIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}
	return block
}
func initClient(configuration *Configuration)  {
	Client = neo.NewClient(configuration.NodeURI)
	ok := Client.Ping()
	if !ok {
		log.Fatal("Unable to connect to NEO node")
	}
}