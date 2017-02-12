package main

import (
	"os"

	"github.com/umutuzgur/fav/commands"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		commands.Help()
		os.Exit(1)
	}
	op := commands.FindCommand(args)
	switch op.OP {
	case commands.LIST:
		commands.List()
	case commands.GUI:
		commands.Gui()
	case commands.EXEC:
		commands.Exec(op.Index)
	case commands.READ:
		commands.Read(op.Index)
	case commands.SAVE:
		commands.Save(op.Command)
	case commands.DELETE:
		commands.Delete(op.Index)
	case commands.HELP:
		commands.Help()
	}
}
