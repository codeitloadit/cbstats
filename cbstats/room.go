package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "http://chaturbate.com/affiliates/api/onlinerooms/?format=json&wm=qBlp5"

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

// Rooms stores the room data retrieved from the web api.
var Rooms []room

// GetRoomsData fetches the JSON data from the web api and unmarshals it into the Rooms variable.
func GetRoomsData() {
	resp, err := http.Get(url)
	HandleError(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleError(err)

	err = json.Unmarshal(body, &Rooms)
	HandleError(err)
}
