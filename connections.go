package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
)

type query struct {
	Connections []Connections `json:"connections"`
}
type Connections struct {
	Duration  string `json:"duration"`
	From      Location
	To        Location
	Transfers int `json:"transfers"`
}
type Location struct {
	Departure string
	Platform  string
	Delay     int
	Station   Station
}
type Station struct {
	Name string
}

type Options struct {
	From     string `required:"true"  long:"from" description:"declare the name of the departing city"`
	To       string `required:"true" long:"to" description:"declare the name of the destination city"`
	Time     string `required:"true"  long:"time" description:"declare time of departure, format [hh:mm]"`
	Date     string `required:"true"  long:"date" description:"declare date of departure, format [YYYY-MM-DD]"`
	ArriveBy bool   ` long:"arriveby" description:"mark declared time and date as the arrival ones, usage --arriveby"`
	Direct   bool   ` long:"direct" description:"query only direct connections, usage --direct"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	allConnections := opts.get()
	/***PRINT FASTEST***/
	fmt.Println("******FASTEST CONNECTION******")
	print(calculateFastestConn(allConnections))
	fmt.Println()
	/***PRINT EARLIEST***/
	fmt.Println("******EARLIEST CONNECTION******")
	print(calculateEarliest(allConnections))

}
func print(c Connections) {
	fmt.Printf("FROM: %v\n", c.From.Station.Name)
	fmt.Printf("TO: %v\n", c.To.Station.Name)
	fmt.Printf("DEPARTURE TIME: %v\n", c.From.Departure)
	fmt.Printf("DURATION: %v\n", c.Duration)
	fmt.Printf("TRANSFERS: %v\n", c.Transfers)
	fmt.Printf("DELAY: %v\n", c.From.Delay)
	fmt.Printf("DEPARTURE PLATFORM: %v\n", c.From.Platform)
	fmt.Printf("ARRIVAL PLATFORM: %v\n", c.To.Platform)

}
func calculateEarliest(queryresults query) Connections {
	layout := "2006-01-02T15:04:05.000Z"
	earliestConnection := queryresults.Connections[0]
	timeofEarliestConn, _ := time.Parse(layout, earliestConnection.From.Departure)
	for _, v := range queryresults.Connections {
		time, _ := time.Parse(layout, v.From.Departure)
		if time.Before(timeofEarliestConn) {
			earliestConnection = v
		}

	}
	return earliestConnection
}

func calculateFastestConn(queryresults query) Connections {
	fastest := queryresults.Connections[0]
	for _, v := range queryresults.Connections {
		if isFaster(v.Duration, fastest.Duration) {
			fastest = v
		}
	}
	return fastest
}

func isFaster(string1, string2 string) bool {
	d1 := formatDuration(string1)
	d2 := formatDuration(string2)
	for i := range d1 {
		if d1[i]-d2[i] > 0 {
			return false
		}
	}
	return true
}

func formatDuration(duration string) []int {
	split := strings.FieldsFunc(duration, split)
	d := make([]int, len(split))
	for i, v := range split {
		d[i], _ = strconv.Atoi(v)

	}
	return d
}

func split(r rune) bool {
	return r == ':' || r == 'd'
}
func (opt Options) get() query {
	var m query
	URL := "http://transport.opendata.ch/v1/connections" +
		"?from=" + opt.From +
		"&to=" + opt.To +
		"&time=" + opt.Time +
		"&date=" + opt.Date
	if opt.ArriveBy {
		URL += "&arrive=1"
	}
	if opt.Direct {
		URL += "&direct=1"
	}

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("could not query API")

	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("failed reading response body")
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(data, &m); err != nil {
		log.Fatalf("failed unmarshalling %v", err)
	}
	return m
}
