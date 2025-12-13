package uxtm

import "github.com/nsf/termbox-go"

type UI struct {
	name string
	view *View
}

// New creates a new UI instance.
func New(name string) *UI {
	return &UI{
		name: name,
		view: NewView(),
	}
}

// View returns the root view.
func (ui *UI) View() *View {
	return ui.view
}

// Run starts the event loop.
func (ui *UI) Run() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	defer termbox.Close()

	for {
		ui.view.Render(ui.name)

		ev := termbox.PollEvent()
		if ev.Type != termbox.EventKey {
			continue
		}

		switch ev.Key {
		case termbox.KeyEsc, termbox.KeyEnter:
			return nil
		default:
			ui.view.HandleEvent(ev)
		}
	}
}
