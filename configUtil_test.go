package main

import (
	"testing"
)


func TestNewConfiguraion(t *testing.T) {
	configuratin := Configuration{NodeURI:"http://localhost:10332"}
	expectedURI := configuratin.NodeURI
	actualConfiguration := NewConfiguraion()

	actualURI := actualConfiguration.NodeURI
	if actualURI != expectedURI {
		t.Errorf("initConfig returned unexpected confiuration: got %v want %v", actualURI, expectedURI)
	}

}
