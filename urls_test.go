package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// checks if / returns status 200
func TestRootEndpoint(t *testing.T) {
	// Set up a new Gin router for testing
	gin.SetMode(gin.TestMode) // Disable debug output
	router := gin.Default()
	router.GET("/", root)

	// Create a test request
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")
}
