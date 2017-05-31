package main

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
