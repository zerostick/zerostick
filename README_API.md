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

## /events
### GET: Get all events metadata
```bash
curl "http://localhost:8081/events"
```
Returns:
```json
{
    "SavedClips": {
        "2019-07-22_18-12-23": [
            {
                "Id": "682f9be5-63f6-4e65-af90-2b03982dfebe",
                "Name": "2019-07-22_18-01-55-front-fast.mp4",
                "ThumbnailFile": "/var/folders/r6/jbkh0zhs65x5s1l43qc7cz940000gn/T/ZeroStick/TeslaCam/SavedClips/2019-07-22_18-12-23/2019-07-22_18-01-55-front-fast.mp4.jpg",
                "ThumbnailRelPath": "/TeslaCam/SavedClips/2019-07-22_18-12-23/2019-07-22_18-01-55-front-fast.mp4.jpg",
                "FullPath": "/Users/hilli/TeslaCam/SavedClips/2019-07-22_18-12-23",
                "Event": "2019-07-22_18-12-23",
                "EventType": "SavedClips",
                "EventTime": "2019-07-22T18:01:55Z",
                "EventCam": "front-fast",
                "Size": 29368368
            },
            {
                "Id": "d49e9baf-ef9b-4e51-b59e-5c78dd3f82e4",
                "Name": "2019-07-22_18-01-55-front-ordered.mp4",
                "ThumbnailFile": "/var/folders/r6/jbkh0zhs65x5s1l43qc7cz940000gn/T/ZeroStick/TeslaCam/SavedClips/2019-07-22_18-12-23/2019-07-22_18-01-55-front-ordered.mp4.jpg",
                "ThumbnailRelPath": "/TeslaCam/SavedClips/2019-07-22_18-12-23/2019-07-22_18-01-55-front-ordered.mp4.jpg",
                "FullPath": "/Users/hilli/TeslaCam/SavedClips/2019-07-22_18-12-23",
                "Event": "2019-07-22_18-12-23",
                "EventType": "SavedClips",
                "EventTime": "2019-07-22T18:01:55Z",
                "EventCam": "front-ordered",
                "Size": 20399084
            },
```
etc.

## /video/:id
### GET: Get a video given the Id in streaming friendly format
```curl
curl -o TeslaVideo.mp4 http://localhost:8081/video/682f9be5-63f6-4e65-af90-2b03982dfebe
```
Returns:
- A MP4 video file
