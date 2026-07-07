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

func FtchAllData() (model.CombinedData, error) {
	var (
		wg       sync.WaitGroup
		mux      sync.Mutex
		finalErr error
	)

	fetchdata := func(endpoint string, dest any) {
		defer wg.Done()
		err := Fetch(endpoint, dest)
		if err != nil {
			mux.Lock()
			err = errors.Join(finalErr, fmt.Errorf("Error fetching data from %s: %v, ", endpoint, err))
			mux.Unlock()
		}
	}

	wg.Add(4)
	go fetchdata("/artist", &data.Artists)
	go fetchdata("/location", &data.Locations)
	go fetchdata("/date", &data.Dates)
	go fetchdata("/relation", &data.Relations)
	wg.Wait()

	if finalErr != nil {
		return model.CombinedData{}, finalErr
	}
	return model.CombinedData{
		Artist:   data.Artists,
		Location: data.Locations.Index,
		Date:     data.Dates.Index,
		Relation: data.Relations.Index,
	}, nil
}
func Fetch(endpoint string, dest any) error {
	Urlresp, err := http.Get(config.Api_url + endpoint)
	if err != nil {
		return err
	}
	defer Urlresp.Body.Close()

	if Urlresp.StatusCode != http.StatusOK {
		return fmt.Errorf("Api return status code %d", Urlresp.StatusCode)
	}
	err = json.NewDecoder(Urlresp.Body).Decode(dest)
	if err != nil {
		return err
	}
	return nil
}
