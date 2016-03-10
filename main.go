package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const port = "8080"
const seperator = "-------------------------"

var rooms []room

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

	broadcasterCounts := make(map[string]uint)
	viewerCounts := make(map[string]uint)
	typeCounts := make(map[string]uint)
	tagCounts := make(map[string]uint)
	var tagCount uint
	var hdCount uint
	var newCount uint
	var ageSum uint
	var followersSum uint
	var secondsSum uint
	for _, room := range rooms {
		broadcasterCounts[room.Gender]++
		broadcasterCounts["a"]++
		viewerCounts[room.Gender] += room.Viewers
		viewerCounts["a"] += room.Viewers
		if len(room.Tags) > 0 {
			tagCount++
			for _, tag := range room.Tags {
				tagCounts[tag]++
			}
		}
		typeCounts[room.Type]++
		if room.IsHD {
			hdCount++
		}
		if room.IsNew {
			newCount++
		}
		ageSum += room.Age
		followersSum += room.Followers
		secondsSum += room.Seconds
	}

	fmt.Fprintln(w, "CB Stats")
	fmt.Fprintln(w, seperator)
	fmt.Fprintln(w, "Broadcasters:", broadcasterCounts["a"])
	fmt.Fprintln(w, "Females:", broadcasterCounts["f"])
	fmt.Fprintln(w, "Males:", broadcasterCounts["m"])
	fmt.Fprintln(w, "Couples:", broadcasterCounts["c"])
	fmt.Fprintln(w, "Trans:", broadcasterCounts["s"])
	fmt.Fprintln(w, seperator)
	fmt.Fprintln(w, "Viewers:", viewerCounts["a"])
	fmt.Fprintln(w, "Females:", viewerCounts["f"])
	fmt.Fprintln(w, "Males:", viewerCounts["m"])
	fmt.Fprintln(w, "Couples:", viewerCounts["c"])
	fmt.Fprintln(w, "Trans:", viewerCounts["s"])
	fmt.Fprintln(w, seperator)
	fmt.Fprintln(w, "Unique Tags:", len(tagCounts))
	fmt.Fprintln(w, "Rooms With Tags:", tagCount)
	fmt.Fprintln(w, "Public Rooms:", typeCounts["public"])
	fmt.Fprintln(w, "Private Rooms:", typeCounts["private"])
	fmt.Fprintln(w, "Group Rooms:", typeCounts["group"])
	fmt.Fprintln(w, "Away Rooms:", typeCounts["away"])
	fmt.Fprintln(w, "HD Rooms:", hdCount)
	fmt.Fprintln(w, "New Rooms:", newCount)
	fmt.Fprintln(w, seperator)
	fmt.Fprintln(w, "Average Minutes:", secondsSum/broadcasterCounts["a"]/60.0)
	fmt.Fprintln(w, "Average Age:", ageSum/broadcasterCounts["a"])
	fmt.Fprintln(w, "Average Viewers:", viewerCounts["a"]/broadcasterCounts["a"])
	fmt.Fprintln(w, "Average Followers:", followersSum/broadcasterCounts["a"])

}

func main() {
	fmt.Println("Listening on port", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
