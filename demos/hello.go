package main

import (
	"runtime"

	"engo.io/audio"
)

func main() {
	player, err := audio.NewSimplePlayer("assets/birds.wav")
	if err != nil {
		panic(err)
	}
	player.Play()
	player.Total()

	if runtime.GOARCH != "js" {
		for {

		}
	}
}
