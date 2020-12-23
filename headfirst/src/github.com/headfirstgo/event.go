package main

import (
	"github.com/headfirstgo/gadget"
)

func playList(device gadget.TapeRecorder, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapeRecorder{}
	mixtape := []string{"Jessie`s Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}
