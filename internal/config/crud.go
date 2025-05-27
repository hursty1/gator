package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func getConfigFile() (Config, error) {
	data, err := os.Open(getConfigFilePath())
	if err != nil {
		return Config{}, err
	}

	defer data.Close()

	byteData, _ := ioutil.ReadAll(data)
	var config Config

	err = json.Unmarshal(byteData, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}


func (c Config) write(username string) error {
	c.Current_user_name = username
	byteData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(getConfigFilePath(), byteData, 0644)
	if err != nil {
		return err
	}
	return nil
}