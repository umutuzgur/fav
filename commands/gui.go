package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

type selected struct {
	command string
}

func (selected selected) Error() string {
	return "none"
}

var logFile = createLogFile()
var list []string

// Gui creates the UI to interract with commands
func Gui() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	closed := false
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		if !closed {
			g.Close()
		}
	}()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, up); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, down); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("side", gocui.KeyEnter, gocui.ModNone, enter); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		selected, ok := err.(*selected)
		if ok {
			closed = true
			g.Close()
			run(selected.command)
		} else {
			log.Panicln(err)
		}
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := correctCursor(v, true)
		v.SetCursor(cx, cy)
	}
	return nil
}

func enter(g *gocui.Gui, v *gocui.View) error {
	_, y := v.Cursor()
	if command, err := v.Line(y); err == nil {
		return &selected{command: command}
	}
	return gocui.ErrQuit
}

func up(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		nextCx, nextCy := correctCursor(v, false)
		v.SetCursor(nextCx, nextCy)
	}
	return nil
}

func correctCursor(v *gocui.View, down bool) (int, int) {
	cx, cy := v.Cursor()
	if down {
		return cx % len(list), (cy + 1) % len(list)
	}

	if y := (cy - 1) % len(list); y < 0 {
		return 0, y + len(list)
	} else {
		return 0, y
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		createGuiList(v)
	}

	if _, err := g.SetCurrentView("side"); err != nil {
		return err
	}

	return nil
}

func createGuiList(v *gocui.View) {
	v.Highlight = true
	v.SelBgColor = gocui.ColorWhite
	v.SelFgColor = gocui.ColorBlack
	list = retrieveList().TerminalCommands
	for _, elem := range list {
		fmt.Fprintln(v, elem)
	}
}

func createLogFile() *os.File {
	file, _ := os.Create("/tmp/fav.log")
	return file
}

func logToFile(a ...interface{}) {
	fmt.Fprintln(logFile, a)
}
