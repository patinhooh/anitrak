package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/patinhooh/anitrak/internal/config"
)

// Response struct for unmarshalling the API response
type Response struct {
	Data struct {
		Page struct {
			MediaList []MediaList `json:"mediaList"`
		} `json:"Page"`
	} `json:"data"`
}

// MediaList wraps the media data for the list response
type MediaList struct {
	Media Media `json:"media"`
}

// Media represents the media object returned from the GraphQL query
type Media struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Episodes    int    `json:"episodes"`
	BannerImage string `json:"bannerImage"`
	Title       struct {
		UserPreferred string `json:"userPreferred"`
	} `json:"title"`
	CoverImage struct {
		Large string `json:"large"`
	} `json:"coverImage"`
	NextAiringEpisode *struct {
		AiringAt        int `json:"airingAt"`
		TimeUntilAiring int `json:"timeUntilAiring"`
		Episode         int `json:"episode"`
	} `json:"nextAiringEpisode"`
}

// makeGraphQLCall sends the GraphQL request using HTTP and returns the response
func makeGraphQLCall(userId int, mediaType string, perPage int) (*Response, error) {
	// Create the query and variables in a map
	query := `query($userId: Int, $type: MediaType, $perPage: Int) {
		Page(perPage: $perPage) {
			mediaList(userId: $userId, type: $type, status_in: [CURRENT, REPEATING], sort: UPDATED_TIME_DESC) {
				media {
					id
					type
					status
					episodes
					bannerImage
					title {
						userPreferred
					}
					coverImage {
						large
					}
					nextAiringEpisode {
						airingAt
						timeUntilAiring
						episode
					}
				}
			}
		}
	}`

	vars := map[string]interface{}{
		"userId":  userId,
		"type":    mediaType,
		"perPage": perPage,
	}

	// Create the request payload
	payload := map[string]interface{}{
		"query":     query,
		"variables": vars,
	}

	// Marshal payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", config.AnilistGraphql, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers (Content-Type, and possibly Authorization if needed)
	req.Header.Set("Content-Type", "application/json")
	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if the response is successful (HTTP status 200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %s, body: %s", resp.Status, body)
	}

	// Unmarshal the JSON response into the Response struct
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// Return the response
	return &response, nil
}

// PrintAnimeList is the main function to fetch and print the anime list
func PrintAnimeList() {
	// Example usage
	userId := 6476156    // Use the Viewer ID you fetched
	mediaType := "ANIME" // Specify the type of media
	perPage := 10        // Number of items per page

	fmt.Println("Fetching media list...")

	// Make the GraphQL call
	response, err := makeGraphQLCall(userId, mediaType, perPage)
	if err != nil {
		fmt.Println("Error making GraphQL call:", err)
		return
	}

	// Check if the response contains media
	if len(response.Data.Page.MediaList) == 0 {
		fmt.Println("No media found for the given Viewer ID.")
		return
	}

	fmt.Println("Fetched media list:")
	for _, mediaListItem := range response.Data.Page.MediaList {
		media := mediaListItem.Media
		fmt.Printf("ID: %d, Title: %s, Status: %s, Episodes: %d\n", media.ID, media.Title.UserPreferred, media.Status, media.Episodes)
		// fmt.Printf("Banner Image: %s\n", media.BannerImage)
		// fmt.Printf("Cover Image: %s\n", media.CoverImage.Large)

		// If there is a next airing episode, print the details
		if media.NextAiringEpisode != nil {
			fmt.Printf("Next Episode: %d, Airing At: %d, Time Until Airing: %d\n", media.NextAiringEpisode.Episode, media.NextAiringEpisode.AiringAt, media.NextAiringEpisode.TimeUntilAiring)
		} else {
			fmt.Print("No upcoming episode.")
		}
		fmt.Print("\n\n")
	}
}
