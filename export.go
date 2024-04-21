package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/wav"
)

func exportAudio(a fyne.App, w fyne.Window) {
	lenEntry := widget.NewEntry()
	lenEntry.SetText("1m")

	srEntry := NewNumEntry()
	srEntry.SetText(fmt.Sprint(sampleRate))

	chansEntry := widget.NewRadioGroup([]string{"mono", "stereo"}, func(s string) {
		return
	})
	chansEntry.SetSelected("stereo")
	chansEntry.Horizontal = true

	percisionEntry := NewNumEntry()
	percisionEntry.SetText("2")

	rqEntry := NewNumEntry()
	rqEntry.SetText("1")

	dialog.NewForm("Export Audio", "OK", "Cancel", []*widget.FormItem{
		widget.NewFormItem("Length", lenEntry),
		widget.NewFormItem("Sample Rate", srEntry),
		widget.NewFormItem("Channels", chansEntry),
		widget.NewFormItem("Percision", percisionEntry),
		widget.NewFormItem("Resample Quality", rqEntry),
	}, func(b bool) {
		if !b {
			return
		}

		dlg := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			if uc == nil {
				return
			}

			path := uc.URI().Path()

			slog.Info("exporting audio", "path", path)

			len, err := time.ParseDuration(lenEntry.Text)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			_sr, _ := strconv.Atoi(srEntry.Text)
			sr := beep.SampleRate(_sr)

			percision, _ := strconv.Atoi(percisionEntry.Text)

			rq, _ := strconv.Atoi(rqEntry.Text)

			var numChannels int
			switch chansEntry.Selected {
			case "mono":
				numChannels = 1
			case "stereo":
				numChannels = 2
			}

			pb := widget.NewProgressBar()
			pb.Min = 0
			pb.Max = float64(bbSampleRate.N(len))

			pbLbl := widget.NewLabel(fmt.Sprintf("%d/%d samples (%.2f%%)", bbgen.T, bbSampleRate.N(len), (float64(bbgen.T)/float64(bbSampleRate.N(len)))*100))

			dlg := dialog.NewCustomWithoutButtons("Rendering...", container.NewVBox(pbLbl, pb), w)
			dlg.Show()

			f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}
			defer f.Close()

			go func() {
				for bbgen.T != bbSampleRate.N(len) {
					pb.SetValue(float64(bbgen.T))
					slog.Info("exporting", "value", bbgen.T, "max", bbSampleRate.N(len), "percentage", (float64(bbgen.T)/float64(bbSampleRate.N(len)))*100)
					pbLbl.SetText(fmt.Sprintf("%d/%d samples (%.2f%%)", bbgen.T, bbSampleRate.N(len), (float64(bbgen.T)/float64(bbSampleRate.N(len)))*100))
					time.Sleep(100 * time.Millisecond)
				}
				dlg.Hide()
			}()

			if err := wav.Encode(f, beep.Resample(rq, bbSampleRate, sr, beep.Take(bbSampleRate.N(len), bbgen)), beep.Format{
				SampleRate:  sr,
				NumChannels: numChannels,
				Precision:   percision,
			}); err != nil {
				slog.Error(err.Error())
				dialog.NewError(err, w).Show()
				return
			}

			slog.Info("done")

		}, w)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".wav"}))
		dlg.SetFileName("out.wav")
		dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
		dlg.Show()
	}, w).Show()
}
