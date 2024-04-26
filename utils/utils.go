package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendSlackMessage(message string) error {
	client := http.Client{}

	js, err := json.Marshal(map[string]string{"text": fmt.Sprintf("Current price of BTC : $%s", message)})

	if err != nil {
		return err
	}

	body := bytes.NewBuffer(js)

	req, err := http.NewRequest(http.MethodPost, os.Getenv("SLACK_URL"), body)

	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(respBody))

	return nil
}
