package main

import (
	"log"
	"os"
	"weather-bot/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error")
	}
	openWeatherToken := os.Getenv("OPEN_WEATHER_TOKEN")
	botToken := os.Getenv("BOT_TOKEN")
	bot.BotToken = botToken
	bot.OpenWeatherToken = openWeatherToken
	bot.Run()
}
