package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ce191383/task_management/internal/middleware"
)

func TestUnauthorized_NoToken(t *testing.T) {

	protected := middleware.AuthMiddleware(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/projects",
		nil,
	)

	rr := httptest.NewRecorder()

	protected.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf(
			"expected 401 got %d",
			rr.Code,
		)
	}
}

func TestUnauthorized_InvalidToken(t *testing.T) {

	protected := middleware.AuthMiddleware(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/projects",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer abcxyz123",
	)

	rr := httptest.NewRecorder()

	protected.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf(
			"expected 401 got %d",
			rr.Code,
		)
	}
}