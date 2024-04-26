package models

type CoinData struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Price     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
}
