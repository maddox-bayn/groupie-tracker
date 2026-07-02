package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Character represents a single character from the API
type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"` // Fixed: added string type and JSON tag
	Status  string `json:"status"`
	Species string `json:"species"`
}

// APIResponse matches the top-level JSON structure
type APIResponse struct {
	Results []Character `json:"results"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	// Requesting filtered data directly from the official API
	resp, err := client.Get("https://rickandmortyapi.com")
	if err != nil {
		fmt.Printf("Network error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned an error status: %s\n", resp.Status)
		return
	}

	var apiData APIResponse
	// Stream and decode the JSON body safely
	if err := json.NewDecoder(resp.Body).Decode(&apiData); err != nil {
		fmt.Printf("Decoding failed: %v\n", err)
		return
	}

	fmt.Println("--- Rick and Morty Characters Found ---")
	for _, character := range apiData.Results {
		fmt.Printf("[%d] %s (%s - %s)\n", character.ID, character.Name, character.Status, character.Species)
	}
}
