package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func HandlChar(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Requesting filtered data directly from the official API
	resp, err := client.Get("https://rickandmortyapi.com/api/character/398")
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
		http.Error(w, "Error decoding json Character data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(apiData); err != nil {
		http.Error(w, "Error Ecoding Character data", http.StatusInternalServerError)
		return
	}
}

func HomeHanler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<div class=\"mr-2 text-xs inline-flex items-center font-bold leading-sm uppercase px-4 py-2 bg-blue-800 rounded\"><a href=\"/rickCharacter\">Character Information</a></div>")
}
func main() {
	http.HandleFunc("/", HomeHanler)
	http.HandleFunc("/rickCharacter", HandlChar)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
