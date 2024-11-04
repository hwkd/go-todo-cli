package parser

import (
	"testing"
)

func TestParserParserArgs(t *testing.T) {
	parser := NewParser([]string{"-t", "Task 1", "-d", "Create a todo"})
	args, err := parser.ParseArgs()
	expected := [][]string{
		{"t", "Task 1"},
		{"d", "Create a todo"},
	}
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(args) != len(expected) {
		t.Errorf("Expected 2 arguments, got %d", len(args))
	}
	for i, arg := range args {
		if arg[0] != expected[i][0] {
			t.Errorf("Expected \"%s\", got \"%s\"", expected[i][0], arg[0])
		}
		if arg[1] != expected[i][1] {
			t.Errorf("Expected \"%s\", got \"%s\"", expected[i][1], arg[1])
		}
	}
}

func TestParserParseArgsWithError(t *testing.T) {
	parser := NewParser([]string{"-t", "Task 1", "-d"})
	_, err := parser.ParseArgs()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	parser = NewParser([]string{"-", "Task 1"})
	_, err = parser.ParseArgs()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestToStringArg(t *testing.T) {
	args := []string{"-t", "Task 1", "-d", "Create a todo"}
	expected := "-t \"Task 1\" -d \"Create a todo\""
	if toStringArg(args) != expected {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, toStringArg(args))
	}
}
