package main

import (
	"os"
	"encoding/json"
	"fmt"
)

//from config.json
type Configuration struct {
	NodeURI string
	WaitTimeSec int
}

func NewConfiguraion() (*Configuration) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("something wrong with TestConfiguration ", err)
		os.Exit(3)
	}
	return &configuration
}

