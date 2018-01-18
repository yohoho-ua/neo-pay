package main

import (
	"testing"
	"github.com/CityOfZion/neo-go-sdk/neo/models"
)

func TestCheckVouts(t *testing.T) {

	customer := Customer{AssignedAddress: "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4", StartBlock: 182117, StatusPaid: false, Balance: 0}

	var vouts [] models.Vout

	//add 10
	voutGoodWithMoney := models.Vout{Address:"AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4",
		Value:"10",
		Asset: "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"}

	//add 3
	voutGoodWithMoneySecond := models.Vout{Address:"AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4",
		Value:"3",
		Asset: "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"}

	voutGoodWithOUTMoney := models.Vout{Address:"AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4",
		Value:"0",
		Asset: "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"}

	voutWrongAddress := models.Vout{Address:"WRONGNbdFMdYLBronyM3cHBzi49WKEwJWD4",
		Value:"100",
		Asset: "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"}

	voutWrongAssetType := models.Vout{Address:"AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4",
		Value:"1000",
		Asset: "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7"}

	vouts = append(vouts, voutGoodWithMoney, voutGoodWithMoneySecond, voutGoodWithOUTMoney, voutWrongAddress, voutWrongAssetType)

	transaction := models.Transaction{Vout:vouts}

	checkVouts(transaction, &customer)
	expectedBalance := int64(13)

	actualBalance := customer.Balance

	if actualBalance != expectedBalance {
		t.Errorf("checkVouts returned unexpected customer balance: got %v want %v", actualBalance, expectedBalance)
	}

}


func TestInitConfig(t *testing.T) {
	configuratin := Configuration{NodeURI:"http://localhost:10332"}
	expectedURI := configuratin.NodeURI
	actualURI := initConfig().NodeURI

	if actualURI != expectedURI {
		t.Errorf("initConfig returned unexpected confiuration: got %v want %v", actualURI, expectedURI)
	}

}

func TestIsAddressStillValid(t *testing.T) {
	expectedTrue := true
	actualTrue := isAddressStillValid(100, 50)

	expectedFalse := false
	actualFalse := isAddressStillValid(1000000000000, 50)


	if actualTrue != expectedTrue {
		t.Errorf("isAddressStillValid returned unexpected bool value: got %v want %v", actualTrue, expectedTrue)
	}

	if actualFalse != expectedFalse {
		t.Errorf("isAddressStillValid returned unexpected bool value: got %v want %v", actualTrue, expectedFalse)
	}

}



