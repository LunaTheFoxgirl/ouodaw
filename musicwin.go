package main

import "github.com/andlabs/ui"

type MusicWindow struct {
	title string
	window *ui.Window
}

func (mw* MusicWindow) Show() {
	mw.window.Show()
}

func StartApplication() MusicWindow {
	var mw MusicWindow

	mw.title = "OUODaw"
	w := ui.NewWindow(mw.title, 420, 420, true)
	mw.window = w

	return mw
}
