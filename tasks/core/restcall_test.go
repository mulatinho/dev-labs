package core

import (
	"log"
	"net/http/httptest"
	"testing"
)

func RestCallSimpleTest(t *testing.T) {
	SetupRouter()
	TestRecord := httptest.NewRecorder()
	restCall := RestCall("GET", URL_API_TASKS, nil)
	taskApp.gin.ServeHTTP(TestRecord, restCall.RestRequest)

	if restCall.RestError != nil {
		log.Println(restCall.RestError)
	}

	equals(t, 200, restCall.RestStatus)
}
