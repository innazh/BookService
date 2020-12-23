package settings

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Settings struct {
	Port             string
	ConnectionString string
}

//Reads the configuration file containing port and connection string, returns both values in a struct
func GetSettings() Settings {
	filePath := filepath.FromSlash("./.config.txt")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.SplitAfter(string(data), "\n")
	port := strings.Split(lines[0], " ")[1]
	connStr := strings.Split(lines[1], " ")[1]

	settings := Settings{
		strings.TrimSpace(port),    //PORT
		strings.TrimSpace(connStr), //CONNECTION_STRING
	}

	return settings
}

func GetKey() []byte {
	filePath := filepath.FromSlash("./.keys")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.SplitAfter(string(data), "\n")
	signingKey := strings.Split(lines[0], " ")[1]

	return []byte(strings.TrimSpace(signingKey)) //SIGNING_KEY
}
