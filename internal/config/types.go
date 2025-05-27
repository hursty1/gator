package config

import (
	"fmt"
	"os"
)


type Config struct {
	Db_url 					string 		`json:"db_url"`
	Current_user_name  		string 		`json:"current_user_name"`
}


const configFileName = "boot.dev/gator/.gatorconfig.json"

func getConfigFilePath() string {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	config_fp := home_dir + "/" + configFileName
	return config_fp
}