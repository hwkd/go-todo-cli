package parser

import (
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	scanner *Scanner
}

func NewParser(input []string) *Parser {
	return &Parser{scanner: NewScanner(toStringArg(input))}
}

func (p *Parser) ParseArgs() ([][]string, error) {
	args := [][]string{}
	for {
		option, err := p.parseOption()
		if err == io.EOF {
			return args, nil
		} else if err != nil {
			return nil, err
		}
		args = append(args, []string{option.Option, option.Value.Literal})
	}
}

func (p *Parser) parseOption() (*OptionNode, error) {
	token := p.nextToken()
	if token.Type == EOF {
		return nil, io.EOF
	}
	if token.Type != OPTION {
		return nil, fmt.Errorf("Expected \"%s\" type token, got \"%s\"", OPTION, token.Type)
	}
	if token.Literal == "" {
		return nil, fmt.Errorf("Expected option literal token, got \"%s\"", token.Literal)
	}

	argNode, err := p.parseValue()
	if err != nil {
		return nil, err
	}
	return &OptionNode{
		Option: token.Literal,
		Value:  argNode,
	}, nil
}

func (p *Parser) parseValue() (*ValueNode, error) {
	token := p.nextToken()
	if token.Type != VALUE {
		return nil, fmt.Errorf("Expected \"%s\" type token, got \"%s\"", VALUE, token.Type)
	}
	if token.Literal == "" {
		return nil, fmt.Errorf("Expected a non-empty string value, got \"\"")
	}
	return &ValueNode{
		Literal: token.Literal,
	}, nil
}

func (p *Parser) nextToken() *Token {
	return p.scanner.NextToken()
}

func toStringArg(args []string) string {
	s := ""
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			arg = "\"" + arg + "\""
		}
		s += arg + " "
	}
	return strings.TrimSpace(s)
}
