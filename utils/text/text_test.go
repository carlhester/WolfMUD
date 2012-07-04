// Copyright 2012 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package text

import (
	"testing"
	. "wolfmud.org/utils/test"
)

var testFoldSubjects = []struct {
	input  string
	width  int
	output string
}{
	{"The quick brown fox jumps over the lazy dog", 10, "The quick\nbrown fox\njumps over\nthe lazy\ndog"},
	{"Line one\nLine two", 10, "Line one\nLine two"},
	{"One\nTwo\nThree", 10, "One\nTwo\nThree"},
	{"One\n\n\nTwo", 10, "One\n\n\nTwo"},
	{"abcdefghi\njklmnopqr", 10, "abcdefghi\njklmnopqr"},
	{"abcdefghij\nklmnopqrst", 10, "abcdefghij\nklmnopqrst"},
	{"abcdefghijk\nlmnopqrstuv", 10, "abcdefghijk\nlmnopqrstuv"},
	{"ab cd efgh\nab cd ef gh", 10, "ab cd efgh\nab cd ef\ngh"},
	{"", 10, ""},
	{"A zero width test", 0, "A\nzero\nwidth\ntest"},
	{" test\n  test\n   test", 10, " test\n  test\n   test"},
}

var testColorSubjects = []struct {
	input  string
	output string
}{
	{"[BLACK]Black", "\033[30mBlack"},
	{"[RED]Red", "\033[31mRed"},
	{"[GREEN]Green", "\033[32mGreen"},
	{"[YELLOW]Yellow", "\033[33mYellow"},
	{"[BLUE]Blue", "\033[34mBlue"},
	{"[MAGENTA]Magenta", "\033[35mMagenta"},
	{"[CYAN]Cyan", "\033[36mCyan"},
	{"[WHITE]White", "\033[37mWhite"},
	{"[BLACK]R[RED]a[GREEN]i[YELLOW]n[BLUE]b[MAGENTA]o[CYAN]w[WHITE]", "\033[30mR\033[31ma\033[32mi\033[33mn\033[34mb\033[35mo\033[36mw\033[37m"},
	{"", ""},
}

var testMonochromeSubjects = []struct {
	input  string
	output string
}{
	{"[BLACK]Black", "Black"},
	{"[RED]Red", "Red"},
	{"[GREEN]Green", "Green"},
	{"[YELLOW]Yellow", "Yellow"},
	{"[BLUE]Blue", "Blue"},
	{"[MAGENTA]Magenta", "Magenta"},
	{"[CYAN]Cyan", "Cyan"},
	{"[WHITE]White", "White"},
	{"[BLACK]R[RED]a[GREEN]i[YELLOW]n[BLUE]b[MAGENTA]o[CYAN]w[WHITE]", "Rainbow"},
	{"", ""},
}

var testColorFoldSubjects = []struct {
	input  string
	width  int
	output string
}{
	{"[BLACK]R[RED]a[GREEN]i[YELLOW]n[BLUE]b[MAGENTA]o[CYAN]w[WHITE]", 10, "\033[30mR\033[31ma\033[32mi\033[33mn\033[34mb\033[35mo\033[36mw\033[37m"},
	{"[CYAN]Test test? [RED]More more?", 10, "\033[36mTest test?\n\033[31mMore more?"},
	{"[CYAN]Test test![RED]More more!", 10, "\033[36mTest\ntest!\033[31mMore\nmore!"},
	{"\x1b[37m[CYAN]South Bridge[WHITE]\nYou are standing on the west side of an incomplete bridge. By the looks of it the city wants to expand onto the far banks of the river. Up river to the north you can see another bridge in a similar state of construction.\n[GREEN]\n[CYAN]You can see exits: [YELLOW]West\n\x1b[35m>", 80, "\x1b[37m\x1b[36mSouth Bridge\x1b[37m\nYou are standing on the west side of an incomplete bridge. By the looks of it\nthe city wants to expand onto the far banks of the river. Up river to the north\nyou can see another bridge in a similar state of construction.\n\x1b[32m\n\x1b[36mYou can see exits: \x1b[33mWest\n\x1b[35m>"},
	{"\x1b[37m[CYAN]Trading Post[WHITE]\nYou are standing in a large Trading Post . The only exit is west into the street.\n[GREEN]\n[CYAN]You can see exits: [YELLOW]South\n\x1b[35m>", 80, "\x1b[37m\x1b[36mTrading Post\x1b[37m\nYou are standing in a large Trading Post . The only exit is west into the\nstreet.\n\x1b[32m\n\x1b[36mYou can see exits: \x1b[33mSouth\n\x1b[35m>"},
	{"", 10, ""},
}

var testMonochromeFoldSubjects = []struct {
	input  string
	width  int
	output string
}{
	{"[BLACK]R[RED]a[GREEN]i[YELLOW]n[BLUE]b[MAGENTA]o[CYAN]w[WHITE]", 10, "Rainbow"},
	{"[CYAN]Test test! [RED]More more!", 10, "Test test!\nMore more!"},
	{"[CYAN]Test test![RED]More more!", 10, "Test\ntest!More\nmore!"},
	{"", 10, ""},
}

func TestFold(t *testing.T) {
	for _, s := range testFoldSubjects {
		Equal(t, "Fold", s.output, Fold(s.input, s.width))
	}
}

func TestColorize(t *testing.T) {
	for _, s := range testColorSubjects {
		Equal(t, "Colorize", s.output, Colorize(s.input))
	}
}

func TestMonochrome(t *testing.T) {
	for _, s := range testMonochromeSubjects {
		Equal(t, "Monochrome", s.output, Monochrome(s.input))
	}
}

func TestColorizeAndFold(t *testing.T) {
	for _, s := range testColorFoldSubjects {
		Equal(t, "Colorize & Fold", s.output, Fold(Colorize(s.input), s.width))
	}
}

func TestMonochromeAndFold(t *testing.T) {
	for _, s := range testMonochromeFoldSubjects {
		Equal(t, "Colorize & Fold", s.output, Fold(Monochrome(s.input), s.width))
	}
}

func BenchmarkFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fold("the quick brown fox jumps over the lazy dog.", 15)
		Fold("the quick brown fox\njumps over the lazy dog.", 15)
		Fold("the\nquick\nbrown\nfox\njumps\nover\nthe\nlazy\ndog.", 15)
		Fold("[RED]the [GREEN]quick [BROWN]brown [YELLOW]fox [BLUE]jumps [MAGENTA]over [CYAN]the [WHITE]lazy dog.", 15)
	}
}
