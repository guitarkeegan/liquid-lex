package main

import "fmt"

type ItemType int

const (
	// tags
	itemLeftCurly ItemType = iota
	itemRightCurly
	// keywords
	itemComment
	itemEndComment
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
