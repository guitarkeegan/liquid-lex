package main

import "sync"

type Lexer struct {
	Input      string
	Start      int
	Cur        int
	Width      int
	LastItem   Item
	KeywordMap map[string]ItemType
	OpMap      map[string]ItemType
	tokens     []Item
	mu         sync.Mutex
}

func NewLexer(input string) (*Lexer, error) {
	l := &Lexer{
		Input:      input,
		KeywordMap: initKeywordMap(),
		OpMap:      initOperatorsMap(),
		tokens:     make([]Item, 0),
	}
	return l, nil
}
