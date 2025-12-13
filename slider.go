package uxtm

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Slider struct {
	label string
	min   int
	max   int
	value int
	width int
}

// NewSlider creates a slider component.
func NewSlider(label string, min, max, value int) *Slider {
	return &Slider{
		label: label,
		min:   min,
		max:   max,
		value: value,
		width: 16,
	}
}

func (s *Slider) Height() int {
	return 1
}

func (s *Slider) Render(x, y int, focused bool) {
	pos := int(
		float64(s.value-s.min) /
			float64(s.max-s.min) *
			float64(s.width),
	)

	bar := ""
	for i := 0; i < s.width; i++ {
		if i == pos {
			bar += "_"
		} else {
			bar += "_"
		}
	}

	prefix := "  "
	color := termbox.ColorWhite
	if focused {
		prefix = "> "
		color = termbox.ColorYellow | termbox.AttrBold
	}

	drawString(
		x,
		y,
		color,
		0,
		fmt.Sprintf(
			"%s%s [%d %s %d]: %d",
			prefix,
			s.label,
			s.min,
			bar,
			s.max,
			s.value,
		),
	)
}

func (s *Slider) HandleEvent(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowLeft:
		if s.value > s.min {
			s.value--
		}
	case termbox.KeyArrowRight:
		if s.value < s.max {
			s.value++
		}
	}
}
