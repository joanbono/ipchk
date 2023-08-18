# ipchk

Simple IP checker

## Usage

```sh
$ ipchk -version

     📟  ipchk v1.0.0
```

```sh
# Returns information about your own IP
$ ipchk 

# Returns information about the IP 8.8.8.8
$ ipchk -ip 8.8.8.8
{
    "as": "AS15169 Google LLC",
    "city": "Ashburn",
    "country": "United States",
    "countryCode": "US",
    "isp": "Google LLC",
    "lat": 39.03,
    "lon": -77.5,
    "org": "Google Public DNS",
    "query": "8.8.8.8",
    "region": "VA",
    "regionName": "Virginia",
    "status": "success",
    "timezone": "America/New_York",
    "zip": "20149"
}

# Returns information about the IP 8.8.8.8 using boring output
$ ipchk -ip 8.8.8.8 -nc
{"status":"success","country":"United States","countryCode":"US","region":"VA","regionName":"Virginia","city":"Ashburn","zip":"20149","lat":39.03,"lon":-77.5,"timezone":"America/New_York","isp":"Google LLC","org":"Google Public DNS","as":"AS15169 Google LLC","query":"8.8.8.8"}
```