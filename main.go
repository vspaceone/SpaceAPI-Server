package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var token = ""

func main() {
	if strings.Compare("print", os.Args[1]) == 0 {
		fmt.Println(createAPIString())
	} else if strings.Compare("create-token", os.Args[1]) == 0 {
		token, _ := GenerateRandomString(64)
		fmt.Println("Generated token: " + token)
		ioutil.WriteFile("token", []byte(token), 0622)

	} else if strings.Compare("serve", os.Args[1]) == 0 {
		port := "8080"
		if len(os.Args) > 2 {
			port = os.Args[2]
		}

		loadToken()
		fmt.Println("Serving SpaceAPI server on port " + port)

		http.HandleFunc("/spaceapi.json", spaceapi)
		http.HandleFunc("/set", set)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}

func spaceapi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, createAPIString())
}
func set(w http.ResponseWriter, r *http.Request) {

}

func createAPIString() string {
	s := status{
		API:   "0.13",
		Space: "vspace.one",
		Logo:  "https://vspace.one/pic/logo_vspaceone.svg",
		URL:   "https://vspace.one",
		location: location{
			Address: "Wilhelm-Binder-Str. 19, 78048 VS-Villingen, Germany",
			Lat:     48.065003,
			Lon:     8.456495,
		},
		contact: contact{
			Phone:   "+49 221 596196638",
			Email:   "info@vspace.one",
			Twitter: "@vspace.one",
		},
		state: state{
			Open: false,
		},
		IssueReportChannels: []string{
			"email",
			"twitter",
		},
		sensors: sensors{
			Temperature: []temperature{
				temperature{
					Value:    25.5,
					Unit:     "°C",
					Location: "Maschinenraum",
				},
				temperature{
					Value:    26.6,
					Unit:     "°C",
					Location: "Brücke",
				},
			},
			Humidity: []humidity{
				humidity{
					Value:    50.5,
					Unit:     "%",
					Location: "Maschinenraum",
				},
				humidity{
					Value:    60.6,
					Unit:     "%",
					Location: "Brücke",
				},
			},
		},
	}
	b, _ := json.MarshalIndent(s, "", "    ")
	return string(b)
}

func loadToken() {
	dat, err := ioutil.ReadFile("token")
	if err != nil {
		panic("Couldn't read token file. You can generate a new one with \"SpaceAPI-Server create-token\"\n" + err.Error())
	}
	token = string(dat)
	fmt.Println("Token:\n" + string(dat))
}
