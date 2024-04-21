package main

import (
	"fmt"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func showAboutDialog(a fyne.App, w fyne.Window) {
	dialog.NewInformation("About gecfg-editor", fmt.Sprintf(
		"%s version %s\nGo version %s %s/%s\n",
		a.Metadata().Name,
		a.Metadata().Version,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	), w).Show()
}
