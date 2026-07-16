package control_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/data"
	"groupie-tracker/model"
	"net/http"
	"sync"
)

// global var used to identify 404 error
var Err404 = errors.New("404")

// FtchAllData  all data form all api endpoint
func FtchAllData() (model.CombinedData, error) {
	var (
		wg       sync.WaitGroup
		mux      sync.Mutex
		finalErr error
	)

	// helper function to fatch data cuncurrently 
	fetchdata := func(endpoint string, dest any) {
		defer wg.Done()
		err := Fetch(endpoint, dest)
		if err != nil {
			mux.Lock()
			finalErr = errors.Join(finalErr, fmt.Errorf("error fetching data from %s: %v, ", endpoint, err))
			mux.Unlock()
		}
	}

	wg.Add(4)
	go fetchdata("/artists", &data.Artists)
	go fetchdata("/locations", &data.Locations)
	go fetchdata("/dates", &data.Dates)
	go fetchdata("/relation", &data.Relations)
	wg.Wait()

	if finalErr != nil {
		return model.CombinedData{}, finalErr
	}
	return model.CombinedData{
		Artists:   data.Artists,
		Locations: data.Locations.Index,
		Dates:     data.Dates.Index,
		Relations: data.Relations.Index,
	}, nil
}
// FetchArtist return all artist detail using map look up 
// if data is not present it returns error  404
func FetchArtist(Id int) (model.Artist, error) {
	artist, found := data.ArtistByID[Id]
	if !found {
		return model.Artist{}, Err404
	}
	return artist, nil
}

// helper function to get content of api content into required destination 
// if err ocured error will be returned
func Fetch(endpoint string, dest any) error {
	urlResp, err := http.Get(config.Api_url + endpoint)
	if err != nil {
		return err
	}
	defer urlResp.Body.Close()

	if urlResp.StatusCode != http.StatusOK {
		return fmt.Errorf("api return status code %d", urlResp.StatusCode)
	}
	err = json.NewDecoder(urlResp.Body).Decode(dest)
	if err != nil {
		return err
	}
	return nil
}
