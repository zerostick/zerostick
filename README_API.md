# API Endpoints:

## /wifilist
### GET: List of scanned SSIDs
Returns:
```json
{
    "AirExtreme": {
        "bssid":"80:2a:a8:c2:b5:c1",
        "frequency":"5640",
        "signal_level":"-81",
        "flags":"[WPA2-PSK-CCMP][ESS]",
        "ssid":"AirExtreme"
    },
    "Tesla Guest": {
        "bssid":"92:2a:a8:c2:b5:c2",
        "frequency":"5640",
        "signal_level":"-82",
        "flags":"[WPA2-PSK-CCMP][ESS]",
        "ssid":"Tesla Guest"
    }
}
```

## /wifi/
### GET: SSID/SSIDs known
```bash
curl http://localhost:8081/wifi
```

Returns:
```json
[
    {
        "ssid": "flaf128",
        "encrypted_password": "25e95c13631d1d99be9c51db13b714d26bde19d0d84851bf99a4bb2a4e2478da",
        "priority": 128,
        "use_for_sync": false
    },
    {
        "ssid": "flaf22",
        "encrypted_password": "9816f2e2f268fd66b600e58ae5a3ce02cdfff0aaa57750e918e77e772fb0871a",
        "priority": 22,
        "use_for_sync": false
    }
]
```

### POST: Add/update SSID
```bash
curl -d '{
    "ssid": "flaf22",
    "password": "testhest",
    "priority": 22,
    "use_for_sync": false
}' -H "Content-Type: application/json" "http://localhost:8081/wifi"
```

Returns:
```json
{"ssid":"flaf22","encrypted_password":"9816f2e2f268fd66b600e58ae5a3ce02cdfff0aaa57750e918e77e772fb0871a","priority":22,"use_for_sync":false}
```

### DELETE: Deletes :id
```bash

```

Returns:
```json

```

## /nabto/
GET: Get DeviceID
POST: Save deviceID and key
{ "deviceid": "devid", "key": "nabtokey" }

## /nabto/delete_acl
DELETE: Deletes Nabto ACL file

## /events
GET: Get all events metadata
