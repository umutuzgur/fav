package commands

import "testing"
import "fmt"

func TestParseHelp(t *testing.T) {
	op, _ := FindCommand([]string{"fav", "-h"})
	if op.OP != HELP {
		t.Error("Expected help, got ", op)
	}
}

func TestParseRead(t *testing.T) {
	op, err := FindCommand([]string{"fav", "get", "1"})
	if err != nil {
		fmt.Println(err)
	}
	if op.OP != READ && op.Index != 1 {
		t.Error("Expected get, got ", op)
	}
}

func TestParseExec(t *testing.T) {
	op, _ := FindCommand([]string{"fav", "exec", "1"})
	if op.OP != EXEC && op.Index != 1 {
		t.Error("Expected exec, got ", op)
	}
}

func TestParseDelete(t *testing.T) {
	op, _ := FindCommand([]string{"fav", "delete", "1"})
	if op.OP != DELETE && op.Index != 1 {
		t.Error("Expected delete, got ", op)
	}
}

func TestParseList(t *testing.T) {
	op, _ := FindCommand([]string{"fav", "list"})
	if op.OP != LIST {
		t.Error("Expected list, got ", op)
	}
}

func TestFailParseSave(t *testing.T) {
	command := "ssh root@127.0.0.1"
	op, _ := FindCommand([]string{"fav", "save", command})
	if op.OP != SAVE && op.Command != &command {
		t.Error("Expected save, got ", op)
	}
}

func TestMissingIndex(t *testing.T) {
	expected := "exec requires more arguments"
	op, err := FindCommand([]string{"fav", "exec"})
	if err.Error() != expected {
		t.Errorf("Expected %s, got %v", expected, op)
	}
}

func TestNotAnIndex(t *testing.T) {
	expected := "a is not an integer"
	op, err := FindCommand([]string{"fav", "exec", "a"})
	if err.Error() != expected {
		t.Errorf("Expected %s, got %v", expected, op)
	}
}

func TestMissingCommand(t *testing.T) {
	expected := "save requires more arguments"
	op, err := FindCommand([]string{"fav", "save"})
	if err.Error() != expected {
		t.Errorf("Expected %s, got %v", expected, op)
	}
}
