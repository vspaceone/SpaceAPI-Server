package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var token = ""

func serve(port int16) {

	loadToken()
	fmt.Println("Serving SpaceAPI server on port " + fmt.Sprintf("%d", port))

	http.HandleFunc("/spaceapi.json", spaceapi)
	http.HandleFunc("/spaceapi", spaceapiEp)
	http.HandleFunc("/spaceapi/sensors", sensorsEp)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", port), nil))
}

func loadToken() {
	dat, err := ioutil.ReadFile("data/token")
	if err != nil {
		log.Println("Couldn't read token file. Creating one under data/token")
		generateToken()
		dat, err = ioutil.ReadFile("data/token")
		if err != nil {
			log.Fatalln("Token creation failed.")
			panic("Exiting")
		}
	}
	token = string(dat)
	fmt.Println("Token:\n" + string(dat))
}

func spaceapi(w http.ResponseWriter, r *http.Request) {

	// Allow access from all locations
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var state state
	var temps []temperature
	var hums []humidity

	db.Find(&state).Order("state.LastChange DESC")

	db.Raw("SELECT * FROM temperatures t1 " +
		"WHERE t1.changed = " +
		"(SELECT MAX(t2.changed) FROM temperatures t2 WHERE t1.location = t2.location);").Scan(&temps)

	db.Raw("SELECT * FROM humidities t1 " +
		"WHERE t1.changed = " +
		"(SELECT MAX(t2.changed) FROM humidities t2 WHERE t1.Location = t2.Location);").Scan(&hums)

	fmt.Fprint(w, createAPIString(state, temps, hums))
}

func spaceapiEp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Just return the SpaceApi string from database
		spaceapi(w, r)
	} else if r.Method == http.MethodPost {
		// Accept changes and write them to a database

		if strings.Compare(r.Header.Get("X-Auth-Token"), token) != 0 {
			w.WriteHeader(401)
			return
		}

		buf := bbuf(r.Body)

		log.Println("Received: " + string(buf))

		st, exists := getState(buf)

		if exists {
			log.Printf("Open: %t", st.Open)
			dbState := state{st.Open, time.Now().Unix()}
			db.Create(&dbState)
		}

		sensors := getSensors(buf)

		for _, elem := range sensors.Temperature {
			log.Printf("Temperature %s: %.2f %s", elem.Location, elem.Value, elem.Unit)
			db.Create(&temperature{elem.Value, elem.Unit, elem.Location, time.Now().Unix()})
		}

		for _, elem := range sensors.Humidity {
			log.Printf("Humidity %s: %.2f %s", elem.Location, elem.Value, elem.Unit)
			db.Create(&humidity{elem.Value, elem.Unit, elem.Location, time.Now().Unix()})
		}

	}
}

func sensorsEp(w http.ResponseWriter, r *http.Request) {
	// Only Post method allowed
	if r.Method != http.MethodPost {
		log.Println("not POST")
		w.WriteHeader(404)
		return
	}
	buf := bbuf(r.Body)
	fmt.Fprint(w, createSensorsResponse(buf))
}
