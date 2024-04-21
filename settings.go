package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	appearance "fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showSettings(a fyne.App) {
	w := a.NewWindow("Settings")

	srEntry := NewNumEntry()
	srEntry.SetText(fmt.Sprint(a.Preferences().IntWithFallback("OutputSampleRate", 48000)))
	srEntry.OnChanged = func(s string) {
		sr, _ := strconv.Atoi(s)
		a.Preferences().SetInt("OutputSampleRate", sr)
	}

	dpsrEntry := NewNumEntry()
	dpsrEntry.SetText(fmt.Sprint(a.Preferences().IntWithFallback("DefaultProjectSampleRate", 8000)))
	dpsrEntry.OnChanged = func(s string) {
		sr, _ := strconv.Atoi(s)
		a.Preferences().SetInt("DefaultProjectSampleRate", sr)
	}

	rqEntry := NewNumEntry()
	rqEntry.SetText(fmt.Sprint(a.Preferences().IntWithFallback("ResampleQuality", 1)))
	rqEntry.OnChanged = func(s string) {
		sr, _ := strconv.Atoi(s)
		a.Preferences().SetInt("ResampleQuality", sr)
	}

	w.SetContent(container.NewAppTabs(
		container.NewTabItem("Audio", container.NewVScroll(container.NewVBox(
			container.NewGridWithColumns(2, newBoldLabel("Output Sample Rate"), srEntry),
			container.NewGridWithColumns(2, newBoldLabel("Default Project Sample Rate"), dpsrEntry),
			container.NewGridWithColumns(2, newBoldLabel("Resample Quality"), rqEntry),
		))),
		container.NewTabItem("Appearance", appearance.NewSettings().LoadAppearanceScreen(w)),
	))

	w.Show()
}

func newBoldLabel(text string) *widget.Label {
	return &widget.Label{Text: text, TextStyle: fyne.TextStyle{Bold: true}}
}
