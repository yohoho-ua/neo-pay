package main

import (
	"testing"
)


func TestInitConfig(t *testing.T) {
	configuratin := Configuration{NodeURI:"http://localhost:10332"}
	expectedURI := configuratin.NodeURI
	actualConfiguration, err := InitConfig()

	if err != nil {
		t.Errorf("initConfig returned error confiuration: got %v want %v", err, actualConfiguration)
	}

	actualURI := actualConfiguration.NodeURI
	if actualURI != expectedURI {
		t.Errorf("initConfig returned unexpected confiuration: got %v want %v", actualURI, expectedURI)
	}

}