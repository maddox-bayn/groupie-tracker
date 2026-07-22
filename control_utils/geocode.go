package control_utils

import (
	"groupie-tracker/data"
	"strings"
)

func PreloadGeocodes() {

	// create catch
	geoCatche := make(map[string]bool)
	for _, locData := range data.Locations.Index {
		for _, loc := range locData.Locations {
			cleanLoc := Totitle(strings.ReplaceAll(loc, "-", " "))
		    geoCatche[cleanLoc] = true
		}
	}
}