package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Location represents an artist's concert locations and the URL to the dates
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"` // URL to fetch concert dates
}

// Locations is a slice of Location structs
type Locations struct {
	Index []Location `json:"index"`
}

// URL for fetching locations data
const locationsURL = "https://groupietrackers.herokuapp.com/api/locations"

// FetchAllLocations fetches and returns all artist locations from the API
func FetchAllLocations() (Locations, error) {
	// Make the HTTP request to fetch locations data
	resp, err := http.Get(locationsURL)
	if err != nil {
		return Locations{}, fmt.Errorf("error fetching locations: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body) // Using io.ReadAll instead of ioutil.ReadAll
	if err != nil {
		return Locations{}, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON data into the Locations struct
	var locations Locations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Locations{}, fmt.Errorf("error unmarshalling locations data: %v", err)
	}

	return locations, nil
}

// GetLocationsByArtistID returns the locations for a specific artist using their ID
func GetLocationsByArtistID(artistID int) ([]string, error) {
	locations, err := FetchAllLocations()
	if err != nil {
		return nil, err
	}

	// Find the location data for the given artist ID
	for _, location := range locations.Index {
		if location.ID == artistID {
			return location.Locations, nil
		}
	}

	return nil, fmt.Errorf("locations for artist with ID %d not found", artistID)
}
