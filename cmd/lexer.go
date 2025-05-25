package main

type Lexer struct {
	Input     string
	Start     int
	Cur       int
	Width     int
	LastToken string
	// output chan
}

// TODO: Rename
func NewLexer(input string) (*Lexer, error) {
	l := &Lexer{
		Input: input,
	}

	// call go run()
	return l, nil
}
