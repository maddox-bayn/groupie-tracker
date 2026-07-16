package controllers

import (
	"groupie-tracker/data"
	"groupie-tracker/model"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// main for loading preloaded data
func TestMain(m *testing.M) {
	data.ArtistByID = map[int]model.Artist{
		1: {ID: 1, Name: "Mock Artist One"},
		2: {ID: 2, Name: "Mock Artist Two"},
	}
	var err error
	Tmpl, err = template.ParseGlob("../templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates for testing: %v", err)
	}
	// Run you r tests
	exitCode := m.Run()
	// Exit the tests process
	os.Exit(exitCode)
}
func TestHandleArtist(t *testing.T) {

	tests := []struct {
		Name           string
		UrlPath        string
		ExpectedStatus int
		Method         string
	}{
		{
			Name:           "valid id -> 200",
			UrlPath:        "/artist?id=1",
			ExpectedStatus: http.StatusOK,
			Method:         "GET",
		},
		{
			Name:           "non-numeric id",
			UrlPath:        "/artist?id=abc",
			ExpectedStatus: http.StatusBadRequest,
			Method:         "GET",
		},
		{
			Name:           "Well formed but midding id -> 404",
			UrlPath:        "/artist?id=99999999",
			ExpectedStatus: http.StatusNotFound,
			Method:         "GET",
		},
		{
			Name:           "Wrong Path -> 404",
			UrlPath:        "/art",
			ExpectedStatus: http.StatusNotFound,
			Method:         "GET",
		},
		{
			Name:           "Invalid OR NON-GET MEthod used",
			UrlPath:        "/artist?id=1",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Method:         "POST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(tt.Method, tt.UrlPath, nil)
			rr := httptest.NewRecorder()
			HandleArtist(rr, req)
			resp := rr.Result()
			defer resp.Body.Close()
			if resp.StatusCode != tt.ExpectedStatus {
				t.Errorf("error... expected %d status code but found %d status", tt.ExpectedStatus, resp.StatusCode)
			}
		})

	}
}

// test for HandleMain handler
func TestHandleMain(t *testing.T) {
	tests := []struct {
		Name    string
		UrlPath string
		Method  string
		Status  int
	}{
		{
			Name:    "Invalid path",
			UrlPath: "/invalid",
			Method:  "GET",
			Status:  http.StatusNotFound,
		},
		{
			Name:    "invalid Method",
			UrlPath: "/",
			Method:  "POST",
			Status:  http.StatusMethodNotAllowed,
		},
		{
			Name:    "valide check",
			UrlPath: "/",
			Method:  "GET",
			Status:  http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(tt.Method, tt.UrlPath, nil)
			rr := httptest.NewRecorder()
			HandleMain(rr, req)
			resp := rr.Result()
			defer resp.Body.Close()
			if resp.StatusCode != tt.Status {
				t.Errorf("expected %d but got %d", tt.Status, resp.StatusCode)
			}
		})
	}
}
