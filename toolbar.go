package main

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeToolbar(a fyne.App, w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewToolbar(
			widget.NewToolbarAction(theme.FileIcon(), func() {
				slog.Info("selected toolbar item New")
				newFile(a, w)
			}),
			widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
				slog.Info("selected toolbar item Open")
				openFile(a, w)
			}),
			widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
				slog.Info("selected toolbar item Save")
				saveFile(a, w)
			}),
			widget.NewToolbarSeparator(),
			widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
				slog.Info("selected toolbar item Play")
				play()
			}),
			widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
				slog.Info("selected toolbar item Pause")
				pause()
			}),
			widget.NewToolbarAction(theme.MediaStopIcon(), func() {
				slog.Info("selected toolbar item Stop")
				stop()
			}),
			widget.NewToolbarAction(theme.MediaFastRewindIcon(), func() {
				slog.Info("selected toolbar item Rewind")
				rewind()
			}),
			widget.NewToolbarAction(theme.MediaFastForwardIcon(), func() {
				slog.Info("selected toolbar item Forward")
				forward()
			}),
		),
		widget.NewSeparator(),
	)
}
