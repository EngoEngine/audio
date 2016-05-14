// +build netgo

package audio

import (
	"honnef.co/go/js/dom"
	"log"
	"time"
)

type Player struct {
	player           *dom.HTMLAudioElement
	duration         time.Duration
	src              string
	format           int
	samplesPerSecond int64
}

type Resource interface {
	Url() string
}

func notImplemented(msg string) {
	log.Println("[WARNING] " + msg + " is not implemented on Gopherjs")
}

//TODO: Use format and samples per second
func NewPlayer(src string, format int, samplesPerSecond int64) (*Player, error) {
	player := dom.GetWindow().Document().CreateElement("audio").(*dom.HTMLAudioElement)
	player.SetAttribute("src", src)
	player.AddEventListener("loadeddata", false, func(ev dom.Event) {
		duration, err := time.ParseDuration(player.Underlying().Get("duration").String() + "s")
		if err != nil {
			log.Println("[ERROR] Failed to get audio duration: " + err.Error())
		}
		return &Player{player: player, duration: duration, src: src, format: format, samplesPerSecond: samplesPerSecond}, nil
	})
}

func NewSimplePlayer(src string) (*Player, error) {
	return NewPlayer(src, 0, 0)
}

func (p *Player) Close() error {
	notImplemented("close")
	return nil
}

func (p *Player) Current() time.Duration {
	notImplemented("current")
	return time.Now().Sub(time.Now())
}

func (p *Player) Pause() error {
	p.player.Pause()
	return nil
}

func (p *Player) Play() error {
	p.player.Play()
	return nil
}

func (p *Player) Seek(offset time.Duration) error {
	notImplemented("seek")
	return nil
}

func (p *Player) SetVolume(vol float64) {
	p.player.Underlying().Set("volume", vol)
}

func (p *Player) State() State {
	if p.player.Paused {
		return Paused
	} else {
		return Playing
	}
}

func (p *Player) Stop() error {
	notImplemented("stop")
	return nil
}

func (p *Player) Total() time.Duration {
	return p.duration
}

func (p *Player) Volume() float64 {
	notImplemented("volume")
	return 0
}

func (p *Player) Rewind() {
	notImplemented("rewind")
}

func (p *Player) PlaySources() {
	notImplemented("playSources")
}
func Preload() error {
	notImplemented("preload")
	return nil
}
