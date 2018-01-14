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

### Planned

*  Endpoints to get information that is not present in the SpaceAPI string (sensor or state changes over time etc.)
*  Dockerfile
*  Support for the whole SpaceAPI (with all specified fields) includeing modification
*  Support for HTTPS (use with reverse proxies for now)
