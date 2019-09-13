package main

/*type data struct {
	Request  string `json:"request"`
	Location string `json:"location"`
	Filter   string `json:"filter"`
	Since    uint64 `json:"since"`
}

// filters
var minuteFilter = "%Y-%m-%d %H:%M"
var hourFilter = "%Y-%m-%d %H"
var dayFilter = "%Y-%m-%d"
var monthFilter = "%Y-%m"
var yearFilter = "%Y"

func marshalResponse(o interface{}) string {
	b, _ := json.MarshalIndent(o, "", "    ")
	return string(b)
}

func createSensorsResponse(buf []byte) string {
	var body data
	err := json.Unmarshal(buf, &body)
	if err != nil {
		panic(err)
	}

	var filter string
	switch body.Filter {
	case "minute":
		filter = minuteFilter
	case "hour":
		filter = hourFilter
	case "day":
		filter = dayFilter
	case "month":
		filter = monthFilter
	case "year":
		filter = yearFilter
	default:
		filter = minuteFilter
	}

	switch body.Request {
	case "temperature":
		var temp []temperature
		db.Find(&temp, "location = ? AND changed > ? "+
			"AND changed IN( "+
			"SELECT MIN(t2.changed) FROM temperatures t2 "+
			"WHERE location = ? "+
			"GROUP BY strftime(?,t2.changed, 'unixepoch') )", body.Location, body.Since,
			body.Location, filter).
			Order("changed DESC")
		return marshalResponse(temp)
	case "humidity":
		var hum []humidity
		db.Find(&hum, "location = ? AND changed > ? "+
			"AND changed IN( "+
			"SELECT MIN(t2.changed) FROM humidities t2 "+
			"WHERE location = ? "+
			"GROUP BY strftime(?,t2.changed, 'unixepoch') )", body.Location, body.Since,
			body.Location, filter).
			Order("changed DESC")
		return marshalResponse(hum)
	case "all":
		var temp []temperature
		db.Find(&temp, "location = ? AND changed > ? "+
			"AND changed IN( "+
			"SELECT MIN(t2.changed) FROM temperatures t2 "+
			"WHERE location = ? "+
			"GROUP BY strftime(?,t2.changed, 'unixepoch') )", body.Location, body.Since,
			body.Location, filter).
			Order("changed DESC")

		var hum []humidity
		db.Find(&hum, "location = ? AND changed > ? "+
			"AND changed IN( "+
			"SELECT MIN(t2.changed) FROM humidities t2 "+
			"WHERE location = ? "+
			"GROUP BY strftime(?,t2.changed, 'unixepoch') )", body.Location, body.Since,
			body.Location, filter).
			Order("changed DESC")
		return "{\"temperature\":" + marshalResponse(hum) + ",\"humidity\":" + marshalResponse(temp) + "}"
	}

	return "[]"
}
*/
