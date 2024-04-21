package main

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func makeMainMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	// File
	file_new := fyne.NewMenuItem("New", func() {
		slog.Info("selected menu item File>New")
		newFile(a, w)
	})
	file_new.Icon = theme.FileIcon()

	file_open := fyne.NewMenuItem("Open", func() {
		slog.Info("selected menu item File>Open")
		openFile(a, w)
	})
	file_open.Icon = theme.FolderOpenIcon()

	file_save := fyne.NewMenuItem("Save", func() {
		slog.Info("selected menu item File>Save")
		saveFile(a, w)
	})
	file_save.Icon = theme.DocumentSaveIcon()

	file_saveAs := fyne.NewMenuItem("Save As", func() {
		slog.Info("selected menu item File>Save As")
		saveFileAs(a, w)
	})
	file_saveAs.Icon = theme.DocumentSaveIcon()

	file_export_audio := fyne.NewMenuItem("Export Audio", func() {
		slog.Info("selected menu item File>Export Audio")
		exportAudio(a, w)
	})
	file_export_audio.Icon = theme.MediaMusicIcon()

	file := fyne.NewMenu("File",
		file_new,
		fyne.NewMenuItemSeparator(),
		file_open,
		fyne.NewMenuItemSeparator(),
		file_save,
		file_saveAs,
		fyne.NewMenuItemSeparator(),
		file_export_audio,
	)

	// Edit
	edit_settings := fyne.NewMenuItem("Settings", func() {
		slog.Info("selected menu item Edit>Settings")
		showSettings(a)
	})
	edit_settings.Icon = theme.SettingsIcon()

	edit := fyne.NewMenu("Edit",
		edit_settings,
	)

	// Help
	help_about := fyne.NewMenuItem("About", func() {
		slog.Info("selected menu item Help>About")
		showAboutDialog(a, w)
	})
	help_about.Icon = theme.InfoIcon()

	help := fyne.NewMenu("Help",
		help_about,
	)

	return fyne.NewMainMenu(file, edit, help)
}
