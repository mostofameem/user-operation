package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"project/web"
)

// Config represents the structure of the JSON configuration file
type Config struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func main() {

	config, err := readConfigFromFile("/home/mostofa/go_code/config.json")
	if err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	mux := web.StartServer()
	log.Printf("Server running on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, mux))
}

func readConfigFromFile(filename string) (*Config, error) {
	// Get the absolute path of the JSON file
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	// Read the JSON file
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
