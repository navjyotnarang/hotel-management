package api

import (
	"net/http"

	"github.com/gofr/gofr"
)

var rooms []Room

func init() {
	// Sample data for testing
	rooms = append(rooms, Room{ID: 1, Number: 101, Capacity: 2, Status: "Available"})
	rooms = append(rooms, Room{ID: 2, Number: 102, Capacity: 4, Status: "Occupied"})
	// Add more sample data as needed
}

func getAllRooms(w http.ResponseWriter, r *http.Request) {
	gofr.JSON(w, http.StatusOK, rooms)
}

func getRoomByID(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve room by ID
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create a new room
}

func updateRoom(w http.ResponseWriter, r *http.Request) {
	// Implement logic to update an existing room
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	// Implement logic to delete a room
}

// SetupRoutes initializes API routes
func SetupRoutes(router *gofr.Router) {
	router.Get("/rooms", getAllRooms)
	router.Get("/rooms/{id}", getRoomByID)
	router.Post("/rooms", createRoom)
	router.Put("/rooms/{id}", updateRoom)
	router.Delete("/rooms/{id}", deleteRoom)
}
