package main

import (
	"fmt"

	"github.com/patinhooh/anitrak/internal/api"
	"github.com/patinhooh/anitrak/internal/config"
)

var cfg *config.Config

type Response struct {
	User struct {
		Name string `json:"name"`
	} `json:"User"`
}

func main() {
	var err error
	cfg, err = config.InitConfig()
	if err != nil {
		fmt.Println("Warning: Using default config due to error:", err)
	}

	var response Response
	api.Test(
		cfg.GraphqlURL,
		`
		query($userId: Int)  {
			User(id: $userId) {
				name
			}
		}`,
		map[string]interface{}{
			"userId": 1,
		},
		&response,
	)

	fmt.Printf("User Name: %s\n", response.User.Name)
}
