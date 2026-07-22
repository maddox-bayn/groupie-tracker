package data

import "groupie-tracker/model"

// global state variable
var (
	Artists      []model.Artist
	Locations    model.Locations
	Dates        model.Dates
	Relations    model.Relations
	CombinedData model.CombinedData
	ArtistByID   map[int]model.Artist
	SearchIndex []model.SearchEntry
	GeoCache = make(map[string]model.Coordinate)
)
