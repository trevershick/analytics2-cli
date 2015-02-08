package config

import (
	"os/user"
	"encoding/json"
	"os"
	// "fmt"
	"strings"
)

type Configuration struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func getConfigFileName() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	s := []string{ usr.HomeDir, ".a2mrc" }
	return strings.Join(s, "/"), nil
}

func GetConfiguration() (*Configuration,error) {
	rcFileName, _ := getConfigFileName()

	if _, err := os.Stat(rcFileName); os.IsNotExist(err) {
		//fmt.Printf("Creating initial configuration file in %s", rcFileName)
		err = UpdateConfiguration(Configuration{
			UserName:"nobody",
			Password:"unknown",
		})
		if err != nil {
			return nil, err
		}
	}


	// fmt.Printf("Loading %s", rcFileName)
	file, err := os.Open(rcFileName)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	if err := decoder.Decode(&configuration); err != nil {
		return nil, err
	}
	return &configuration, nil
}

func UpdateConfiguration(configuration Configuration) error {
	configFile, err := getConfigFileName()
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	//fmt.Printf("Encode %s", configuration)
	err = encoder.Encode(configuration)
	return err
}
