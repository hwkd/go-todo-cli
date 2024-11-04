package parser

type Scanner struct {
	input        string
	char         byte
	position     int
	readPosition int
}

func NewScanner(input string) *Scanner {
	scanner := &Scanner{input: input}
	scanner.readChar()
	return scanner
}

// TODO: Remove error and let the parser decide if there's an error
func (s *Scanner) NextToken() *Token {
	var token *Token

	s.skipWhiteSpaces()

	switch s.char {
	case '-':
		s.readChar()
		if isLetter(s.char) {
			token = &Token{
				Type:    OPTION,
				Literal: string(s.char),
			}
		} else {
			token = &Token{
				Type:    OPTION,
				Literal: "",
			}
		}
	case '"':
		s.readChar()
		token = &Token{
			Type:    VALUE,
			Literal: s.readString(),
		}
	case 0:
		token = &Token{
			Type:    EOF,
			Literal: EOF,
		}
	}

	s.readChar()
	return token
}

func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input) {
		s.char = 0
	} else {
		s.char = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition++
}

func (s *Scanner) peekChar() byte {
	if s.readPosition >= len(s.input) {
		return 0
	} else {
		return s.input[s.readPosition]
	}
}

func (s *Scanner) skipWhiteSpaces() {
	for s.char == ' ' || s.char == '\n' || s.char == '\t' || s.char == '\r' {
		s.readChar()
	}
}

func (s *Scanner) readString() string {
	position := s.position
	for s.char != '"' {
		s.readChar()
	}
	str := s.input[position:s.position]
	s.readChar()
	return str
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
