package main

const (
	lowercaseAlph = "abcdefghijklmnopqrstuvwxyz"
	uppercaseAlph = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits        = "0123456789"
	letters       = lowercaseAlph + uppercaseAlph
	alphanumeric  = letters + digits
	punctuation   = "!#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	whitespace    = " \t\n\r\f\v"
)

const (
	leftCodeInput       = "{%"
	rightCodeInput      = "%}"
	leftOutput          = "{{"
	rightOutput         = "}}"
	dot                 = "."
	eof            rune = -1
)

// Operators
const (
	equals               = "=="
	assignEquals         = "="
	doesNotEqual         = "!="
	greaterThan          = ">"
	lessThan             = "<"
	greaterThanOrEqualTo = ">="
	lessThanOrEqualTo    = "<="
)

// Keywords - Opening tags
const (
	cycle     = "cycle"
	render    = "render"
	raw       = "raw"
	comment   = "comment"
	increment = "increment"
	unless    = "unless"
	decrement = "decrement"
	capture   = "capture"
	continue_ = "continue"
	include   = "include"
	case_     = "case"
	ifchanged = "ifchanged"
	assign    = "assign"
	for_      = "for"
	break_    = "break"
	if_       = "if"
	echo      = "echo"
	tablerow  = "tablerow"
	doc       = "doc"
	liquid    = "liquid"
)

// Keywords - Closing tags
const (
	endraw       = "endraw"
	endcomment   = "endcomment"
	endunless    = "endunless"
	endcapture   = "endcapture"
	endcase      = "endcase"
	endfor       = "endfor"
	endif        = "endif"
	endtablerow  = "endtablerow"
	endifchanged = "endifchanged"
)

// Keywords - Conditional/flow control
const (
	when       = "when"
	else_      = "else"
	elsif      = "elsif"
	elseif     = "elseif"
	logicalOr  = "or"
	logicalAnd = "and"
	contains   = "contains"
)

func initKeywordMap() map[string]ItemType {
	kwMap := make(map[string]ItemType)

	// Opening tags
	kwMap[cycle] = itemCycle
	kwMap[render] = itemRender
	kwMap[raw] = itemRaw
	kwMap[comment] = itemComment
	kwMap[increment] = itemIncrement
	kwMap[unless] = itemUnless
	kwMap[decrement] = itemDecrement
	kwMap[capture] = itemCapture
	kwMap[continue_] = itemContinue
	kwMap[include] = itemInclude
	kwMap[case_] = itemCase
	kwMap[ifchanged] = itemIfchanged
	kwMap[assign] = itemAssign
	kwMap[for_] = itemFor
	kwMap[break_] = itemBreak
	kwMap[if_] = itemIf
	kwMap[echo] = itemEcho
	kwMap[tablerow] = itemTablerow
	kwMap[doc] = itemDoc
	kwMap[liquid] = itemLiquid

	// Closing tags
	kwMap[endraw] = itemEndraw
	kwMap[endcomment] = itemEndComment
	kwMap[endunless] = itemEndunless
	kwMap[endcapture] = itemEndcapture
	kwMap[endcase] = itemEndcase
	kwMap[endfor] = itemEndfor
	kwMap[endif] = itemEndif
	kwMap[endtablerow] = itemEndtablerow
	kwMap[endifchanged] = itemEndifchanged

	// Conditional/flow control
	kwMap[when] = itemWhen
	kwMap[else_] = itemElse
	kwMap[elsif] = itemElsif
	kwMap[elseif] = itemElseif
	kwMap[logicalOr] = itemLogicalOr
	kwMap[logicalAnd] = itemLogicalAnd
	kwMap[contains] = itemContains

	return kwMap
}

func initOperatorsMap() map[string]ItemType {
	opMap := make(map[string]ItemType)

	opMap[equals] = itemEquals
	opMap[assignEquals] = itemAssignEquals
	opMap[doesNotEqual] = itemDoesNotEqual
	opMap[greaterThan] = itemGreaterThan
	opMap[lessThan] = itemLessThan
	opMap[greaterThanOrEqualTo] = itemGreaterThanOrEqualTo
	opMap[lessThanOrEqualTo] = itemLessThanOrEqualTo

	return opMap
}
