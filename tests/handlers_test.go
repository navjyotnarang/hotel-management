package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/hotel-management/api"
)

func TestGetAllRooms(t *testing.T) {
	router := gofr.NewRouter()
	api.SetupRoutes(router)

	req, err := http.NewRequest("GET", "/rooms", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var rooms []api.Room
	err = json.Unmarshal(rr.Body.Bytes(), &rooms)
	assert.NoError(t, err)

	// Add more assertions based on your expected API response
}
