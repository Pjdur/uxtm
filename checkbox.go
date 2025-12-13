package uxtm

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Checkbox struct {
	title    string
	options  []string
	selected []bool
	cursor   int
}

func NewCheckbox(title string, options []string) *Checkbox {
	return &Checkbox{
		title:    title,
		options:  options,
		selected: make([]bool, len(options)),
	}
}

func (c *Checkbox) Height() int {
	return len(c.options) + 1
}

func (c *Checkbox) Render(x, y int, focused bool) {
	drawString(
		x, y,
		termbox.ColorCyan|termbox.AttrBold,
		0,
		c.title,
	)

	for i, opt := range c.options {
		prefix := "  "
		color := termbox.ColorWhite

		if focused && i == c.cursor {
			prefix = "> "
			color = termbox.ColorYellow | termbox.AttrBold
		}

		check := "[ ]"
		if c.selected[i] {
			check = "[x]"
		}

		drawString(
			x,
			y+i+1,
			color,
			0,
			fmt.Sprintf("%s%s %s", prefix, check, opt),
		)
	}
}

func (c *Checkbox) HandleEvent(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowUp:
		if c.cursor > 0 {
			c.cursor--
		}
	case termbox.KeyArrowDown:
		if c.cursor < len(c.options)-1 {
			c.cursor++
		}
	case termbox.KeySpace:
		c.selected[c.cursor] = !c.selected[c.cursor]
	}
}
