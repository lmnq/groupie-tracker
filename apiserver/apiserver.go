package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Artist ..
type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
}

// Relation ..
type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// const ..
const (
	APIUrl = "https://groupietrackers.herokuapp.com/api"
)

// Artists ..
var (
	Artists   []Artist
	Relations Relation
)

// initialize API Structures ..
func init() {
	err := GetAPIData(fmt.Sprintf("%s/artists", APIUrl), &Artists)
	if err != nil {
		log.Println(err)
	}
	err = GetAPIData(fmt.Sprintf("%s/relation", APIUrl), &Relations)
	if err != nil {
		log.Println(err)
	}
	relationsToArtists()
}

// GetAPIData ..
func GetAPIData(url string, str interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &str)
	if err != nil {
		return err
	}
	return nil
}
