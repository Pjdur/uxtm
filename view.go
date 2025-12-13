package uxtm

import "github.com/nsf/termbox-go"

type View struct {
	components []Component
	focus      int
}

func NewView() *View {
	return &View{}
}

func (v *View) Add(c ...Component) *View {
	v.components = append(v.components, c...)
	return v
}

// Render draws all components.
func (v *View) Render(title string) {
	termbox.Clear(0, 0)

	drawString(
		0, 0,
		termbox.ColorCyan|termbox.AttrBold,
		0,
		title,
	)

	y := 2
	for i, c := range v.components {
		c.Render(0, y, i == v.focus)
		y += c.Height() + 1
	}

	drawString(
		0, y+1,
		termbox.ColorGreen,
		0,
		"Tab Switch | Arrows Move | Space Toggle | Enter Confirm | Esc Exit",
	)

	termbox.Flush()
}


func (v *View) HandleEvent(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyTab:
		v.focus = (v.focus + 1) % len(v.components)
	default:
		if len(v.components) > 0 {
			v.components[v.focus].HandleEvent(ev)
		}
	}
}
