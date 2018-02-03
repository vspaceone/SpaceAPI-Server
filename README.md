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

*  Endpoints to get information that is not present in the SpaceAPI string (sensor or state changes over time etc.)
*  Support for the whole SpaceAPI (with all specified fields) includeing modification
*  Support for HTTPS (use with reverse proxies for now)

## Running
For now no docker registry is provided, so the container needs to be built first.

* `docker build -t sapi .`
* Create a config and an override file in /srv/sapi like in this example (with your values)

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

* `docker run --name sapi -v /srv/sapi/:/go/src/github.com/vspaceone/SpaceAPI-Server/data sapi`
* Token and database files should be created automagically