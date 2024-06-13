package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename = ".bear.json"

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
}

const (
	usernameLabel = "Username"
	passwordLabel = "Password"
	hostnameLabel = "Hostname"
)

func Setup() error {
	cfg := Config{}

	fmt.Printf("%s:\t", usernameLabel)
	_, err := fmt.Scanln(&cfg.Username)
	if err != nil {
		return err
	}

	fmt.Printf("%s:\t", passwordLabel)
	_, err = fmt.Scanln(&cfg.Password)
	if err != nil {
		return err
	}

	fmt.Printf("%s:\t", hostnameLabel)
	_, err = fmt.Scanln(&cfg.Hostname)
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filename); err == nil {
		fmt.Println("!!!!!!!!!!!!!!!")
		fmt.Println("! File exists !")
		fmt.Println("!!!!!!!!!!!!!!!")
		fmt.Println("")
		fmt.Printf("Overwrite? (y/n)")

		var yn string
		fmt.Scanln(&yn)
		if yn == "n" {
			fmt.Println("Exiting without writing config")
			return nil
		}
	}

	os.WriteFile(filename, data, 0600)

	return nil
}
