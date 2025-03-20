package api

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func Test(graphqlURL string, query string, variables map[string]any, response any) error {
	// create a client (safe to share across requests)
	client := graphql.NewClient(graphqlURL)
	req := graphql.NewRequest(query)

	// set variables
	// XXX would be nice to pass just the variables to req
	// make graphql from scratch?
	for key, value := range variables {
		req.Var(key, value)
	}

	// TODO set header fields
	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	if err := client.Run(ctx, req, response); err != nil {
		return fmt.Errorf("GraphQL request failed: %w", err)
	}
	return nil
}
