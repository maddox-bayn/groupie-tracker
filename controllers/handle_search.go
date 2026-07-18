package controllers

import (
	"encoding/json"
	"groupie-tracker/data"
	"groupie-tracker/model"
	"net/http"
	"strings"
)

// handler Seach fo r handling searchse for bat
func Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		RenderError(w, http.StatusNotFound, "Request path not found")
		return
	}

	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	query := strings.ToLower(r.URL.Query().Get("q"))
	if strings.TrimSpace(query) == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]model.SearchEntry{})
		return
	}

	var results []model.SearchEntry
	for _, entery := range data.SearchIndex {
		if strings.Contains(entery.SearchText, query) {
			results = append(results, entery)
		}
	}
	if len(results) > 20 {
		results = results[:20]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
