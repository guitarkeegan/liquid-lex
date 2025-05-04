package main

type lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	items chan Item
}

type stateFn func(*lexer) stateFn

func Run() {
	for state := startState; state != nil; {
		state = state(lexer)
	}
}
