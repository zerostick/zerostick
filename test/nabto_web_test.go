package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	zsweb "github.com/zerostick/zerostick/daemon/web"
	//_ "github.com/zerostick/zerostick/daemon"
)

func TestNabtoGetDeviceIDEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/nabto", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NabtoConfig)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"deviceid":"","devicekey":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestNabtoSetup(t *testing.T) {
	var jsonStr = []byte(`{ 
		"deviceid": "devid", 
		"devicekey": "nabtokey"
	}`)
	req, err := http.NewRequest("POST", "/nabto", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NabtoSetup)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"deviceid":"devid","devicekey":"nabtokey"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestNabtoGetConfig(t *testing.T) {
	req, err := http.NewRequest("GET", "/nabto", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.NabtoConfig)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"deviceid":"devid","devicekey":"nabtokey"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
