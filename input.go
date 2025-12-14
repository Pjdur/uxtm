package uxtm

import (
	"github.com/nsf/termbox-go"
)

type Input struct {
	Title  string
	Text   string
	Cursor int
	X, Y   int
}

func NewInput(title string) *Input {
	return &Input{
		Title:  title,
		Text:   "",
		Cursor: 0,
	}
}

func (i *Input) Render(x, y int, focused bool) {
	i.X, i.Y = x, y
	termbox.SetCell(x, y, '>', termbox.ColorDefault, termbox.ColorDefault)
	x++
	for _, r := range i.Title {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
	termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	x++
	termbox.SetCell(x, y, ';', termbox.ColorDefault, termbox.ColorDefault)
	x++
	termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	x++

	inputStartX := x
	for _, r := range i.Text {
		termbox.SetCell(x, y, r, termbox.ColorRed, termbox.ColorDefault)
		x++
	}
	underscores := 20 - len(i.Text)
	if underscores > 0 {
		for j := 0; j < underscores; j++ {
			termbox.SetCell(x, y, '_', termbox.ColorRed, termbox.ColorDefault)
			x++
		}
	}
	cursorX := inputStartX + i.Cursor
	if focused {
		termbox.SetCursor(cursorX, y)
	}
}

func (i *Input) HandleEvent(ev termbox.Event) {
	if ev.Type != termbox.EventKey {
		return
	}
	switch ev.Key {
	case termbox.KeyEnter:
		// Handle enter if needed
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		if i.Cursor > 0 {
			i.Text = i.Text[:i.Cursor-1] + i.Text[i.Cursor:]
			i.Cursor--
		}
	case termbox.KeyDelete:
		if i.Cursor < len(i.Text) {
			i.Text = i.Text[:i.Cursor] + i.Text[i.Cursor+1:]
		}
	case termbox.KeyArrowLeft:
		if i.Cursor > 0 {
			i.Cursor--
		}
	case termbox.KeyArrowRight:
		if i.Cursor < len(i.Text) {
			i.Cursor++
		}
	default:
		if ev.Ch != 0 {
			i.Text = i.Text[:i.Cursor] + string(ev.Ch) + i.Text[i.Cursor:]
			i.Cursor++
		}
	}
}

func (i *Input) Height() int {
	return 1
}
