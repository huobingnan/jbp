package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func _quit(ui *gocui.Gui, v *gocui.View) error {
	fmt.Printf("Call")
	return gocui.ErrQuit
}

func _layout(ui *gocui.Gui) error {
	mx, _ := ui.Size()
	v, err := ui.SetView("Title", 0, 0, mx-10, 2)
	v.Clear()
	v.Title = "Title"
	_, err = fmt.Fprintf(v, "JVM Byte Code Viewer")
	return err
}

func main() {
	// refactor
	ui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer ui.Close()
	err = ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, _quit)
	if err != nil {
		log.Panicln(err)
	}
	ui.SetManagerFunc(_layout)
	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
