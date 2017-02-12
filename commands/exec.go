package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// Exec executes the command with the speficfied index
func Exec(index int) {
	terminalCommandList := retrieveList()
	if index >= len(terminalCommandList.TerminalCommands) {
		fmt.Println("Can't find any command with that index")
		os.Exit(1)
	}
	command := terminalCommandList.TerminalCommands[index]
	run(command)
}

// Exec executes the command with the speficfied index
func run(command string) {
	shell := findCurrentShell()
	cmd := exec.Command(shell, "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

// FindCurrentShell finds the chosen shell app from $SHELL env variable
func findCurrentShell() string {
	return os.Getenv("SHELL")
}
