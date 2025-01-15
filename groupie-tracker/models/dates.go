package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Date represents the dates for a specific artist
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Dates is a slice containing multiple Date structs
type Dates struct {
	Index []Date `json:"index"`
}

// URL for fetching dates data
const datesURL = "https://groupietrackers.herokuapp.com/api/dates"

// FetchAllDates fetches and returns all artist dates from the API
func FetchAllDates() (Dates, error) {
	// Make the HTTP request to fetch dates data
	resp, err := http.Get(datesURL)
	if err != nil {
		return Dates{}, fmt.Errorf("error fetching dates: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Dates{}, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON data into the Dates struct
	var dates Dates
	err = json.Unmarshal(body, &dates)
	if err != nil {
		return Dates{}, fmt.Errorf("error unmarshalling dates data: %v", err)
	}

	return dates, nil
}

// GetDatesByArtistID returns the dates for a specific artist using their ID
func GetDatesByArtistID(artistID int) ([]string, error) {
	dates, err := FetchAllDates()
	if err != nil {
		return nil, err
	}

	// Find the dates data for the given artist ID
	for _, date := range dates.Index {
		if date.ID == artistID {
			return date.Dates, nil
		}
	}

	return nil, fmt.Errorf("dates for artist with ID %d not found", artistID)
}
