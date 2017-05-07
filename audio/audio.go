package audio

import "github.com/gordonklaus/portaudio"

var RunningEngine Engine = StartNew()

type Engine struct {
	AudioBuffer portaudio.Buffer
}

func StartNew() Engine {
	var eng Engine
	//portaudio.
	return eng
}