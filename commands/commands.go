package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OP string

const (
	READ   = OP("get")
	EXEC   = OP("exec")
	SAVE   = OP("save")
	DELETE = OP("delete")
	LIST   = OP("list")
	GUI    = OP("gui")
	HELP   = OP("--help")
)

type IndexOp struct {
	OP      OP
	Index   int
	Command *string
}

//FindCommand parses the args and returns the correstponding command
func FindCommand(args []string) *IndexOp {
	baseOp := findBaseCommand(args)
	index := findIndex(baseOp, args)
	command := findExecCommand(baseOp, args)
	return &IndexOp{OP: baseOp, Index: index, Command: command}
}

func findBaseCommand(args []string) OP {
	baseCommand := args[1]
	switch baseCommand {
	case string(READ):
		return READ
	case string(EXEC):
		return EXEC
	case string(SAVE):
		return SAVE
	case string(DELETE):
		return DELETE
	case string(LIST):
		return LIST
	case string(GUI):
		return GUI
	default:
		return HELP
	}
}

func findIndex(baseOp OP, args []string) int {
	if baseOp == LIST || baseOp == HELP || baseOp == SAVE || baseOp == GUI {
		return -1
	}
	if len(args) <= 2 {
		fmt.Print("%+v requires more arguments", baseOp)
		os.Exit(1)
	}
	index, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("%s is not an integer", args[2])
		os.Exit(1)
	}
	return index
}

func findExecCommand(baseOp OP, args []string) *string {
	if baseOp != SAVE {
		return nil
	}
	if len(args) <= 2 {
		fmt.Printf("%+v requires more arguments", baseOp)
		os.Exit(1)
	}
	concat := strings.Join(args[2:], " ")
	return &concat

}
