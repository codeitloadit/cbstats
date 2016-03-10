package main

import (
	"html/template"
	"io"
)

const text = `CB Stats
-------------------------
All Broadcasters: {{.BroadcasterCounts.a}}
Female Broadcasters: {{.BroadcasterCounts.f}}
Male Broadcasters: {{.BroadcasterCounts.m}}
Couples Broadcasters: {{.BroadcasterCounts.c}}
Trans Broadcasters: {{.BroadcasterCounts.s}}
-------------------------
All Viewers: {{.ViewerCounts.a}}
Female Viewers: {{.ViewerCounts.f}}
Male Viewers: {{.ViewerCounts.m}}
Couples Viewers: {{.ViewerCounts.c}}
Trans Viewers: {{.ViewerCounts.s}}
-------------------------
Unique Tags: {{.TagCounts | len}}
Rooms With Tags: {{.RoomsWithTags}}
Public Rooms: {{.TypeCounts.public}}
Private Rooms: {{.TypeCounts.private}}
Group Rooms: {{.TypeCounts.group}}
Away Rooms: {{.TypeCounts.away}}
HD Rooms: {{.HDRooms}}
New Rooms: {{.NewRooms}}
-------------------------
Average Minutes: {{.AverageMinutes}}
Average Age: {{.AverageAge}}
Average Viewers: {{.AverageViewers}}
Average Followers: {{.AverageFollowers}}
`

// RenderText renders the template to the specified writer.
func RenderText(w io.Writer) {
	tmpl, err := template.New("cbstats").Parse(text)
	HandleError(err)

	err = tmpl.Execute(w, Stats)
	HandleError(err)
}
