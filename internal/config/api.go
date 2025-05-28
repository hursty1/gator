package config


func Read() (Config, error) {
	
	return getConfigFile()
}

func (c *Config) SetUser(username string) error {
	c.Current_user_name = username
	return write(*c)
}