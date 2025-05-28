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
		return lexInsideCodeInput
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
		return lexInsideCodeInput
	}

	// check for var
	if l.LastItem.Typ == itemAssign {
		l.Emit(itemVar)
		return lexInsideCodeInput
	}

	return l.errorf("lexKeyword: unknown keyword: %.10q..., is not a known keyword", word)
}

func lexString(l *Lexer) stateFn {
	l.AcceptRun(alphanumeric + punctuation + whitespace)
	if l.Input[l.Cur] == '"' {
		l.Next()
	} else {
		l.errorf("lexString: did not end with '\"' token")
	}
	l.Emit(itemString)
	return lexInsideCodeInput
}

// TODO: just int for now
func lexNumber(l *Lexer) stateFn {
	l.AcceptRun(digits)
	l.Emit(itemNumber)
	return lexInsideCodeInput
}

func lexInsideCodeInput(l *Lexer) stateFn {
	for {
		if strings.HasPrefix(l.Input[l.Cur:], rightCodeInput) {
			if l.LastItem.Val == leftCodeInput {
				return l.errorf("empty output: %s %s must not be empty", leftCodeInput, rightCodeInput)
			}
			dbg("  lexInsideCodeInput: rightCodeInput")
			return lexRightCodeInput
		}
		switch r := l.Next(); {
		case r == eof:
			return l.errorf("EOF: unclosed output: %q must have a closing tag: %q", leftCodeInput, rightCodeInput)
		case isSpace(r):
			dbg("  lexInsideCodeInput: isSpace")
			l.Ignore()
		case isLetter(r):
			// check keywords
			dbg("  lexInsideCodeInput: isLetter")
			l.Backup()
			return lexKeyword
		case isString(r):
			return lexString
		case isNumber(r):
			return lexNumber
		case isOperator(r):
			return lexOperator
			// TODO: isString
		}
	}
}

func lexRightOutput(l *Lexer) stateFn {
	l.Cur += len(rightOutput)
	l.Emit(itemRightOutput)
	return lexText
}

func lexInsideOuput(l *Lexer) stateFn {
	for {
		if strings.HasPrefix(l.Input[l.Cur:], rightOutput) {
			if l.LastItem.Val == leftOutput {
				return l.errorf("empty output: %s %s must not be empty", leftOutput, rightOutput)
			}
			dbg("  lexInsideCodeInput: rightCodeInput")
			return lexRightOutput
		}
		switch r := l.Next(); {
		case r == eof:
			return l.errorf("EOF: unclosed output: %q must have a closing tag: %q", leftCodeInput, rightCodeInput)
		case isSpace(r):
			dbg("  lexInsideCodeInput: isSpace")
			l.Ignore()
		case isLetter(r):
			// check keywords
			dbg("  lexInsideCodeInput: isLetter")
			l.Backup()
			return lexKeyword
		case isString(r):
			return lexString
		case isNumber(r):
			return lexNumber
		case isOperator(r):
			return lexOperator
		}
	}
}

func lexLeftCodeInput(l *Lexer) stateFn {
	l.Cur += len(leftCodeInput)
	l.Emit(itemLeftCodeInput)
	return lexInsideCodeInput
}

func lexLeftOutput(l *Lexer) stateFn {
	l.Cur += len(leftOutput)
	l.Emit(itemLeftCodeInput)
	return lexInsideOuput
}

func lexText(l *Lexer) stateFn {

	for {
		if strings.HasPrefix(l.Input[l.Cur:], leftCodeInput) {
			if l.Cur > l.Start {
				l.Emit(itemText)
			}
			return lexLeftCodeInput
		}
		if strings.HasPrefix(l.Input[l.Cur:], leftOutput) {
			if l.Cur > l.Start {
				l.Emit(itemText)
			}
			return lexLeftOutput
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
