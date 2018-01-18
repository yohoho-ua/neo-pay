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

//from config.json
type Configuration struct {
	NodeURI string
}

var client neo.Client

func GetCurrentBlockIndex() int64 {
	if client.NodeURI == "" {
		initClient()
	}
	currentBlockIndex, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	return currentBlockIndex
}

func GetBlockByIndex(index int64) *models.Block {
	block, err := client.GetBlockByIndex(index)
	if err != nil {
		log.Fatal(err)
	}
	return block
}
func initClient() {
	configuration := initConfig()
	client = neo.NewClient(configuration.NodeURI)
	ok := client.Ping()
	if !ok {
		log.Fatal("Unable to connect to NEO node")
	}
}