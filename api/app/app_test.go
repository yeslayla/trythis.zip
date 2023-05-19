package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yeslayla/trythis.zip/api/api"
)

func TestGetStatus(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	// Intialize app
	app := App{}
	router := app.GetRouter()

	// Send test request to router
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("GET", "/status", nil))

	assert.Equal(http.StatusOK, w.Code)

	// Parse response object
	response := &api.ErrorResponse{}
	err := json.Unmarshal(w.Body.Bytes(), response)
	require.Nilf(err, "Failed to parse response: %s")

	assert.Equal(int32(http.StatusOK), response.Code)
	assert.Greater(len(response.Message), 0)
}
