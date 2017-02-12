package commands

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

//createHomeDirectoyIfMissing creates the missing work folder
func createHomeDirectoyIfMissing() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	favDir := filepath.Join(usr.HomeDir, ".fav")
	if _, err := os.Stat(favDir); os.IsNotExist(err) {
		err := os.Mkdir(favDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	favConfigDir := filepath.Join(favDir, ".config")
	if _, err := os.Stat(favConfigDir); os.IsNotExist(err) {
		_, err := os.Create(favConfigDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return favConfigDir
}
func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func contains(s []string, elem *string) bool {
	for _, e := range s {
		if e == *elem {
			return true
		}
	}
	return false
}
