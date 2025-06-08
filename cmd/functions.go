package main

import (
	"strings"

	"github.com/charmbracelet/log"
)

type stateFn func(*Lexer) stateFn

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

func lexLeftCodeInput(l *Lexer) stateFn {
	l.Cur += len(leftCodeInput)
	l.Emit(itemLeftCodeInput)
	return lexInsideCodeInput
}

func lexLeftOutput(l *Lexer) stateFn {
	l.Cur += len(leftOutput)
	l.Emit(itemLeftOutput)
	return lexInsideCodeInput
}

func lexRightCodeInput(l *Lexer) stateFn {
	l.Cur += len(rightCodeInput)
	l.Emit(itemRightCodeInput)
	return lexText
}

func lexRightOutput(l *Lexer) stateFn {
	l.Cur += len(rightOutput)
	l.Emit(itemRightOutput)
	return lexText
}

func lexInsideCodeInput(l *Lexer) stateFn {
	for {
		if strings.HasPrefix(l.Input[l.Cur:], rightCodeInput) {
			if l.LastItem.Val == leftCodeInput {
				return l.errorf("empty code tags: %s %s must not be empty", leftCodeInput, rightCodeInput)
			}
			return lexRightCodeInput
		}
		if strings.HasPrefix(l.Input[l.Cur:], rightOutput) {
			if l.LastItem.Val == leftOutput {
				return l.errorf("empty output tags: %s %s must not be empty", leftOutput, rightOutput)
			}
			return lexRightOutput
		}

		switch r := l.Next(); {
		case r == eof:
			return l.errorf("unclosed tag: expected closing tag")
		case isSpace(r):
			l.Ignore()
		case isLetter(r):
			l.Backup()
			return lexKeyword
		case isDot(r):
			return lexDot
		case isString(r):
			return lexString
		case isNumber(r):
			l.Backup()
			return lexNumber
		case isOperator(r):
			l.Backup()
			return lexOperator
		default:
			return l.errorf("unexpected character: %q", r)
		}
	}
}

func lexKeyword(l *Lexer) stateFn {
	dbg("lexKeyword")

	// Accept letters, numbers, and underscores for keywords/variables
	l.AcceptRun(alphanumeric + "_")

	word := l.Input[l.Start:l.Cur]
	if itemType, ok := l.KeywordMap[word]; ok {
		l.Emit(itemType)

		// Special handling for comment blocks
		if itemType == itemComment {
			return lexCommentBlock
		}
		if itemType == itemRaw {
			return lexRawBlock
		}

		return lexInsideCodeInput
	}

	// Not a keyword, treat as variable
	l.Emit(itemVar)
	return lexInsideCodeInput
}

func lexCommentBlock(l *Lexer) stateFn {
	// We've already emitted the comment keyword
	// Skip any whitespace after the comment keyword
	for isSpace(l.Peek()) {
		l.Next()
	}
	l.Ignore() // Ignore the whitespace

	// Now we need to find the end of the code block
	if !strings.HasPrefix(l.Input[l.Cur:], rightCodeInput) {
		return l.errorf("expected %s after comment", rightCodeInput)
	}

	l.Cur += len(rightCodeInput)
	l.Emit(itemRightCodeInput)

	// Now we're in comment content - consume everything until {% endcomment %}
	for {
		if strings.HasPrefix(l.Input[l.Cur:], leftCodeInput) {
			// Check if this is the end comment
			pos := l.Cur + len(leftCodeInput)
			if pos < len(l.Input) {
				// Skip whitespace
				for pos < len(l.Input) && isSpace(rune(l.Input[pos])) {
					pos++
				}
				if strings.HasPrefix(l.Input[pos:], endcomment) {
					// Emit any text before the end comment
					if l.Cur > l.Start {
						l.Emit(itemText)
					}
					// Process the end comment tag
					return lexLeftCodeInput
				}
			}
		}
		if l.Next() == eof {
			return l.errorf("unclosed comment block: expected {% endcomment %}")
		}
	}
}

func lexRawBlock(l *Lexer) stateFn {
	// We've already emitted the raw keyword
	// Skip any whitespace after the raw keyword
	for isSpace(l.Peek()) {
		l.Next()
	}
	l.Ignore() // Ignore the whitespace

	// Now we need to find the end of the code block
	if !strings.HasPrefix(l.Input[l.Cur:], rightCodeInput) {
		return l.errorf("expected %s after raw", rightCodeInput)
	}

	l.Cur += len(rightCodeInput)
	l.Emit(itemRightCodeInput)

	// Consume everything until {% endraw %}
	for {
		if strings.HasPrefix(l.Input[l.Cur:], leftCodeInput) {
			pos := l.Cur + len(leftCodeInput)
			if pos < len(l.Input) {
				// Skip whitespace
				for pos < len(l.Input) && isSpace(rune(l.Input[pos])) {
					pos++
				}
				if strings.HasPrefix(l.Input[pos:], endraw) {
					if l.Cur > l.Start {
						l.Emit(itemText)
					}
					return lexLeftCodeInput
				}
			}
		}
		if l.Next() == eof {
			return l.errorf("unclosed raw block: expected {% endraw %}")
		}
	}
}

func lexString(l *Lexer) stateFn {
	quote := l.Input[l.Cur-1] // The quote character we just read

	for {
		r := l.Next()
		if r == eof {
			return l.errorf("unclosed string")
		}
		if r == rune(quote) {
			break
		}
		if r == '\\' {
			// Handle escaped characters
			if l.Next() == eof {
				return l.errorf("unclosed string: unexpected EOF after backslash")
			}
		}
	}

	l.Emit(itemString)
	return lexInsideCodeInput
}

func lexNumber(l *Lexer) stateFn {
	l.AcceptRun(digits)

	// Handle decimal numbers
	if l.Accept(".") {
		l.AcceptRun(digits)
	}

	l.Emit(itemNumber)
	return lexInsideCodeInput
}

func lexDot(l *Lexer) stateFn {
	l.Emit(itemDot)
	return lexInsideCodeInput
}

func lexOperator(l *Lexer) stateFn {
	// Read the operator characters
	start := l.Cur
	for isOperator(l.Peek()) {
		l.Next()
	}

	op := l.Input[start:l.Cur]
	if item, ok := l.OpMap[op]; ok {
		l.Emit(item)
		return lexInsideCodeInput
	}

	return l.errorf("unknown operator: %q", op)
}

func (l *Lexer) Run() []Item {
	// Start the lexing process
	for state := lexText; state != nil; {
		state = state(l)
	}

	log.Info("Lexing completed")
	return l.tokens
}
