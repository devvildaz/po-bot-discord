package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/pomodoro-discgo/internal/config"
)

func main() {
	log.Println("++++ DISCORD BOT ++++")

	const configFile = "./config/config.json"

	cfg, err := config.ParseConfigFromJSONFile(configFile)

	log.Println(cfg)

	if err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	if err = discord.Open(); err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
