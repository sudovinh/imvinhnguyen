package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockDB struct{}

func (mdb *mockDB) Close() error {
	return nil
}

func (mdb *mockDB) Health() map[string]string {
	return map[string]string{
		"status": "healthy",
	}
}

func TestRegisterRoutes(t *testing.T) {
	s := &Server{db: &mockDB{}}
	e := s.RegisterRoutes()

	tests := []struct {
		method       string
		target       string
		expectedCode int
	}{
		{"GET", "/", http.StatusOK},
		{"GET", "/links", http.StatusOK},
		{"GET", "/health", http.StatusOK},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.target, nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, tt.expectedCode, rec.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	e := echo.New()
	s := &Server{db: &mockDB{}}
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, s.healthHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"status":"healthy"}`, rec.Body.String())
	}
}
