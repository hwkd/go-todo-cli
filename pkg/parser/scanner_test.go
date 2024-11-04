package parser

import "testing"

func TestScannerSuccessCase(t *testing.T) {
	scanner := NewScanner("-t \"Title\" -d \"Description\"")
	tokens := []Token{
		{OPTION, "t"},
		{VALUE, "Title"},
		{OPTION, "d"},
		{VALUE, "Description"},
		{EOF, EOF},
	}

	for _, expected := range tokens {
		token := scanner.NextToken()
		if token.Type != expected.Type {
			t.Errorf("Expected \"%s\", got \"%s\"", expected.Type, token.Type)
		}
		if token.Literal != expected.Literal {
			t.Errorf("Expected \"%s\", got \"%s\"", expected.Literal, token.Literal)
		}
	}
}

func TestScannerAbnormalOption(t *testing.T) {
	scanner := NewScanner("- \"Title\"")

	token := scanner.NextToken()
	if token.Type != OPTION {
		t.Errorf("Expected \"%s\", got \"%s\"", OPTION, token.Type)
	}
	if token.Literal != "" {
		t.Errorf("Expected \"\", got \"%s\"", token.Literal)
	}
}
