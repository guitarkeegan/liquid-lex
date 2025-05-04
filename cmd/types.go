package main

import "fmt"

type itemType int

const (
	itemError itemType = iota
	itemDot
	itemEOF
	itemElse

	itemIf
	itemEndIf

	itemLeftOutputTag
	itemRightOutputTag

	itemLeftLogicTag
	itemRightLogicTag

	itemLeftWhiteSpaceCtlTag
	itemRightWhiteSpaceCtlTag

	itemLeftOutputWhiteSpaceCtlTag
	itemRightOutputWhiteSpaceCtlTag

	itemComment
	itemCommentEnd
)

type Item struct {
	typ itemType
	val string
}

func (i Item) String() string {
	switch i.typ {
	case itemError:
		return i.val
	case itemEOF:
		return "EOF"
	case itemElse:
		return "else"
	case itemIf:
		return "if"
	case itemEndIf:
		return "endif"
	case itemLeftOutputTag:
		return "{{"
	case itemRightOutputTag:
		return "}}"
	case itemLeftLogicTag:
		return "{%"
	case itemRightLogicTag:
		return "%}"
	case itemLeftWhiteSpaceCtlTag:
		return "{%-"
	case itemRightWhiteSpaceCtlTag:
		return "-%}"
	case itemLeftOutputWhiteSpaceCtlTag:
		return "{{-"
	case itemRightOutputWhiteSpaceCtlTag:
		return "-}}"
	case itemComment:
		return "comment"
	case itemCommentEnd:
		return "endcomment"
	default:
		if len(i.val) > 10 {
			return fmt.Sprintf("%.10q...", i.val)
		}
		return fmt.Sprintf("%q", i.val)
	}
}
