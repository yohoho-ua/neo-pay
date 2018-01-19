package main

import (
	"os"
	"encoding/json"
)

//from config.json
type Configuration struct {
	NodeURI string
	WaitTimeSec int
}

func NewConfiguraion() (*Configuration, error) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, err
}

