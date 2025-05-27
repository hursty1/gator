package main

import (
	"fmt"

	"github.com/hursty1/gator/internal/config"
)
func main() {
	fmt.Println("hello")

	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.SetUser("Hurst")
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.Current_user_name)
}