package main

type statsStruct struct {
	BroadcasterCounts map[string]uint
	ViewerCounts      map[string]uint
	TagCounts         map[string]uint
	TypeCounts        map[string]uint
	RoomsWithTags     uint
	HDRooms           uint
	NewRooms          uint
	AverageMinutes    uint
	AverageAge        uint
	AverageViewers    uint
	AverageFollowers  uint
}

// Stats store all of the data to be rendered by the template.
var Stats statsStruct

// PopulateStats aggregates the data from Rooms into Stats.
func PopulateStats() {
	Stats.BroadcasterCounts = make(map[string]uint)
	Stats.ViewerCounts = make(map[string]uint)
	Stats.TagCounts = make(map[string]uint)
	// Initializing the types to remove logic from the template and avoid <no value>
	Stats.TypeCounts = map[string]uint{
		"public": 0,
		"private": 0,
		"group": 0,
		"away": 0,
	}

	var totalAge uint
	var totalFollowers uint
	var totalSeconds uint

	for _, room := range Rooms {
		Stats.BroadcasterCounts[room.Gender]++
		Stats.BroadcasterCounts["a"]++
		Stats.ViewerCounts[room.Gender] += room.Viewers
		Stats.ViewerCounts["a"] += room.Viewers
		if len(room.Tags) > 0 {
			Stats.RoomsWithTags++
			for _, tag := range room.Tags {
				Stats.TagCounts[tag]++
			}
		}
		Stats.TypeCounts[room.Type]++
		if room.IsHD {
			Stats.HDRooms++
		}
		if room.IsNew {
			Stats.NewRooms++
		}
		totalAge += room.Age
		totalFollowers += room.Followers
		totalSeconds += room.Seconds
	}

	Stats.AverageMinutes = totalSeconds / Stats.BroadcasterCounts["a"] / 60
	Stats.AverageAge = totalAge / Stats.BroadcasterCounts["a"]
	Stats.AverageViewers = Stats.ViewerCounts["a"] / Stats.BroadcasterCounts["a"]
	Stats.AverageFollowers = totalFollowers / Stats.BroadcasterCounts["a"]
}
