package control_utils

import (
	"errors"
	"groupie-tracker/data"
	"groupie-tracker/model"
	"reflect"
	"testing"
)

func TestFetchArtist(t *testing.T) {

	// setup : isolate this test form the real global - swap in known data
	// and restore whatever was there afterward
	original := data.ArtistByID
	defer func() {
		data.ArtistByID = original
	}()
	// mock changes
	data.ArtistByID = map[int]model.Artist{1: {ID: 1, Name: "Queen"}}
	tests := []struct {
		Name        string
		Input       int
		ExpectedOut model.Artist
		ExpectedErr error
	}{
		{
			Name:        "Existing id case",
			Input:       1,
			ExpectedOut: model.Artist{ID: 1, Name: "Queen"},
			ExpectedErr: nil,
		},
		{
			Name:        "Non-existing ID case",
			Input:       0,
			ExpectedOut: model.Artist{},
			ExpectedErr: Err404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := FetchArtist(tt.Input)
			if !reflect.DeepEqual(got, tt.ExpectedOut) {
				t.Errorf("error.... expected %v, but got %v\n", tt.ExpectedOut, got)
			}
			if !errors.Is(err, tt.ExpectedErr) {
				t.Errorf("error.... expected %s, but got %s", tt.ExpectedErr, err)
			}
		})
	}

}
