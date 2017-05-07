package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func () {
		mw := StartApplication()
		mw.Show()
	})

	if err != nil {
		println("An error occured! :C\n" + err.Error())
	}
}