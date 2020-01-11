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

### DELETE: Deletes /wifi/:id
```bash
curl -X DELETE http://localhost:8081/wifi/flaf22
```

Returns:
*Nothing*

## /nabto/
### GET: Get DeviceID
```bash
curl "http://localhost:8081/nabto"
```
Return:
```json
{
    "deviceid": "devid",
    "devicekey": "nabtokey"
}
```

### POST: Save deviceID and key
```bash
curl -v -d '{ 
    "deviceid": "devid", 
    "devicekey": "nabtokey"
}' -H "Content-Type: application/json" "http://localhost:8081/nabto"
```
Returns:
```json
{
    "deviceid": "devid",
    "devicekey":"nabtokey"
}
```


## /nabto/delete_acl
### DELETE: Deletes Nabto ACL file
```bash
curl -X DELETE "http://localhost:8081/nabto/delete_acl"
```
Returns:
```
{}
```

## /events/:type
### GET: Get all events metadata
Type is usualy one of "RecentClips", "SavedClips" and "SentryClips".

```bash
curl "http://localhost:8081/events/SavedClips"
```
Returns:
```json
{
   "2019-07-23_17-22-24" : [
      {
         "file_size" : 29407105,
         "file_time" : "2019-07-23T17:11:41Z",
         "event" : "2019-07-23_17-22-24",
         "filename" : "2019-07-23_17-11-41-front.mp4",
         "camera" : "front",
         "thumbnail_path" : "/TeslaCam/SavedClips/2019-07-23_17-22-24/2019-07-23_17-11-41-front.mp4.jpg",
         "id" : "481543f3-8afe-4b4f-ad5b-642dae1c2f03"
      },
      {
         "file_size" : 21898242,
         "file_time" : "2019-07-23T17:11:41Z",
         "filename" : "2019-07-23_17-11-41-left_repeater.mp4",
         "event" : "2019-07-23_17-22-24",
         "camera" : "left_repeater",
         "thumbnail_path" : "/TeslaCam/SavedClips/2019-07-23_17-22-24/2019-07-23_17-11-41-left_repeater.mp4.jpg",
         "id" : "e1c7fdcc-6f62-4a9f-be67-c5bb568b2afc"
      },
```
etc.

## /video/:id
### GET: Get a video given the Id in streaming friendly format
```curl
curl -o TeslaVideo.mp4 http://localhost:8081/video/481543f3-8afe-4b4f-ad5b-642dae1c2f03
```
Returns:
- A MP4 video file
