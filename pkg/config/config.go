package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/kayceeDev/caspa-events/internal/shared/domains"
	"github.com/kelseyhightower/envconfig"
)

func replaceQuote(word string) string {
	regexPattern := "[\"]"
	re := regexp.MustCompile(regexPattern)
	replacement := ""
	result := re.ReplaceAllString(word, replacement)
	return result
}

// var storedEnv *Config
var storedEnv *domains.Config

func GetEnv(key string, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return replaceQuote(value)
	}

	if fallback == "" {
		log.Fatalf("Missing environment variable of key: %s", key)
	}

	return replaceQuote(fallback)
}

func loadConfigMap() (*domains.Config, error) {
	log.SetOutput(os.Stderr)
	checkEnv()
	var cfg domains.Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func checkEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println("no env file located...")
	} else {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("Some error occurred. Err: %s", err)
		}
	}

}

func EnvVars() domains.Config {
	if storedEnv != nil {
		return *storedEnv
	}

	vars, err := loadConfigMap()

	storedEnv = vars

	if err != nil {
		log.Fatal(err)
	}

	return *storedEnv
}
