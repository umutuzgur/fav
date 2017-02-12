package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Delete(index int) {
	terminalCommandList := retrieveList()
	if index >= len(terminalCommandList.TerminalCommands) {
		fmt.Println("Can't find any command with that index")
		os.Exit(1)
	}
	terminalCommandList.TerminalCommands = removeIndex(terminalCommandList.TerminalCommands, index)
	dir := createHomeDirectoyIfMissing()
	listJSON, err := json.Marshal(terminalCommandList)
	if err != nil {
		fmt.Println("Can't turn the terminal commands to JSON: ", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(dir, []byte(listJSON), 0)
	if err != nil {
		fmt.Println("Can't write JSON to the file: ", err)
		os.Exit(1)
	}
}
