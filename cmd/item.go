package main

import "fmt"

type ItemType int

const (
	// Basic symbols
	itemLeftCodeInput ItemType = iota
	itemRightCodeInput
	itemLeftOutput
	itemRightOutput
	itemDot

	// Opening tag keywords
	itemCycle
	itemRender
	itemRaw
	itemComment
	itemIncrement
	itemUnless
	itemDecrement
	itemCapture
	itemContinue
	itemInclude
	itemCase
	itemIfchanged
	itemAssign
	itemFor
	itemBreak
	itemIf
	itemEcho
	itemTablerow
	itemDoc
	itemLiquid

	// Closing tag keywords
	itemEndraw
	itemEndComment
	itemEndunless
	itemEndcapture
	itemEndcase
	itemEndfor
	itemEndif
	itemEndtablerow
	itemEndifchanged

	// Conditional keywords
	itemWhen
	itemElse
	itemElsif
	itemElseif
	itemLogicalOr
	itemLogicalAnd
	itemContains

	// Operators
	itemEquals
	itemAssignEquals
	itemDoesNotEqual
	itemGreaterThan
	itemLessThan
	itemGreaterThanOrEqualTo
	itemLessThanOrEqualTo

	// Variables and literals
	itemVar
	itemText
	itemString
	itemNumber

	// Special tokens
	itemEOF
	itemError
)

var itemTypeNames = map[ItemType]string{
	itemLeftCodeInput:        "LEFT_CODE_INPUT",
	itemRightCodeInput:       "RIGHT_CODE_INPUT",
	itemLeftOutput:           "LEFT_OUTPUT",
	itemRightOutput:          "RIGHT_OUTPUT",
	itemDot:                  "DOT",
	itemCycle:                "CYCLE",
	itemRender:               "RENDER",
	itemRaw:                  "RAW",
	itemComment:              "COMMENT",
	itemIncrement:            "INCREMENT",
	itemUnless:               "UNLESS",
	itemDecrement:            "DECREMENT",
	itemCapture:              "CAPTURE",
	itemContinue:             "CONTINUE",
	itemInclude:              "INCLUDE",
	itemCase:                 "CASE",
	itemIfchanged:            "IFCHANGED",
	itemAssign:               "ASSIGN",
	itemFor:                  "FOR",
	itemBreak:                "BREAK",
	itemIf:                   "IF",
	itemEcho:                 "ECHO",
	itemTablerow:             "TABLEROW",
	itemDoc:                  "DOC",
	itemLiquid:               "LIQUID",
	itemEndraw:               "ENDRAW",
	itemEndComment:           "ENDCOMMENT",
	itemEndunless:            "ENDUNLESS",
	itemEndcapture:           "ENDCAPTURE",
	itemEndcase:              "ENDCASE",
	itemEndfor:               "ENDFOR",
	itemEndif:                "ENDIF",
	itemEndtablerow:          "ENDTABLEROW",
	itemEndifchanged:         "ENDIFCHANGED",
	itemWhen:                 "WHEN",
	itemElse:                 "ELSE",
	itemElsif:                "ELSIF",
	itemElseif:               "ELSEIF",
	itemLogicalOr:            "OR",
	itemLogicalAnd:           "AND",
	itemContains:             "CONTAINS",
	itemEquals:               "EQUALS",
	itemAssignEquals:         "ASSIGN_EQUALS",
	itemDoesNotEqual:         "NOT_EQUALS",
	itemGreaterThan:          "GREATER_THAN",
	itemLessThan:             "LESS_THAN",
	itemGreaterThanOrEqualTo: "GREATER_THAN_OR_EQUAL",
	itemLessThanOrEqualTo:    "LESS_THAN_OR_EQUAL",
	itemVar:                  "VARIABLE",
	itemText:                 "TEXT",
	itemString:               "STRING",
	itemNumber:               "NUMBER",
	itemEOF:                  "EOF",
	itemError:                "ERROR",
}

type Item struct {
	Typ ItemType
	Val string
	Pos int // Position in input for better error reporting
}

func (i Item) String() string {
	switch i.Typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.Val
	}
	if len(i.Val) > 10 {
		return fmt.Sprintf("%.10q...", i.Val)
	}
	return fmt.Sprintf("%q", i.Val)
}

func (i Item) TypeString() string {
	if name, ok := itemTypeNames[i.Typ]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN(%d)", int(i.Typ))
}
