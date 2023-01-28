package main

import (
	"encoding/json"
	"github/ho3eintry/ModernGolangProgramming/dino/dynoWebPortal"
	"log"
	"os"
)

type configuration struct {
	ServerAddress string `json:"web_server"`
}

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Printf("starting web server on address %s", config.ServerAddress)
	dynowebportal.RunWebPortal(config.ServerAddress)
}
