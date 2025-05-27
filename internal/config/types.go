package config


type Config struct {
	Db_url 		string 		`json:"db_url"`
}


func Get() (Config, error) {
	data, err
}