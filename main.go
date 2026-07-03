package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Character represents a single character from the API
type Character struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"` // Fixed: added string type and JSON tag
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	Origin   Origin   `json:"origin"`
	Location Location `json:"location"`
	Image    string   `json:"image"`
	Episode  []string `json:"episode"`
	Url      string   `json:"url"`
	Created  string   `json:"created"`
}

// APIResponse matches the top-level JSON structure
type Origin struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	// Requesting filtered data directly from the official API
	resp, err := client.Get("https://rickandmortyapi.com/api/character/279")
	if err != nil {
		fmt.Printf("Network error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned an error status: %s\n", resp.Status)
		return
	}

	var apiData Character
	// Stream and decode the JSON body safely
	if err := json.NewDecoder(resp.Body).Decode(&apiData); err != nil {
		fmt.Printf("Decoding failed: %v\n", err)
		return
	}

	fmt.Println("--- Rick and Morty Characters Found ---")
	fmt.Printf("ID: [%d]\nName is %s\nStatus: %s and Species: %s\nLocation: %s\n", apiData.ID, apiData.Name, apiData.Status, apiData.Species, apiData.Location.Name)
}
