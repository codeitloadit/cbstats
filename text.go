package main

const Text = `CB Stats
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
Private Rooms: {{if .TypeCounts.private}}{{.TypeCounts.private}}{{- else}}0{{- end}}
Group Rooms: {{if .TypeCounts.group}}{{.TypeCounts.group}}{{- else}}0{{- end}}
Away Rooms: {{if .TypeCounts.away}}{{.TypeCounts.away}}{{- else}}0{{- end}}
HD Rooms: {{.HDRooms}}
New Rooms: {{.NewRooms}}
-------------------------
Average Minutes: {{.AverageMinutes}}
Average Age: {{.AverageAge}}
Average Viewers: {{.AverageViewers}}
Average Followers: {{.AverageFollowers}}
`
