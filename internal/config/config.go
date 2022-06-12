package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Token  string `json: "token"`
	Prefix string `json: "prefix"`
}

func ParseConfigFromJSONFile(fileName string) (c *Config, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	c = new(Config)
	json.NewDecoder(f).Decode(c)
	log.Println(c)
	return c, nil
}
