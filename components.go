package uxtm

import "github.com/nsf/termbox-go"


type Component interface {
	Render(x, y int, focused bool)
	HandleEvent(ev termbox.Event)
	Height() int
}
