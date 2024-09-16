package core

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTasksRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	SetupRouter()

	TestRecord := httptest.NewRecorder()
	restCall := RestCall("GET", URL_API_TASKS, nil)
	taskApp.gin.ServeHTTP(TestRecord, restCall.RestRequest)

	if restCall.RestStatus != 200 {
		t.Error("statusCode != 200", restCall.RestStatus)
	}
}

func TestGetTasksByIdRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	SetupRouter()

	url := URL_API_TASKS + "/0"

	TestRecord := httptest.NewRecorder()
	restCall := RestCall("GET", url, nil)
	taskApp.gin.ServeHTTP(TestRecord, restCall.RestRequest)

	equals(t, 200, restCall.RestStatus)
}
