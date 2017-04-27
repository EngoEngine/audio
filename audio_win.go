// +build windows

package audio

import (
	"log"
	"time"

	"bitbucket.org/StephenPatrick/go-winaudio/winaudio"
)

type Player struct {
	player           winaudio.Audio
	duration         time.Duration
	src              string
	format           int
	samplesPerSecond int64
}

func notImplemented(msg string) {
	log.Println("[WARNING] " + msg + " is not implemented on Windows")
}

func NewPlayer(src string, format int, samplesPerSecond int64) (*Player, error) {
	winaudio.InitWinAudio()
	wavfile, _ := winaudio.LoadWav(src)
	wavfile.SetLooping(true)

	return &Player{player: wavfile, src: src, format: format, samplesPerSecond: samplesPerSecond}, nil
}

func NewSimplePlayer(src string) (*Player, error) {
	return NewPlayer(src, 0, 0)
}

func (p *Player) Close() error {
	notImplemented("close")
	return nil
}

func (p *Player) Current() time.Duration {
	dur, _ := time.ParseDuration("3000ms")
	return dur
}

func (p *Player) Play() error {
	p.player.Play()
	return nil
}

func (p *Player) Pause() error {
	p.player.Pause()
	return nil
}

func (p *Player) SetVolume(volume int32) error {
	p.player.SetVolume(volume)
	return nil
}

func (p *Player) Volume() int32 {
	vol, _ := p.player.GetVolume()
	return vol
}
