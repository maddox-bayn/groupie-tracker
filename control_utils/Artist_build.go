package control_utils

import (
	"groupie-tracker/data"
	"groupie-tracker/model"
	"strings"
)

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

func BuildArtistIndex() {
	data.ArtistByID = make(map[int]model.Artist)
	locationIndex := make(map[int]model.Location)
	dateIndex := make(map[int]model.Date)
	relationIndex := make(map[int]model.Relation)

	for _, loc := range data.Locations.Index {
		locationIndex[loc.ID] = loc
	}
	for _, date := range data.Dates.Index {
		dateIndex[date.ID] = date
	}

	for _, rel := range data.Relations.Index {
		formattedRelations := make(map[string][]string)
		for k, v := range rel.DateLocation {
			k = strings.ReplaceAll(k, "_", " ")
			k = strings.ReplaceAll(k, "-", " ")
			k = Totitle(k)
			formattedRelations[k] = v
		}
		rel.DateLocation = formattedRelations
		relationIndex[rel.ID] = rel
	}
	for _, artist := range data.Artists {
		artist.Location = locationIndex[artist.ID]
		artist.Date = dateIndex[artist.ID]
		artist.Relation = relationIndex[artist.ID]
		data.ArtistByID[artist.ID] = artist
	}
}
