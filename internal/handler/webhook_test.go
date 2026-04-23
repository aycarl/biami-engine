package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMobileHookHandler_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := []byte(`{"from":"+123456789","text":"*123#","session_id":"xyz123"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/mobile-hook", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	MobileHookHandler(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	if got := w.Body.String(); got == "" {
		t.Fatalf("expected non-empty response body")
	}
}

func TestMobileHookHandler_BadPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/mobile-hook", bytes.NewReader([]byte(`{"from":`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	MobileHookHandler(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

