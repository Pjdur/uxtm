package uxtm

import "github.com/nsf/termbox-go"

func drawString(x, y int, fg, bg termbox.Attribute, msg string) {
	for i, r := range msg {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}
