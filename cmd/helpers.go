package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/charmbracelet/log"
)

func (l *Lexer) Next() rune {
	dbg("Next")
	if l.Cur >= len(l.Input) {
		l.Width = 0
		return eof
	}
	char, width := utf8.DecodeRuneInString(l.Input[l.Cur:])
	l.Width = width
	l.Cur += l.Width
	dbg("  l.Cur: %d, char: %s", l.Cur, string(char))
	return char
}

func (l *Lexer) Ignore() {
	dbg("Ignore")
	l.Start = l.Cur
}

func (l *Lexer) Backup() {
	dbg("Backup")
	l.Cur -= l.Width
}

func (l *Lexer) Peek() rune {
	char := l.Next()
	l.Backup()
	return char
}

func (l *Lexer) Accept(valid string) bool {
	if strings.IndexRune(valid, l.Next()) >= 0 {
		return true
	}
	l.Backup()
	return false
}

func (l *Lexer) AcceptRun(valid string) {
	for strings.IndexRune(valid, l.Next()) >= 0 {
	}
	l.Backup()
}

func (l *Lexer) errorf(format string, args ...any) stateFn {
	item := Item{
		Typ: itemError,
		Val: fmt.Sprintf(format, args...),
		Pos: l.Start,
	}

	l.mu.Lock()
	l.tokens = append(l.tokens, item)
	l.mu.Unlock()

	log.Errorf("Error at position %d: %s", l.Start, item.Val)
	return nil
}

func (l *Lexer) Emit(it ItemType) {
	item := Item{
		Typ: it,
		Val: l.Input[l.Start:l.Cur],
		Pos: l.Start,
	}

	l.mu.Lock()
	l.tokens = append(l.tokens, item)
	l.mu.Unlock()

	dbg("Emit: %+v", item)
	l.LastItem = item
	l.Start = l.Cur
}

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func isOperator(r rune) bool {
	operators := "=><!"
	return strings.ContainsRune(operators, r)
}

func isString(r rune) bool {
	return r == '"' || r == '\''
}

func isNumber(r rune) bool {
	return unicode.IsNumber(r)
}

func isDot(r rune) bool {
	return r == '.'
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}
