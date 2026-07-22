package control_utils

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var geoMutex sync.RWMutex

func PreloadGeocodes() {
	fmt.Println("🗺️  Background Geocode Preloader Started...")
	// create catch
	uniqueLocations := make(map[string]bool)
	for _, locData := range data.Locations.Index {
		for _, loc := range locData.Locations {
			cleanLoc := Totitle(strings.ReplaceAll(loc, "-", " "))
		    uniqueLocations[cleanLoc] = true
		}
	}
	// Loop through and execute the fetch
	for  loc := range uniqueLocations {
		// calling Geocode 
		_, err := Geocode(loc); if err != nil {
			fmt.Printf("⚠️ Failed to geocode %s: %v\n", loc, err)
		}
	}
	fmt.Println("✅ Background Geocoding Complete!")
}

    // build Nominatim URL with address as query param
    // http.Get
    // decode into []nominatimResult (array!)
    // check len(results) == 0 -> not found error
    // parse Lat/Lon strings -> float64
    // return model.Coordinate{Lat: ..., Lng: ...}
func Geocode(address string) (model.Coordinate, error) {
	// checking the cache first
	geoMutex.RLock()
	if coords, exists := data.GeoCache[address]; exists {
		geoMutex.RUnlock()
		return coords, nil
	}
	geoMutex.RUnlock()
	// http and parsing logic
	escapedLoc := url.QueryEscape(address)
	apiUrl := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", escapedLoc)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", apiUrl, nil); if err != nil {
		return model.Coordinate{}, err
	}
	req.Header.Set("User-Agent", "GroupieTrackerApp//1.0")
	resp, err := client.Do(req); if err != nil {
		return model.Coordinate{}, err
	}
	defer resp.Body.Close()

	var result []model.NominatimResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return model.Coordinate{}, err
	}

	// 6. Handle empty results (location not found)
	if len(result) == 0 {
		return model.Coordinate{}, fmt.Errorf("location not found")
	}

	latFloat, err := strconv.ParseFloat(result[0].Lat, 64); if err != nil {
		return model.Coordinate{}, err
	} 

	lonFloat, err := strconv.ParseFloat(result[0].Lon, 64); if err != nil {
		return model.Coordinate{}, err
	} 
	coords := model.Coordinate{
		Lat: latFloat,
		Lng: lonFloat,
	}

	// save to cache and sleep
	geoMutex.Lock()
	data.GeoCache[address]= coords
	geoMutex.Unlock()

	time.Sleep(1 * time.Second)
	return coords, nil
}