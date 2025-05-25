package main

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/log"
)

type stateFn func(*Lexer) stateFn

func lexRightCodeInput(l *Lexer) stateFn {
	l.Cur += len(rightCodeInput)
	l.Emit(itemRightCurly)
	return lexText
}

func isLetter(r rune) bool {
	if unicode.IsLetter(r) {
		return true
	}

	return false
}

func lexKeyword(l *Lexer) stateFn {

	dbg("lexKeyword")

	// TODO: would a trie work well here?
	if strings.HasPrefix(l.Input[l.Cur:], comment) {
		l.Cur += len(comment)
		l.Emit(itemComment)
		return lexInsideOutput
	}
	if strings.HasPrefix(l.Input[l.Cur:], endcomment) {
		l.Cur += len(endcomment)
		l.Emit(itemEndComment)
		return lexInsideOutput
	}

	return l.errorf("lexKeyword: unknown keyword: %.10q..., is not a known keyword", l.Input[l.Cur:])
}

func lexInsideOutput(l *Lexer) stateFn {
	for {
		if strings.HasPrefix(l.Input[l.Cur:], rightCodeInput) {
			// TODO: need to know if keyword was found...
			if l.LastToken == leftCodeInput {
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
		}
	}
}

func lexLeftCodeInput(l *Lexer) stateFn {
	l.Cur += len(leftCodeInput)
	l.Emit(itemLeftCurly)
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
