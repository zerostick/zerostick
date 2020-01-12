package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	zsweb "github.com/zerostick/zerostick/daemon/web"
	//_ "github.com/zerostick/zerostick/daemon"
)

func TestPushbulletGetEmptyWeb(t *testing.T) {
	req, err := http.NewRequest("GET", "/notifications/provider/pushbullet", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushbulletConfig)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"api_key":"","enabled":false}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPushbulletConfigSet(t *testing.T) {
	var jsonStr = []byte(`{"api_key":"testkey","enabled":true}`)

	req, err := http.NewRequest("POST", "/notifications/provider/pushbullet", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushbulletConfigSet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"api_key":"testkey","enabled":true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPushbulletGetWeb(t *testing.T) {
	req, err := http.NewRequest("GET", "/notifications/provider/pushbullet", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushbulletConfig)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"api_key":"testkey","enabled":true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPushbulletDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/notifications/provider/pushbullet", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NotificationPushbulletConfigDelete)
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
