package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	// Data is the configuration data for the application.
	Data config
)

const (
	configFile     = "./conf/config.json"
	defaultBaseURL = "https://api.clashofclans.com/v1"
)

type config struct {
	BaseURL string `json:"base_url"`
}

// init reads in the configuration file
func init() {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		log.Errorf("unable to open file %s", configFile)
		// Use the default values for the configuration
		Data = config{
			BaseURL: defaultBaseURL,
		}
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Data)
}
