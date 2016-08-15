package main

import (
	"fmt"
	"runtime"
	"time"

	"engo.io/audio"
)

func main() {
	player, err := audio.NewSimplePlayer("assets/birds.wav")
	if err != nil {
		panic(err)
	}
	player.Play()
	player.Total(false)
	fmt.Println(player.Volume())

	time.Sleep(time.Second * 5)
	player.SetVolume(0.2)
	fmt.Println(player.Volume())

	if runtime.GOARCH != "js" {
		for {
			time.Sleep(time.Hour)
		}
	}
}
