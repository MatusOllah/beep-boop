package main

import (
	"fyne.io/fyne/v2/widget"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
)

var (
	openFileName string = "Untitled"
	openFilePath string
	bbgen        *BytebeatGenerator = NewBytebeatGenerator("")
	bbCtrl       *beep.Ctrl         = &beep.Ctrl{
		Streamer: beep.Resample(1, bbSampleRate, beep.SampleRate(48000), bbgen),
		Paused:   false,
	}
	bbVol *effects.Volume = &effects.Volume{
		Streamer: bbCtrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	isPlaying    bool            = false
	bbSampleRate beep.SampleRate = 8000
	sampleRate   beep.SampleRate
	exprEntry    *widget.Entry
	mode         string
	modeEntry    *widget.Select
)
