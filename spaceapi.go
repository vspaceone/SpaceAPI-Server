package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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
}

type humidity struct {
	Value    float32 `json:"value"`
	Unit     string  `json:"unit"`
	Location string  `json:"location"`
}

var override status

func (st *status) set(st2 status) {
	st.Space = st2.Space
	st.Logo = st2.Logo
	st.URL = st2.URL
	st.location = st2.location
	st.contact = st2.contact
	st.IssueReportChannels = st2.IssueReportChannels
}

func loadOverride() {
	var st status

	dat, err := ioutil.ReadFile("data/override.json")
	if err != nil {
		log.Fatal("Couldn't read override file")
		return
	}

	err = json.Unmarshal(dat, &st)
	if err != nil {
		log.Fatal("Couldn't unmarshal override file")
		return
	}

	override = st
}

func bbuf(r io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.Bytes()
}

func createAPIString(st state, temps []temperature, hums []humidity) string {
	loadOverride()

	s := status{
		API:   "0.13",
		Space: "",
		Logo:  "",
		URL:   "",
		location: location{
			Address: "",
			Lat:     0,
			Lon:     0,
		},
		contact: contact{
			Phone:   "",
			Email:   "",
			Twitter: "",
		},
		state: state{
			Open:       st.Open,
			LastChange: st.LastChange,
		},
		IssueReportChannels: []string{},
		sensors: sensors{
			Temperature: temps,
			Humidity:    hums,
		},
	}

	s.set(override)

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
