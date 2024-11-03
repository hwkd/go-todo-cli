package parser

import "testing"

func TestScanner(t *testing.T) {
	scanner := NewScanner("-t \"Title\" -d \"Description\"")
	token := scanner.NextToken()
	if token.Type != OPTION {
		t.Errorf("Expected %s, got %s", OPTION, token.Type)
	}
	if token.Literal != "t" {
		t.Errorf("Expected t, got %s", token.Literal)
	}
	token = scanner.NextToken()
	if token.Type != VALUE {
		t.Errorf("Expected %s, got %s", VALUE, token.Type)
	}
	if token.Literal != "Title" {
		t.Errorf("Expected \"Title\", got %s", token.Literal)
	}

	token = scanner.NextToken()
	if token.Type != OPTION {
		t.Errorf("Expected %s, got %s", OPTION, token.Type)
	}
	if token.Literal != "d" {
		t.Errorf("Expected d, got %s", token.Literal)
	}
	token = scanner.NextToken()
	if token.Type != VALUE {
		t.Errorf("Expected %s, got %s", VALUE, token.Type)
	}
	if token.Literal != "Description" {
		t.Errorf("Expected \"Description\", got %s", token.Literal)
	}
}
