package main

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI(a fyne.App, w fyne.Window) fyne.CanvasObject {
	exprEntry = widget.NewEntry()
	exprEntry.OnChanged = func(s string) {
		if err := bbgen.SetExpr(s); err != nil {
			slog.Error(err.Error())
		}
	}
	exprEntry.TextStyle = fyne.TextStyle{
		Monospace: true,
	}

	return container.NewBorder(
		makeToolbar(a, w),
		makeStatusBar(a, w),
		nil,
		nil,
		exprEntry,
	)
}
