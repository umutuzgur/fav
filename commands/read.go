package commands

import "fmt"
import "os"

// Read shows the call in the index
func Read(index int) {
	terminalCommandList := retrieveList()
	if index >= len(terminalCommandList.TerminalCommands) {
		fmt.Println("Can't find any command with that index")
		os.Exit(1)
	}
	fmt.Printf("%d. %s\n", index, terminalCommandList.TerminalCommands[index])
}
