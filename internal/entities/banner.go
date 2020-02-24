package entities

type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Index []SubLocation `json:"index"`
}

type SubLocation struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type ConcertDates struct {
	Index []SubDates `json:"index"`
}

type SubDates struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Index []Index `json:"index"`
}

type Index struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
