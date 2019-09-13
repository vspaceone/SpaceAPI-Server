# SpaceAPI-Server
SpaceAPI Server written in Go

This server can return a valid SpaceAPI-string in version 13 as specified 
[here](https://spacedirectory.org/pages/docs.html]) and provides further API endpoints for modification of the returned values.

## Features

### Implemented

*  Return of valid SpaceAPI strings
*  Modification of SpaceAPI
    *  state
    *  sensors
        *  temperature
        *  humidity
*  Persistence using a database (sqlite tested)
*  Simple token authentication of modification-requests
*  Static override of some values (for example to set space details like name, location, url etc.)
*  Dockerfile

### Planned

*  Support for the whole SpaceAPI (with all specified fields) including modification
*  Support for HTTPS (use with reverse proxies for now)

## Running

* Create a config and an override file in /srv/spaceapi-server like in this example (with your values)

config.json
```
{
    "db":"sqlite3",
    "dbconnection":"data/lite.db",
    "port":8080,
    "debug":false
}
```

override.json
```
{
    "API":   "0.13",
    "Space": "vspace.one",
    "Logo":  "https://vspace.one/pic/logo_vspaceone.svg",
    "URL":   "https://vspace.one",
    "location": {
        "Address": "Wilhelm-Binder-Str. 19, 78048 VS-Villingen, Germany",
        "Lat":     48.065003,
        "Lon":     8.456495
    },
    "contact": {
        "Phone":   "+49 221 596196638",
        "Email":   "info@vspace.one",
        "Twitter": "@vspace.one"
    },
    "IssueReportChannels": [
        "email",
        "twitter"
    ]
}
```

* `docker run --name spaceapi-server -v /srv/spaceapi-server/:/go/src/github.com/vspaceone/SpaceAPI-Server/data vspaceone/spaceapi-server`
* Token and database files should be created automagically

## API

### Getting SpaceAPI string

*GET on /spaceapi  
GET on /spaceapi.json*

Returns the whole SpaceAPI string

### Setting SpaceAPI values

*POST on /spaceapi*

Makes it possible to send data similar to the SpaceAPI string to set 
specific values (for now only setting of state.open, sensors.temperature and sensors.humidity is possible).

**Note** that setting these values is only possible if the right token is specified in Header as `X-Auth-Token`. The token you need to specify is generated at first start of this application. When specifying a wrong token or none the server will respond with status 401.

Examples for POST payload:

**Setting state.open**
```
{
    "state": {
        "open": false,
        "lastchange": 1519502622
    }
}
```
**Setting sensors.temperature and sensors.humidity**
```
{
    "sensors": {
        "temperature": [
            {
                "value": 25,
                "unit": "°C",
                "location": "Maschinenraum"
            },
            {
                "value": 22,
                "unit": "°C",
                "location": "Brücke"
            }
        ],
        "humidity": [
            {
                "value": 50,
                "unit": "%",
                "location": "Brücke"
            },
            {
                "value": 40,
                "unit": "%",
                "location": "Maschinenraum"
            }
        ]
    }
}
```

### Getting past sensor data

Not possible anymore. 
This server only saves the current state in its database and presents it as a SpaceAPI.
Other tools can be used for monitoring, logging and analyzing changes to the API.