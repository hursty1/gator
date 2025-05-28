package config

import (
	"encoding/json"
	"os"
)

func getConfigFile() (Config, error) {
	data, err := os.Open(getConfigFilePath())
	if err != nil {
		return Config{}, err
	}

	defer data.Close()

	decoder := json.NewDecoder(data)
	config := Config{}
	if err = decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	
	return config, nil
}


func write(c Config) error {
	
	file, err := os.Create(getConfigFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(c); err != nil {
		return err
	}
	return nil
}