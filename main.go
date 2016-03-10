package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

const port = "8080"

type room struct {
	Username  string   `json:"username"`
	Gender    string   `json:"gender"`
	Viewers   uint     `json:"num_users"`
	Followers uint     `json:"num_followers"`
	Type      string   `json:"current_show"`
	IsHD      bool     `json:"is_hd"`
	IsNew     bool     `json:"is_new"`
	Age       uint     `json:"age"`
	Seconds   uint     `json:"seconds_online"`
	Tags      []string `json:"tags"`
}

type statsStruct struct {
	BroadcasterCounts map[string]uint
	ViewerCounts      map[string]uint
	TagCounts         map[string]uint
	RoomsWithTags     uint
	TypeCounts        map[string]uint
	HDRooms           uint
	NewRooms          uint
	TotalSeconds      uint
	TotalAge          uint
	TotalFollowers    uint
}

func (s statsStruct) AverageMinutes() uint {
	return s.TotalSeconds / s.BroadcasterCounts["a"] / 60
}

func (s statsStruct) AverageAge() uint {
	return s.TotalAge / s.BroadcasterCounts["a"]
}

func (s statsStruct) AverageViewers() uint {
	return s.ViewerCounts["a"] / s.BroadcasterCounts["a"]
}

func (s statsStruct) AverageFollowers() uint {
	return s.TotalFollowers / s.BroadcasterCounts["a"]
}

var rooms []room
var stats statsStruct

func getRooms() {
	resp, err := http.Get("http://chaturbate.com/affiliates/api/onlinerooms/?format=json&wm=qBlp5")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &rooms)
	if err != nil {
		panic(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	getRooms()

	stats.BroadcasterCounts = make(map[string]uint)
	stats.ViewerCounts = make(map[string]uint)
	stats.TagCounts = make(map[string]uint)
	stats.TypeCounts = make(map[string]uint)

	for _, room := range rooms {
		stats.BroadcasterCounts[room.Gender]++
		stats.BroadcasterCounts["a"]++
		stats.ViewerCounts[room.Gender] += room.Viewers
		stats.ViewerCounts["a"] += room.Viewers
		if len(room.Tags) > 0 {
			stats.RoomsWithTags++
			for _, tag := range room.Tags {
				stats.TagCounts[tag]++
			}
		}
		stats.TypeCounts[room.Type]++
		if room.IsHD {
			stats.HDRooms++
		}
		if room.IsNew {
			stats.NewRooms++
		}
		stats.TotalAge += room.Age
		stats.TotalFollowers += room.Followers
		stats.TotalSeconds += room.Seconds
	}

	tmpl, err := template.New("cbstats").Parse(Text)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, stats)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Listening on port", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
