package data

import "groupie-tracker/model"

var (
	Artists      []model.Artist
	Locations    model.Locations
	Dates        model.Dates
	Relations    model.Relations
	CombinedData model.CombinedData
	ArtistByID   map[int]model.Artist
)
