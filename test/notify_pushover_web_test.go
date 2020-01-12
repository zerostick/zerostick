package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	zsweb "github.com/zerostick/zerostick/daemon/web"
	//_ "github.com/zerostick/zerostick/daemon"
)

func TestPushoverGetEmptyWeb(t *testing.T) {
	req, err := http.NewRequest("GET", "/notifications/provider/pushover", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushoverConfig)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"user_key":"","app_key":"","enabled":false}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPushoverConfigSet(t *testing.T) {
	var jsonStr = []byte(`{"user_key":"testuserkey","app_key":"testappkey","enabled":true}`)

	req, err := http.NewRequest("POST", "/notifications/provider/pushover", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushoverConfigSet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"user_key":"testuserkey","app_key":"testappkey","enabled":true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPushoverGetWeb(t *testing.T) {
	req, err := http.NewRequest("GET", "/notifications/provider/pushover", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushoverConfig)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"user_key":"testuserkey","app_key":"testappkey","enabled":true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPushoverDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/notifications/provider/pushover", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushoverConfigDelete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
