package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

/** CONFIG MODEL **/

type Config struct {
	Local EnvConfig
	Tests EnvConfig
}

type EnvConfig struct {
	Driver   string
	Host     string
	User     string
	Password string
	DbName   string
	DbPort   string
	ApiPort  string
}

func GetConfig() *Config {

	filePath := "config.json"

	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(strings.NewReader(string(buffer)))

	var c *Config
	if err := dec.Decode(&c); err == nil {
		log.Println("Configuration loaded")
		return c
	} else if err != nil {
		log.Fatal(err)
	}

	return c
}
