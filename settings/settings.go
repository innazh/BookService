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
	filePath := filepath.FromSlash("./config.txt")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.SplitAfter(string(data), "\n")
	port := strings.Split(lines[0], " ")[1]
	conn_str := strings.Split(lines[1], " ")[1]

	settings := Settings{
		strings.TrimSpace(port),     //PORT
		strings.TrimSpace(conn_str), //CONNECTION_STRING
	}

	return settings
}
