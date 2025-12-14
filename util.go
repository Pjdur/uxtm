package uxtm

import "github.com/nsf/termbox-go"

func drawString(x, y int, fg, bg termbox.Attribute, msg string) {
	for i, r := range msg {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}

func (i *Input) Draw(t *term.Term, x, y, w, h int) {
	// Draw title
	if i.title != "" {
		t.DrawText(x, y, i.title, style.New())
		y++
	}

	for idx, r := range i.value {
		if idx == i.cursor {
			t.DrawRune(x+idx, y, r, i.cursorStyle)
		} else {
			t.DrawRune(x+idx, y, r, i.style)
		}
	}

	if i.cursor == len(i.value) {
		t.DrawRune(x+i.cursor, y, ' ', i.cursorStyle)
	}
}
