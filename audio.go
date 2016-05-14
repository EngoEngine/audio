/*
Package audio provides is a cross platform go audio library.
It currently supports desktop and mobile using golang.org/x/mobile/exp/audio.
To use this library on linux desktop, you will need an OpenAL library
as an external dependency. You can get this on Ubuntu using:
    sudo apt-get install libopenal-dev
Web support via Gopherjs is planned.
*/
package audio //import "engo.io/audio"

import (
	"io"
)

// State indicates the current playing state of the player.
type State int

const (
	Unknown State = iota
	Initial
	Playing
	Paused
	Stopped
)

var stateStrings = [...]string{
	Unknown: "unknown",
	Initial: "initial",
	Playing: "playing",
	Paused:  "paused",
	Stopped: "stopped",
}

func (s State) String() string { return stateStrings[s] }

// ReadSeekCloser is an io.ReadSeeker and io.Closer.
type ReadSeekCloser interface {
	io.ReadSeeker
	io.Closer
}
