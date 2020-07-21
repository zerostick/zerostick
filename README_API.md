# API Endpoints:

## /wifilist
GET: List of scanned SSIDs
{
    wifi: ["ap2", "ap1"]
}

## /wifi/<id?>
GET: SSID/SSIDs know
POST: Add/update SSID
DELETE: Deletes :id

{ "SSID": "AP1", "password": "flaf", "priority": 10, "useforsync": false }

## /nabto/
GET: Get DeviceID
POST: Save deviceID and key
{ "deviceid": "devid", "key": "nabtokey" }

## /nabto/delete_acl
DELETE: Deletes Nabto ACL file

## /events
GET: Get all events metadata
