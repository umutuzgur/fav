package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TerminalCommandList struct {
	TerminalCommands []string
}

// List list all the saved commands
func retrieveList() TerminalCommandList {
	dir := createHomeDirectoyIfMissing()
	file, e := ioutil.ReadFile(dir)
	if e != nil {
		fmt.Printf("File error: %v", e)
		os.Exit(1)
	}
	var terminalCommandList TerminalCommandList
	err := json.Unmarshal(file, &terminalCommandList)
	if err != nil {
		return TerminalCommandList{TerminalCommands: []string{}}
	}
	return terminalCommandList
}

// List list and prints all the saved commands
func List() {
	list := retrieveList()
	for index, command := range list.TerminalCommands {
		fmt.Printf("%d. %s\n", index, command)
	}
}
