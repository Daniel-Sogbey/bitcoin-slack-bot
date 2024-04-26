package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	coinapi "github.com/Daniel-Sogbey/slack-bot/coin_api"
	"github.com/Daniel-Sogbey/slack-bot/models"
	"github.com/Daniel-Sogbey/slack-bot/utils"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	client := redisClient()
	defer client.Close()

	fmt.Println("Redis Client : ", client)

	coinData, err := coinapi.GetCoinPrice()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Coin Data : ", coinData)

	coinDataJSON, err := json.Marshal(coinData)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	err = client.Set(ctx, "coin", coinDataJSON, time.Minute).Err()
	if err != nil {
		log.Fatal(err)
	}

	oldCoinData, err := getPreviousCoinPrice(ctx, client)

	fmt.Println("Old Coin Data : ", oldCoinData)

	if err != nil {
		if err == redis.Nil {
			log.Println("No previous coin data found in Redis")
		} else {
			log.Fatal(err)
		}
	}

	loc, err := time.LoadLocation("Africa/Accra")

	cronJob := cron.NewWithLocation(loc)

	log.Println("TIME NOW : ", time.Now())

	//seconds,minutes, hours, day of month, month, day of week
	cronJob.AddFunc("*/5 * * * *", func() {
		log.Println("Cron job running...")
		if oldCoinData.Data.Price > coinData.Data.Price {
			fmt.Println("HERE")
			utils.SendSlackMessage(coinData.Data.Price)
		}

		log.Println("OLD COIN DATA ", oldCoinData)
	})

	if err != nil {
		log.Fatal(err)
	}

	cronJob.Start()

	select {}
}

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func getPreviousCoinPrice(ctx context.Context, rdb *redis.Client) (*models.CoinData, error) {

	coinDataJson, err := rdb.Get(ctx, "coin").Result()

	if err != nil {
		return nil, err
	}

	var coinData models.CoinData

	err = json.Unmarshal([]byte(coinDataJson), &coinData)

	if err != nil {
		return nil, err
	}

	return &coinData, nil

}
