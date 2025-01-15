package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/models" // Adjust import based on your project structure
)

// Serve the list of all artists
func handleArtistsPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		renderErrorPage(w, http.StatusNotFound)
		return
	}

	// Fetch all artists
	artists, err := models.FetchAllArtists()
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
		return
	}

	// Parse the index (artists list) template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
		return
	}

	// Render the template with the list of artists
	err = tmpl.Execute(w, artists)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
	}
}


// Serve the artist details page
func handleArtistDetailsPage(w http.ResponseWriter, r *http.Request) {
	artistIDStr := r.URL.Path[len("/artist/"):]

	// Convert the artist ID from string to integer
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		renderErrorPage(w, http.StatusBadRequest)
		return
	}

	// Fetch the artist details from the model
	artist, err := models.GetArtistByID(artistID)
	if err != nil {
		renderErrorPage(w, http.StatusNotFound)
		return
	}

	// Fetch the artist's relation (dates and locations)
	datesLocations, err := models.GetRelationByArtistID(artistID)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
		return
	}

	// Create a new map to store the formatted locations
	formattedDatesLocations := make(map[string][]string)

	for location, dates := range datesLocations {
		formattedLocation := FormatLocation(location)
		formattedDatesLocations[formattedLocation] = dates
	}

	artist.DatesLocations = formattedDatesLocations

	// Parse the artist-details template
	tmpl, err := template.ParseFiles("templates/artist-details.html")
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
		return
	}

	// Render the template with the artist details and dates/locations
	err = tmpl.Execute(w, artist)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
	}
}
func renderErrorPage(w http.ResponseWriter, statusCode int) {
	var tmpl *template.Template
	var err error

	switch statusCode {
	case http.StatusNotFound:
		tmpl, err = template.ParseFiles("templates/404.html")
	case http.StatusBadRequest:
		tmpl, err = template.ParseFiles("templates/400.html")
	case http.StatusInternalServerError:
		tmpl, err = template.ParseFiles("templates/500.html")
	default:
		tmpl, err = template.ParseFiles("templates/500.html") // Fallback to 500 for other errors
	}

	if err != nil {
		http.Error(w, "Unable to load error template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render error template", http.StatusInternalServerError)
	}
}
// FormatLocation formats the location string by replacing underscores and capitalizing
func FormatLocation(location string) string {
	// Replace underscores with spaces
	location = strings.ReplaceAll(location, "_", " ")

	// Split by dash to separate city and country
	parts := strings.Split(location, "-")

	// Capitalize the first letter of each word
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Rejoin with dash
	return strings.Join(parts, ", ")
}
func main() {

	// Serve the root URL to display the artist list
	http.HandleFunc("/", handleArtistsPage)
	// Serve the artist list and artist details pages
	http.HandleFunc("/artist/", handleArtistDetailsPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
