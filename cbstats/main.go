package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Print("Fetching CB Stats")

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for _ = range ticker.C {
			fmt.Print(".")
		}
	}()

	GetRoomsData()
	PopulateStats()
	ticker.Stop()
	RenderText(os.Stdout)
}

// HandleError is shorthand for the err -> panic idiom.
func HandleError(err error) {
	if err != nil {
		fmt.Println("\n")
		panic(err.Error())
	}
}
