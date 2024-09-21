package main

import (
	"encoding/json"
	"fmt" // Added for printing server startup message
	"net/http"
	"strconv"
)

// PG struct represents PG data
type PG struct {
	Location string  `json:"location"`
	Price    int     `json:"price"`
	Beds     int     `json:"beds"`
	Baths    int     `json:"baths"`
}

// Sample PG data
var pgData = []PG{
	{Location: "Silk Institute", Price: 6500, Beds: 1, Baths: 1},
	{Location: "Silk Institute", Price: 7000, Beds: 1, Baths: 1},
	{Location: "Other Institute", Price: 6000, Beds: 2, Baths: 2},
}

func main() {
	http.HandleFunc("/search-pg", searchPGHandler)

	// Print message indicating server is starting
	fmt.Println("Server started successfully on port 8080")

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

// Handler for /search-pg
func searchPGHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	priceFromStr := r.URL.Query().Get("priceFrom")
	priceToStr := r.URL.Query().Get("priceTo")
	bedsStr := r.URL.Query().Get("beds")
	bathsStr := r.URL.Query().Get("baths")

	// Convert query strings to integers
	priceFrom, _ := strconv.Atoi(priceFromStr)
	priceTo, _ := strconv.Atoi(priceToStr)
	beds, _ := strconv.Atoi(bedsStr)
	baths, _ := strconv.Atoi(bathsStr)

	// Filter PG data
	var filteredPGs []PG
	for _, pg := range pgData {
		if pg.Location == location && pg.Price >= priceFrom && pg.Price <= priceTo && pg.Beds == beds && pg.Baths == baths {
			filteredPGs = append(filteredPGs, pg)
		}
	}

	// Return the filtered PGs as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredPGs)
}
