# uxtm

A simple and lightweight Go library for building Terminal User Interfaces (TUIs) using [termbox-go](https://github.com/nsf/termbox-go).

## Features

- **Component-based architecture**: Build TUIs using reusable components like checkboxes, sliders, and input fields.
- **Event handling**: Built-in support for keyboard events and focus management.
- **Simple API**: Easy to use with a minimalistic design.
- **Cross-platform**: Works on any terminal that supports termbox-go.

## Installation

```bash
go get github.com/Pjdur/uxtm
```

## Usage

Here's a basic example of creating a TUI with a checkbox and a slider:

```go
package main

import (
    "log"

    "github.com/Pjdur/uxtm"
)

func main() {
    ui := uxtm.New("My TUI App")

    checkbox := uxtm.NewCheckbox("Options", []string{"Option 1", "Option 2", "Option 3"})
    slider := uxtm.NewSlider("Volume", 0, 100, 50)

    ui.View().Add(checkbox, slider)

    if err := ui.Run(); err != nil {
        log.Fatal(err)
    }
}
```

## Controls

- **Tab**: Switch focus between components
- **Arrow Keys**: Navigate within components (e.g., move cursor in checkboxes, adjust sliders)
- **Space**: Toggle selections (e.g., check/uncheck checkboxes)
- **Enter**: Confirm and exit
- **Esc**: Exit without confirming

## API Overview

### UI

- `New(name string) *UI`: Creates a new UI instance.
- `View() *View`: Returns the root view.
- `Run() error`: Starts the event loop.

### View

- `Add(components ...Component) *View`: Adds components to the view.

### Components

All components implement the `Component` interface:

- `Render(x, y int, focused bool)`: Renders the component.
- `HandleEvent(ev termbox.Event)`: Handles input events.
- `Height() int`: Returns the component's height.

#### Checkbox

- `NewCheckbox(title string, options []string) *Checkbox`: Creates a new checkbox component.

#### Slider

- `NewSlider(label string, min, max, value int) *Slider`: Creates a new slider component.

#### Input

- `NewInput(title string) *Input`: Creates a new input component with a title.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
