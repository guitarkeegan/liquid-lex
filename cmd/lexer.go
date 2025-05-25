package main

type Lexer struct {
	Input      string
	Start      int
	Cur        int
	Width      int
	LastItem   Item
	KeywordMap map[string]ItemType
	OpMap      map[string]ItemType
	// output chan
}

// TODO: Rename if package becomes lexer
func NewLexer(input string) (*Lexer, error) {
	l := &Lexer{
		Input:      input,
		KeywordMap: initKeywordMap(),
		OpMap:      initOperatorsMap(),
	}

	// call go run()
	return l, nil
}
