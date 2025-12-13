package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/nsf/termbox-go"
)

func cls() {
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/c", "cls").Run()
	} else {
		exec.Command("clear").Run()
	}
}

func Checkbox(checks []string, title string) ([]bool, error) {
	if err := keyboard.Open(); err != nil {
		return nil, err
	}
	defer keyboard.Close()

	var selected []bool = make([]bool, len(checks))
	currentIndex := 0

	if err := termbox.Init(); err != nil {
		return nil, err
	}
	defer termbox.Close()

	for {
		termbox.SetCursor(0, 0)
		fmt.Print("\033[H\033[2J")
		fmt.Println(title)

		for i := range checks {
			if i == currentIndex {
				fmt.Print("> ")
			} else {
				fmt.Print("  ")
			}
			if selected[i] {
				fmt.Printf("[x] %s\n", checks[i])
			} else {
				fmt.Printf("[ ] %s\n", checks[i])
			}
		}

		if _, key, err := keyboard.GetKey(); err == nil {
			switch key {
			case keyboard.KeySpace:
				selected[currentIndex] = !selected[currentIndex]
			case keyboard.KeyTab:
				currentIndex = (currentIndex + 1) % len(checks)
			case keyboard.KeyEsc:
				return selected, nil
			}
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func Label(content string) {
	fmt.Println(content)
}