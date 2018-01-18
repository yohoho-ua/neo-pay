package main

import (
	"testing"
)


func TestInitConfig(t *testing.T) {
	configuratin := Configuration{NodeURI:"http://localhost:10332"}
	expectedURI := configuratin.NodeURI
	actualURI := InitConfig().NodeURI

	if actualURI != expectedURI {
		t.Errorf("initConfig returned unexpected confiuration: got %v want %v", actualURI, expectedURI)
	}

}