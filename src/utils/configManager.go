package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port          string `yaml:"port" envconfig:"SERVER_PORT"`
		Host          string `yaml:"host" envconfig:"SERVER_HOST"`
		AllowedOrigin string `yaml:"allowedOrigin" envconfig:"SERVER_ALLOWED"`
	} `yaml:"server"`
	Database struct {
		ConnStr string `yaml:"connStr" envconfig:"DB_CONNSTR"`
		DbName  string `yaml:"dbName" envconfig:"DB_NAME"`
	} `yaml:"database"`
}

//Reads the configuration file containing port and connection string, returns both values in a struct
func ParseConfigFile() Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func GetKey() []byte {
	filePath := filepath.FromSlash("./keys.txt")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.SplitAfter(string(data), "\n")
	signingKey := strings.Split(lines[0], " ")[1]

	return []byte(strings.TrimSpace(signingKey)) //SIGNING_KEY
}
