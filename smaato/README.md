

http POST http://localhost:8000/openrtb2/auction < request.json

HTTP/1.1 200 OK
Cache-Control: no-cache, no-store, must-revalidate
Content-Length: 225
Content-Type: application/json
Date: Sat, 16 May 2020 15:42:13 GMT
Expires: 0
Pragma: no-cache
Vary: Origin

{
    "ext": {
        "errors": {
            "smaato": [
                {
                    "code": 5,
                    "message": "The adapter failed to generate any bid requests, but also failed to generate an
error explaining why"
                }
            ]
        },
        "responsetimemillis": {
            "smaato": 0
        },
        "tmaxrequest": 500
    },
    "id": "324234234"
}