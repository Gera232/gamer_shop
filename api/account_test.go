package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccounts(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(getAccounts))
	resp, err := http.Get(server.URL)
	if err != nil {
		test.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		test.Errorf("expected 200 but got %d", resp.StatusCode)
	}
}

/* func TestGetAccountsHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/GetAccounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder to record the response.
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getAccounts)

	// Call the handler directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(recorder, req)

	// Check the status code is what we expect (HTTP 200).
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
} */
