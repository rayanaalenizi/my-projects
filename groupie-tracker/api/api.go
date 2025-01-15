package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"groupie-tracker/models" // Import models package
)

// HandleArtists fetches and returns the list of artists as JSON
func HandleArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := models.FetchAllArtists() // Fetch artists from models package
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch artists: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists) // Return artists as JSON
}

// HandleArtistByID fetches and returns an artist's details by their ID
func HandleArtistByID(w http.ResponseWriter, r *http.Request, artistID int) {
	artist, err := models.GetArtistByID(artistID) // Fetch artist by ID
	if err != nil {
		http.Error(w, fmt.Sprintf("Artist not found: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist) // Return artist details as JSON
}

// API URLs
const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

// Artist represents the artist information
type Artist struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	FirstAlbum string   `json:"firstAlbum"`
	Members    []string `json:"members"`
}

// Artists is a slice containing multiple Artist structs
type Artists []Artist

// Relation represents the relationship between an artist's locations and dates
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Relations is a slice containing multiple Relation structs
type Relations struct {
	Index []Relation `json:"index"`
}

// FetchAllArtists fetches and returns all artist data from the API
func FetchAllArtists() (Artists, error) {
	resp, err := http.Get(artistsURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var artists Artists
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

// FetchAllRelations fetches and returns all relation data from the API
func FetchAllRelations() (Relations, error) {
	resp, err := http.Get(relationURL)
	if err != nil {
		return Relations{}, fmt.Errorf("error fetching relations: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Relations{}, fmt.Errorf("error reading response body: %v", err)
	}

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

	for _, relation := range relations.Index {
		if relation.ID == artistID {
			return relation.DatesLocations, nil
		}
	}

	return nil, fmt.Errorf("relation data for artist with ID %d not found", artistID)
}
