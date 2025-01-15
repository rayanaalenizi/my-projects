package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Relations is a slice containing multiple Relation structs
type Relations struct {
	Index []Relation `json:"index"`
}

// URL for fetching relation data
const relationURL = "https://groupietrackers.herokuapp.com/api/relation"

// FetchAllRelations fetches and returns all artist relations from the API
func FetchAllRelations() (Relations, error) {
	// Make the HTTP request to fetch relation data
	resp, err := http.Get(relationURL)
	if err != nil {
		return Relations{}, fmt.Errorf("error fetching relations: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body) // Use io.ReadAll instead of ioutil.ReadAll
	if err != nil {
		return Relations{}, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON data into the Relations struct
	var relations Relations
	err = json.Unmarshal(body, &relations)
	if err != nil {
		return Relations{}, fmt.Errorf("error unmarshalling relation data: %v", err)
	}

	return relations, nil
}

// GetRelationByArtistID returns the relation data for a specific artist using their ID
func GetRelationByArtistID(artistID int) (map[string][]string, error) {
	relations, err := FetchAllRelations()
	if err != nil {
		return nil, err
	}

	// Find the relation data for the given artist ID
	for _, relation := range relations.Index {
		if relation.ID == artistID {
			return relation.DatesLocations, nil
		}
	}

	return nil, fmt.Errorf("relation data for artist with ID %d not found", artistID)
}
