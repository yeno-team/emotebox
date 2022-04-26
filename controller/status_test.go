package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yeno-team/emotebox"
	"github.com/yeno-team/emotebox/mock"
)

func TestStatusControllerDefault(t *testing.T) {
	r := gin.Default()
	var ss mock.StatusService
	statusController := &StatusController{StatusService: &ss}

	defaultStatus := &emotebox.Status{Message: "OK", Code: 200}

	ss.StatusFn = func() (*emotebox.Status, error) {
		return defaultStatus, nil
	}

	r.GET("/api/v1/status", statusController.Default)

	req := httptest.NewRequest("GET", "/api/v1/status", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	if !ss.StatusFnInvoked {
		t.Error("Expected status function to have been called by the status controller")
	}

	if res.StatusCode != 200 {
		t.Error("Expected status code to be 200")
	}
}
