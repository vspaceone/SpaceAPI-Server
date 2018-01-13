package main

import (
	"bytes"
	"encoding/json"
	"io"
)

type status struct {
	API                 string `json:"api"`
	Space               string `json:"space"`
	Logo                string `json:"logo"`
	URL                 string `json:"url"`
	location            `json:"location"`
	contact             `json:"contact"`
	IssueReportChannels []string `json:"issue_report_channels"`
	state               `json:"state"`
	sensors             `json:"sensors"`
}

type location struct {
	Address string  `json:"address"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

type contact struct {
	Phone   string `json:"phone"`
	Twitter string `json:"twitter"`
	Email   string `json:"email"`
}

type state struct {
	Open       bool  `json:"open"`
	LastChange int64 `json:"lastchange"`
}

type sensors struct {
	Temperature []temperature `json:"temperature"`
	Humidity    []humidity    `json:"humidity"`
}

type temperature struct {
	Value    float32 `json:"value"`
	Unit     string  `json:"unit"`
	Location string  `json:"location"`
	Changed  int64   `json:"-"`
}

type humidity struct {
	Value    float32 `json:"value"`
	Unit     string  `json:"unit"`
	Location string  `json:"location"`
	Changed  int64   `json:"-"`
}

func bbuf(r io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.Bytes()
}

func createAPIString(st state, temps []temperature, hums []humidity) string {
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
			Open:       st.Open,
			LastChange: st.LastChange,
		},
		IssueReportChannels: []string{
			"email",
			"twitter",
		},
		sensors: sensors{
			Temperature: temps,
			Humidity:    hums,
		},
	}
	b, _ := json.MarshalIndent(s, "", "    ")
	return string(b)
}

func getState(buf []byte) (s state, exists bool) {

	var t map[string]interface{}
	err := json.Unmarshal(buf, &t)
	if err != nil {
		panic(err)
	}

	// When no "state" given in data
	_, ok := t["state"]
	if !ok {
		exists = false
		return
	}

	// When state is given return it
	var data status
	err = json.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	return data.state, true
}

func getSensors(buf []byte) sensors {

	// When state is given return it
	var data status
	err := json.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	return data.sensors
}
