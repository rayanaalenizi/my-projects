package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Artist represents an artist with their basic details.
type Artist struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"` // Year of formation
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      []string            `json:"-"` // Populate locations separately
	DatesLocations map[string][]string `json:"-"` // Add DatesLocations to store concert locations and dates

}

// Artists is a slice of Artist.
type Artists []Artist

// URL for the artists' data
const artistURL = "https://groupietrackers.herokuapp.com/api/artists"

// FetchAllArtists fetches and returns all artist data from the API
func FetchAllArtists() ([]Artist, error) {
	resp, err := http.Get(artistURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling artist data: %v", err)
	}

	return artists, nil
}

// GetArtistByID fetches a specific artist by ID
func GetArtistByID(artistID int) (*Artist, error) {
	artists, err := FetchAllArtists()
	if err != nil {
		return nil, err
	}

	for _, artist := range artists {
		if artist.ID == artistID {
			return &artist, nil
		}
	}

	return nil, fmt.Errorf("artist with ID %d not found", artistID)
}
