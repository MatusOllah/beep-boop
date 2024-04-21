package main

import (
	"fmt"
	"log/slog"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func makeStatusBar(a fyne.App, w fyne.Window) fyne.CanvasObject {
	srEntry := NewNumEntry()
	srEntry.SetText(fmt.Sprint(bbSampleRate))
	srEntry.OnChanged = func(s string) {
		sr, _ := strconv.Atoi(s)
		bbSampleRate = beep.SampleRate(sr)
		bbCtrl.Streamer = beep.Resample(1, bbSampleRate, sampleRate, bbgen)
		slog.Info("new sample rate", "bbSampleRate", bbSampleRate)
	}

	modeEntry = widget.NewSelect([]string{"Bytebeat", "Floatbeat"}, func(s string) {
		mode = s
		slog.Info("new mode", "mode", mode)
	})
	modeEntry.SetSelected(mode)

	volSlider := widget.NewSlider(-10, 0)
	volSlider.SetValue(0)
	volSlider.OnChanged = func(f float64) {
		speaker.Lock()
		bbVol.Volume = f
		speaker.Unlock()
	}
	volSlider.Step = 0.1

	bar := container.NewHBox(
		widget.NewLabel("v"+a.Metadata().Version),
		widget.NewSeparator(),
		volSlider,
		widget.NewSeparator(),
		modeEntry,
		widget.NewSeparator(),
		widget.NewLabel("Sample Rate:"),
		srEntry,
	)

	return container.NewBorder(widget.NewSeparator(), nil, nil, nil, bar)
}
