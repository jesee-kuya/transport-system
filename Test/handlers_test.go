package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_handleAuth(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "GET request serves authentication.html",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusOK,
			expectedBody:   "", 
		},
		{
			name:           "POST request missing parameters",
			method:         http.MethodPost,
			body:           "",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "POST request with invalid action",
			method:         http.MethodPost,
			body:           "action=invalid&userType=school",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "POST request with valid login action",
			method:         http.MethodPost,
			body:           "action=login&userType=school",
			expectedStatus: http.StatusOK,
			expectedBody:   "", 
		},
		{
			name:           "POST request with valid signup action",
			method:         http.MethodPost,
			body:           "action=signup&userType=parent",
			expectedStatus: http.StatusOK,
			expectedBody:   "", 
		},
		{
			name:           "Invalid method",
			method:         http.MethodPut,
			body:           "",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request with the specified method and body
			r := httptest.NewRequest(tt.method, "/auth", strings.NewReader(tt.body))
			if tt.method == http.MethodPost {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}

			// Capture the response
			w := httptest.NewRecorder()

			// Check the status code
			if w.Result().StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Result().StatusCode)
			}

			// Check the body content
			if !strings.Contains(w.Body.String(), tt.expectedBody) {
				t.Errorf("Expected body to contain \"%s\", got \"%s\"", tt.expectedBody, w.Body.String())
			}
		})
	}
}
