package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Save add the command to json
func Save(command *string) {
	terminalCommandLists := retrieveList()
	if contains(terminalCommandLists.TerminalCommands, command) {
		fmt.Printf("%s already exists in the list", *command)
		os.Exit(1)
	}
	newCommand := *command
	terminalCommandLists.TerminalCommands = append(terminalCommandLists.TerminalCommands, newCommand)
	dir := createHomeDirectoyIfMissing()
	listJSON, err := json.Marshal(terminalCommandLists)
	if err != nil {
		fmt.Println("Can't turn the terminal commands to JSON: ", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(dir, []byte(listJSON), 0)
	if err != nil {
		fmt.Println("Can't write JSON to the file: ", err)
		os.Exit(1)
	}
	fmt.Printf("%d. %s\n", len(terminalCommandLists.TerminalCommands)-1, newCommand)
}
