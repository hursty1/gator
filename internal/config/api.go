package config

import "fmt"


func Read() (Config, error) {
	
	return getConfigFile()
}

func (c Config) SetUser(username string) error {
	err := c.write(username)
	if err != nil {
		fmt.Println("Failed to write file")
		return err
	}
	return nil
}