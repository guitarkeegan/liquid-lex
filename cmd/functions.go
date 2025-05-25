package main

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/log"
)

type stateFn func(*Lexer) stateFn

func lexRightCodeInput(l *Lexer) stateFn {
	l.Cur += len(rightCodeInput)
	l.Emit(itemRightCodeInput)
	return lexText
}

func isLetter(r rune) bool {
	if unicode.IsLetter(r) {
		return true
	}

	return false
}

func lexOperator(l *Lexer) stateFn {

	l.AcceptRun("=!><")
	op := l.Input[l.Start:l.Cur]
	if item, ok := l.OpMap[op]; ok {
		l.Emit(item)
		return lexInsideOutput
	}

	return l.errorf("lexOperator: op %q not found in OpMap", op)

}

// accept run to get the word, then check agains a map
func lexKeyword(l *Lexer) stateFn {

	dbg("lexKeyword")

	l.AcceptRun(lowercaseAlph)
	// should be at end of keyword

	word := l.Input[l.Start:l.Cur]
	if itemType, ok := l.KeywordMap[word]; ok {
		l.Emit(itemType)
		return lexInsideOutput
	}

	// check for var
	if l.LastItem.Typ == itemAssign {
		l.Emit(itemVar)
		return lexOperator
	}

	return l.errorf("lexKeyword: unknown keyword: %.10q..., is not a known keyword", word)
}

func lexInsideOutput(l *Lexer) stateFn {
	for {
		if strings.HasPrefix(l.Input[l.Cur:], rightCodeInput) {
			if l.LastItem.Val == leftCodeInput {
				return l.errorf("empty output: %s %s must not be empty", leftCodeInput, rightCodeInput)
			}
			dbg("  lexInsideOutput: rightCodeInput")
			return lexRightCodeInput
		}
		switch r := l.Next(); {
		case r == eof:
			return l.errorf("EOF: unclosed output: %q must have a closing tag: %q", leftCodeInput, rightCodeInput)
		case isSpace(r):
			dbg("  lexInsideOutput: isSpace")
			l.Ignore()
		case isLetter(r):
			// check keywords
			dbg("  lexInsideOutput: isLetter")
			l.Backup()
			return lexKeyword
		case isOperator(r):
			return lexOperator
			// TODO: isString
		}
	}
}

func lexLeftCodeInput(l *Lexer) stateFn {
	l.Cur += len(leftCodeInput)
	l.Emit(itemLeftCodeInput)
	return lexInsideOutput
}

func lexText(l *Lexer) stateFn {

	for {
		if strings.HasPrefix(l.Input[l.Cur:], leftCodeInput) {
			if l.Cur > l.Start {
				l.Emit(itemText)
			}
			return lexLeftCodeInput
		}
		if l.Next() == eof {
			break
		}
	}

	if l.Cur > l.Start {
		l.Emit(itemText)
	}
	l.Emit(itemEOF)
	return nil

}

func (l *Lexer) Run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	// close the l.items chan here
	log.Info("Lex Completed")
}
