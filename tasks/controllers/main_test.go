package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mulatinho/golabs/tasks/controllers"
	"github.com/stretchr/testify/assert"
)

func TestGetTasksRoute(t *testing.T) {
	router := controllers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetTasksByIdRoute(t *testing.T) {
	router := controllers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
