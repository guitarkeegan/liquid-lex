package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/charmbracelet/log"
)

// Helpers
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
	// TODO: send to items chan
	item := Item{
		itemError,
		fmt.Sprintf(format, args...),
	}

	log.Errorf("%s\n", item.Val)
	return nil
}

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

// emit item to chan, reset start of next item
func (l *Lexer) Emit(it ItemType) {
	// TODO send to chan
	log.Infof("Emit: %+v", Item{it, l.Input[l.Start:l.Cur]})
	l.LastToken = l.Input[l.Start:l.Cur]
	l.Start = l.Cur
}
