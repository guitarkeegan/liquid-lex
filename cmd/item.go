package main

import "fmt"

type ItemType int

const (
	// tags
	itemLeftCodeInput ItemType = iota
	itemRightCodeInput

	// opening tag keywords
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

	// closing tag keywords
	itemEndraw
	itemEndComment
	itemEndunless
	itemEndcapture
	itemEndcase
	itemEndfor
	itemEndif
	itemEndtablerow
	itemEndifchanged

	// case/conditional keywords
	itemWhen
	itemElse
	itemElsif
	itemElseif
	itemLogicalOr
	itemLogicalAnd
	itemContains

	itemEquals
	itemAssignEquals
	itemDoesNotEqual
	itemGreaterThan
	itemLessThan
	itemGreaterThanOrEqualTo
	itemLessThanOrEqualTo

	// outside of map
	itemVar // escape hatch for vars

	// end of file
	itemEOF
	itemError

	// plain text
	itemText
)

type Item struct {
	Typ ItemType
	Val string
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
