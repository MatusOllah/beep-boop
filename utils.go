package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/gopxl/beep/speaker"
)

func updateWindowTitle(a fyne.App, w fyne.Window) {
	w.SetTitle(openFileName + " - " + a.Metadata().Name)
}

// windowSizeToDialog scales the window size to a suitable dialog size.
func windowSizeToDialog(s fyne.Size) fyne.Size {
	return fyne.NewSize(s.Width*0.8, s.Height*0.8)
}

func play() {
	if isPlaying {
		return
	}
	isPlaying = true

	speaker.Play(bbVol)
}

func stop() {
	if !isPlaying {
		return
	}
	isPlaying = false

	speaker.Clear()
	bbgen.T = 0
}

func pause() {
	bbCtrl.Paused = !bbCtrl.Paused
	slog.Info("", "Paused", bbCtrl.Paused)
}

func rewind() {
	bbgen.T -= 1e4
}

func forward() {
	bbgen.T += 1e4
}

func newFile(a fyne.App, w fyne.Window) {
	exprEntry.SetText("")
	bbgen.SetExpr("")
	openFileName = "Untitled"
	openFilePath = ""
	updateWindowTitle(a, w)
}

func openFile(a fyne.App, w fyne.Window) {
	dlg := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		if uc == nil {
			return
		}

		path := uc.URI().Path()

		slog.Info("opening file", "path", path)

		content, err := os.ReadFile(path)
		if err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}
		expr := string(content)

		openFileName = filepath.Base(path)
		openFilePath = path
		exprEntry.SetText(expr)
		if err := bbgen.SetExpr(expr); err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
		}
		updateWindowTitle(a, w)
	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".js", ".mjs"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}

func saveFile(a fyne.App, w fyne.Window) {
	if openFilePath == "" {
		saveFileAs(a, w)
		return
	}

	path := openFilePath

	slog.Info("saving file", "path", path)

	if err := os.WriteFile(path, []byte(exprEntry.Text), 0666); err != nil {
		slog.Error(err.Error())
		dialog.NewError(err, w).Show()
		return
	}

	openFileName = filepath.Base(path)
	openFilePath = path
	updateWindowTitle(a, w)
}

func saveFileAs(a fyne.App, w fyne.Window) {
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

		slog.Info("saving file", "path", path)

		if err := os.WriteFile(path, []byte(exprEntry.Text), 0666); err != nil {
			slog.Error(err.Error())
			dialog.NewError(err, w).Show()
			return
		}

		openFileName = filepath.Base(path)
		openFilePath = path
		updateWindowTitle(a, w)
	}, w)
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".js", ".mjs"}))
	dlg.Resize(windowSizeToDialog(w.Canvas().Size()))
	dlg.Show()
}
