package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/4ubak/Groupie-tracker/internal/entities"
)

func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		IndexPage(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func IndexPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		t.Execute(w, nil)
		return
	}
	artists, err := GetAllData()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("./front/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, artists); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetAllData() ([]entities.Artist, error) {

	UrlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	UrlLocations := "https://groupietrackers.herokuapp.com/api/locations"
	UrlDate := "https://groupietrackers.herokuapp.com/api/dates"
	UrlRelation := "https://groupietrackers.herokuapp.com/api/relation"
	artists, err := GetArtist(UrlArtists)
	if err != nil {
		fmt.Println(err.Error())
		return artists, err
	}
	locations, err := GetLocation(UrlLocations)
	if err != nil {
		fmt.Println(err.Error())
		return artists, err
	}
	dates, err := GetDate(UrlDate)
	if err != nil {
		fmt.Println(err.Error())
		return artists, err
	}
	relations, err := GetRelation(UrlRelation)
	if err != nil {
		fmt.Println(err.Error())
		return artists, err
	}
	artists = SetData(artists, relations)
	artists = SetDataLocation(artists, locations)
	artists = SetDataDate(artists, dates)
	return artists, nil
}

func SetData(artists []entities.Artist, relations entities.Relation) []entities.Artist {
	for i := range artists {
		artists[i].Relations = parseMapToStr(relations.Index[i].DatesLocations)
	}
	return artists
}

func SetDataLocation(artists []entities.Artist, locations entities.Location) []entities.Artist {
	for i := range artists {
		artists[i].Locations = parseSliceToStr(locations.Index[i].Locations)
	}
	return artists
}

func SetDataDate(artists []entities.Artist, dates entities.ConcertDates) []entities.Artist {
	for i := range artists {
		artists[i].ConcertDates = parseSliceToStrDate(dates.Index[i].Dates)
	}
	return artists
}

func parseSliceToStrDate(dates []string) string {
	str := ""
	for i := range dates {
		str += dates[i] + "; "
	}
	return str
}

func parseSliceToStr(locations []string) string {
	str := ""
	for i := range locations {
		str += locations[i] + "; "
	}
	return str
}

func parseMapToStr(relations map[string][]string) string {
	str := ""
	for key, value := range relations {
		str += key + ":\n"
		for i := range value {
			str += value[i] + "\n"
		}
	}
	return str
}

func GetContent(url string) ([]byte, error) {

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	return contents, nil

}
func GetArtist(url string) ([]entities.Artist, error) {

	contents, err := GetContent(url)
	if err != nil {
		return nil, err
	}
	artists := make([]entities.Artist, 0)
	err2 := json.Unmarshal(contents, &artists)
	if err2 != nil {
		return nil, err2
	}
	return artists, nil

}

func GetLocation(url string) (entities.Location, error) {
	var locations entities.Location
	contents, err := GetContent(url)
	if err != nil {
		return locations, err
	}
	//locations := make(Location, 0)
	err2 := json.Unmarshal(contents, &locations)
	if err2 != nil {
		return locations, err2
	}
	return locations, nil

}
func GetDate(url string) (entities.ConcertDates, error) {
	var dates entities.ConcertDates
	contents, err := GetContent(url)
	if err != nil {
		return dates, err
	}
	//locations := make(Location, 0)
	err2 := json.Unmarshal(contents, &dates)
	if err2 != nil {
		return dates, err2
	}
	return dates, nil

}

func GetRelation(url string) (entities.Relation, error) {
	var relations entities.Relation
	contents, err := GetContent(url)
	if err != nil {
		return relations, err
	}

	err2 := json.Unmarshal(contents, &relations)
	if err2 != nil {

		return relations, err2
	}
	return relations, nil

}
