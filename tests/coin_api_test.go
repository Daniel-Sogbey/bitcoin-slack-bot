package tests

import (
	"testing"

	coinapi "github.com/Daniel-Sogbey/slack-bot/coin_api"
)

func TestCoinApi(t *testing.T) {

	coinData, err := coinapi.GetCoinPrice()

	if err != nil {
		t.Errorf("Test failed with error %s", err.Error())
	}

	if coinData == nil {
		t.Errorf("Test failed, error getting coin data")
	}
}
