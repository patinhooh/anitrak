package api

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
	"github.com/patinhooh/anitrak/internal/config"
	"github.com/patinhooh/anitrak/internal/models"
	"github.com/patinhooh/anitrak/internal/util"
)

func GetUserId() int {
	// Create a new GraphQL client
	client := graphql.NewClient(config.AnilistGraphql)

	// Create a new request
	req := graphql.NewRequest(`
		query {
			Viewer {
				id
				name
			}
		}
	`)

	// Add the Authorization header with the token
	var token *models.Token
	token, err := util.LoadToken()
	if err != nil {
		log.Fatal("Error loading token:", err)
	}
	req.Header.Set("Authorization", "Bearer "+token.Access)

	// Define a response structure to hold the result
	var resp struct {
		Viewer struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"Viewer"`
	}

	// Execute the request
	if err := client.Run(context.Background(), req, &resp); err != nil {
		log.Fatalf("failed to execute GraphQL request: %v", err)
	}

	// Print the response
	fmt.Printf("Viewer ID: %d\n", resp.Viewer.Id)
	fmt.Printf("Viewer Name: %s\n", resp.Viewer.Name)
	return resp.Viewer.Id
}
