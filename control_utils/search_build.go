package control_utils

import (
	"groupie-tracker/data"
	"groupie-tracker/model"
	"strconv"
	"strings"
)

func BuildSearchIndex() {
	// check artist name
	for _, artist := range data.ArtistByID {

		data.SearchIndex = append(data.SearchIndex, model.SearchEntry{
			Text:       artist.Name,
			SearchText: strings.ToLower(artist.Name),
			Type:       "artist/band",
			ArtistID:   artist.ID,
		})

		// check member
		for _, member := range artist.Members {

			data.SearchIndex = append(data.SearchIndex, model.SearchEntry{
				Text:       member,
				SearchText: strings.ToLower(member),
				Type:       "member",
				ArtistID:   artist.ID,
			})

		}

		// check creation date
		creationStr := strconv.Itoa(artist.CreationDate)

		data.SearchIndex = append(data.SearchIndex, model.SearchEntry{
			Text:       creationStr,
			SearchText: strings.ToLower(creationStr),
			Type:       "creationDate",
			ArtistID:   artist.ID,
		})

		// check first album

		data.SearchIndex = append(data.SearchIndex, model.SearchEntry{
			Text:       artist.FirstAlbum,
			SearchText: strings.ToLower(artist.FirstAlbum),
			Type:       "firtsAlbum",
			ArtistID:   artist.ID,
		})

		// check location
		for _, loc := range artist.Location.Locations {
			data.SearchIndex = append(data.SearchIndex, model.SearchEntry{
				Text:       Totitle(strings.ReplaceAll(loc, "-", " ")),
				SearchText: strings.ToLower(loc),
				Type:       "location",
				ArtistID:   artist.ID,
			})

		}
	}
}
