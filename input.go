package uxtm

import (
	"github.com/pjdur/uxtm/style"
	"github.com/pjdur/uxtm/term"
)

type Input struct {
	title       string
	value       []rune
	cursor      int
	style       style.Style
	cursorStyle style.Style
}

func NewInput() *Input {
	return &Input{
		value:  []rune{},
		cursor: 0,
		style:  style.New(),
		cursorStyle: style.New().
			Reverse(true),
	}
}

func (i *Input) SetTitle(t string) *Input {
	i.title = t
	return i
}

func (i *Input) SetStyle(s style.Style) *Input {
	i.style = s
	return i
}

func (i *Input) SetCursorStyle(s style.Style) *Input {
	i.cursorStyle = s
	return i
}
func (i *Input) HandleEvent(e Event) bool {
	switch ev := e.(type) {
	case KeyEvent:
		switch ev.Key {
		case KeyLeft:
			if i.cursor > 0 {
				i.cursor--
			}
		case KeyRight:
			if i.cursor < len(i.value) {
				i.cursor++
			}
		case KeyBackspace:
			if i.cursor > 0 {
				i.value = append(i.value[:i.cursor-1], i.value[i.cursor:]...)
				i.cursor--
			}
		default:
			if ev.Rune != 0 {
				i.value = append(
					i.value[:i.cursor],
					append([]rune{ev.Rune}, i.value[i.cursor:]...)...,
				)
				i.cursor++
			}
		}
		return true
	}
	return false
}

