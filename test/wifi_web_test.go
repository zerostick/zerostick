package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	zsweb "github.com/zerostick/zerostick/daemon/web"
	//_ "github.com/zerostick/zerostick/daemon"
)

func TestWifiListWeb(t *testing.T) {
	req, err := http.NewRequest("GET", "/wifilist", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.Wifilist)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"AirExtreme":{"bssid":"80:2a:a8:c2:b5:c1","frequency":"5640","signal_level":"-81","flags":"[WPA2-PSK-CCMP][ESS]","ssid":"AirExtreme"},"AirKids":{"bssid":"92:2a:a8:c2:b5:c1","frequency":"5640","signal_level":"-82","flags":"[WPA2-PSK-CCMP][ESS]","ssid":"AirKids"},"Hans":{"bssid":"dc:a4:ca:ba:f5:28","frequency":"2437","signal_level":"-73","flags":"[WPA2-PSK-CCMP][ESS]","ssid":"Hans"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestWifiGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/wifi", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.WifiGetEntries)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ssid":"flaf","encrypted_password":"21ad66ddf9a61afa2a66d9cf233c722e3993b2dd361b5ca1c3456dd7ea9d8ff4","priority":1,"use_for_sync":false},{"ssid":"flaf22","encrypted_password":"6abd60875676ec4c046945a7773f09ce7b8d49b219a514479a49d50e85b74629","priority":22,"use_for_sync":false}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestWifiAddEntry(t *testing.T) {
	var jsonStr = []byte(`{"ssid":"flaf128","encrypted_password":"225e95c13631d1d99be9c51db13b714d26bde19d0d84851bf99a4bb2a4e2478da","priority":128,"use_for_sync":false}`)

	req, err := http.NewRequest("POST", "/wifi", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.WifiGetEntries)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ssid":"flaf","encrypted_password":"21ad66ddf9a61afa2a66d9cf233c722e3993b2dd361b5ca1c3456dd7ea9d8ff4","priority":1,"use_for_sync":false},{"ssid":"flaf22","encrypted_password":"6abd60875676ec4c046945a7773f09ce7b8d49b219a514479a49d50e85b74629","priority":22,"use_for_sync":false}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestWifiDeleteEntry(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/wifi/flaf", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(zsweb.WifiDeleteEntry)
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
