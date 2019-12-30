package main

import (
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
	expected := `{"AirExtreme":{"bssid":"80:2a:a8:c2:b5:c1","frequency":"5640","signal_level":"-81","flags":"[WPA2-PSK-CCMP][ESS]","ssid":"AirExtreme"},"AirKids":{"bssid":"92:2a:a8:c2:b5:c1","frequency":"5640","signal_level":"-82","flags":"[WPA2-PSK-CCMP][ESS]","ssid":"AirKids"},"Hans":{"bssid":"dc:a4:ca:ba:f5:28","frequency":"2437","signal_level":"-73","flags":"[WPA2-PSK-CCMP][ESS]","ssid":"Hans"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
