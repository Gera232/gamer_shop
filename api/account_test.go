package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(getAccounts))
	defer s.Close()

	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		t.Error(err)
	}
	defer req.Body.Close()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Error(res.StatusCode)
	}
}
