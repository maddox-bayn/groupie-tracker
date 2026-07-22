package model

// model to contain artist details
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     Location
	Date         Date
	Relation     Relation
}

type ArtistsPage struct {
	Artists []Artist
}

type Location struct {
	ID        int                   `json:"id"`
	Locations []string              `json:"locations"`
	DatesUrl  string                `json:"dates"`
	Coords    map[string]Coordinate `json:"coords"`
}

type Locations struct {
	Index []Location
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"date"`
}

type Dates struct {
	Index []Date
}

type Relation struct {
	ID           int                 `json:"id"`
	DateLocation map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Relation
}

// model struccture to hold all fetched data
type CombinedData struct {
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
}

// model to structure search entery
type SearchEntry struct {
	Text       string `json:"text"`
	SearchText string `json:"-"`
	Type       string `json:"type"`
	ArtistID   int    `json:"artistID"`
}

type NominatimResult struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type Coordinate struct {
	Lat float64
	Lng float64
}
