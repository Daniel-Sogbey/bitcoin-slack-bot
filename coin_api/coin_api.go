package coinapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Daniel-Sogbey/slack-bot/models"
)

func GetCoinPrice() (*models.CoinData, error) {
	client := http.Client{}
	coinUUID := "Qwsogvtv82FCd"

	url := fmt.Sprintf("https://api.coinranking.com/v2/coin/%s/price", coinUUID)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("x-access-token", os.Getenv("API_KEY"))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	var coinData models.CoinData

	err = json.Unmarshal(body, &coinData)

	if err != nil {
		return nil, err
	}

	return &coinData, nil
}
