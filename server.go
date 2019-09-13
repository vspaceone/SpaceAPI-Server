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
	//http.HandleFunc("/spaceapi/sensors", sensorsEp)
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

	var dbState dbState
	var dbTemps []dbTemperature
	var dbHums []dbHumidity
	db.First(&dbState)
	db.Table("db_temperatures").Scan(&dbTemps)
	db.Table("db_humidities").Scan(&dbHums)

	var state state
	state.LastChange = dbState.LastChange
	state.Open = dbState.Open

	var temps []temperature
	for _, elem := range dbTemps {
		var tmp temperature
		tmp.Location = elem.Location
		tmp.Unit = elem.Unit
		tmp.Value = elem.Value
		temps = append(temps, tmp)
	}

	var hums []humidity
	for _, elem := range dbHums {
		var hum humidity
		hum.Location = elem.Location
		hum.Unit = elem.Unit
		hum.Value = elem.Value
		hums = append(hums, hum)
	}

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

			var state dbState
			db.FirstOrCreate(&state, dbState{StateID: 1})

			state.Open = st.Open
			state.LastChange = time.Now().Unix()

			db.Save(&state)
		}

		sensors := getSensors(buf)

		for _, elem := range sensors.Temperature {
			log.Printf("Temperature %s: %.2f %s", elem.Location, elem.Value, elem.Unit)

			var tmp dbTemperature
			db.FirstOrCreate(&tmp, dbTemperature{Location: elem.Location})

			tmp.Changed = time.Now().Unix()
			tmp.Location = elem.Location
			tmp.Unit = elem.Unit
			tmp.Value = elem.Value

			db.Save(&tmp)
		}

		for _, elem := range sensors.Humidity {
			log.Printf("Humidity %s: %.2f %s", elem.Location, elem.Value, elem.Unit)

			var hum dbHumidity
			db.FirstOrCreate(&hum, dbHumidity{Location: elem.Location})

			hum.Changed = time.Now().Unix()
			hum.Location = elem.Location
			hum.Unit = elem.Unit
			hum.Value = elem.Value

			db.Save(&hum)
		}

	}
}

/*func sensorsEp(w http.ResponseWriter, r *http.Request) {
	// Only Post method allowed
	if r.Method != http.MethodPost {
		log.Println("not POST")
		w.WriteHeader(404)
		return
	}
	buf := bbuf(r.Body)
	fmt.Fprint(w, createSensorsResponse(buf))
}
*/
