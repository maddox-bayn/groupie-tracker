package control_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/data"
	"groupie-tracker/model"
	"net/http"
	"strings"
	"sync"
)

var Err404 = errors.New("404")

// function to
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
func FetchArtist(Id int) (model.Artist, error) {
	var artist model.Artist
	found := false
	for _, v := range data.CombinedData.Artists {
		if v.ID == Id {
			found = true
			artist.ID = v.ID
			artist.CreationDate = v.CreationDate
			artist.Image = v.Image
			artist.Members = v.Members
			artist.Name = v.Name
			artist.FirstAlbum = v.FirstAlbum
		}
	}

	if !found {
		return model.Artist{}, Err404
	}

	var locat model.Location
	for _, loc := range data.Locations.Index {
		if loc.ID == Id {
			locat.Locations = loc.Locations
		}
	}
	var date model.Date
	for _, dat := range data.Dates.Index {
		if dat.ID == Id {
			date.ID = Id
			date.Dates = dat.Dates
		}
	}
	var relation model.Relation
	formattedRelations := make(map[string][]string)
	for _, rela := range data.Relations.Index {
		if rela.ID == Id {
			for k, v := range rela.DateLocation {
				k = strings.ReplaceAll(k, "_", " ")
				k = strings.ReplaceAll(k, "-", " ")
				k = Totitle(k)
				formattedRelations[k] = v
			}
			relation.DateLocation = formattedRelations
		}
	}
	artist.Date = date
	artist.Location = locat
	artist.Relation = relation

	return artist, nil
}
func Totitle(key string) string {
	words := strings.Fields(key)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	keyForm := strings.Join(words, " ")
	keyForm = strings.ReplaceAll(keyForm, "Usa", "USA")
	keyForm = strings.ReplaceAll(keyForm, "Uk", "UK")
	return keyForm
}
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
