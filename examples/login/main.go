package main

import (
	"fmt"
	"log"
	"os"

	cloudreve "github.com/Karlbaey/Go-Cloudreve-SDK"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Cloudreve struct {
		URL      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"cloudreve"`
}

// load YAML
func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	log.Println("Loading config from config.yaml...")
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Load config failed: %v", err)
	}

	log.Printf("Creating a Cloudreve client. Target server: %s", config.Cloudreve.URL)
	client := cloudreve.NewClient(config.Cloudreve.URL)

	// Call Login
	log.Printf("Try to login as user '%s'...", config.Cloudreve.Username)
	user, err := client.Login(config.Cloudreve.Username, config.Cloudreve.Password)
	if err != nil {
		log.Fatalf("Login failed: %x", err)
	}

	// Success
	log.Printf("Login success. Welcome, %s!", user.Nickname)
	fmt.Println("------------------------------------")
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Email: %s\n", user.UserName)
	fmt.Printf("User Group: %s\n", user.Group.Name)
	fmt.Println("------------------------------------")
}
