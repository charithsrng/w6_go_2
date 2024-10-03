package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"time"

	"github.com/gorilla/mux"
)

// Grocery struct (Model)
type Grocery struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	//Expiery  time.Time `json:"date"`
}

// In-memory database
var groceries []Grocery

// Create a new grocery item (POST /groceries)
func createGrocery(w http.ResponseWriter, r *http.Request) {
	var grocery Grocery
	json.NewDecoder(r.Body).Decode(&grocery)
	grocery.ID = len(groceries) + 1
	groceries = append(groceries, grocery)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grocery)
}

// Get all groceries (GET /groceries)
func getGroceries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groceries)
}

// Get a single grocery item by ID (GET /groceries/{id})
func getGrocery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, grocery := range groceries {
		if grocery.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(grocery)
			return
		}
	}
	http.Error(w, "Grocery not found", http.StatusNotFound)
}

// Update a grocery item (PUT /groceries/{id})
func updateGrocery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, grocery := range groceries {
		if grocery.ID == id {
			var updatedGrocery Grocery
			json.NewDecoder(r.Body).Decode(&updatedGrocery)
			groceries[i].Name = updatedGrocery.Name
			groceries[i].Category = updatedGrocery.Category
			groceries[i].Quantity = updatedGrocery.Quantity
			groceries[i].Price = updatedGrocery.Price
			//groceries[i].Expiery = updatedGrocery.Expiery
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(groceries[i])
			return
		}
	}
	http.Error(w, "Grocery not found", http.StatusNotFound)
}

// Delete a grocery item (DELETE /groceries/{id})
func deleteGrocery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, grocery := range groceries {
		if grocery.ID == id {
			groceries = append(groceries[:i], groceries[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Grocery not found", http.StatusNotFound)
}

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Route Handlers
	router.HandleFunc("/groceries", createGrocery).Methods("POST")
	router.HandleFunc("/groceries", getGroceries).Methods("GET")
	router.HandleFunc("/groceries/{id}", getGrocery).Methods("GET")
	router.HandleFunc("/groceries/{id}", updateGrocery).Methods("PUT")
	router.HandleFunc("/groceries/{id}", deleteGrocery).Methods("DELETE")

	// Start the server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
