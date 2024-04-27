package tests

import (
	"allergycron/utils"
	"testing"
)

func TestSendSlackMessage(t *testing.T) {
	err := utils.SendSlackMessage("Test Message!")

	if err != nil {
		t.Errorf("Error sending message to slack. Error : %v", err)
	}
}
