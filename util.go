package uxtm

import "github.com/nsf/termbox-go"

func drawString(x, y int, fg, bg termbox.Attribute, msg string) {
	for i, r := range msg {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}
func drawInput(x, y int, title string, value string, cursorPos int) {
 // Draw the title
 drawString(x, y, termbox.ColorWhite, termbox.ColorDefault, ">"+title+" ; ")

 // Calculate input box start position
 inputStart := x + len(">"+title+" ; ")

 // Draw the input box (red background)
 inputLen := 16 // length of input box
 for i := 0; i < inputLen; i++ {
  ch := ' '
  fg := termbox.ColorWhite
  bg := termbox.ColorRed
  if i < len(value) {
   ch = rune(value[i])
  }
  termbox.SetCell(inputStart+i, y, ch, fg, bg)
 }

 // Draw the right border
 termbox.SetCell(inputStart+inputLen, y, '|', termbox.ColorWhite, termbox.ColorDefault)

 // Draw the cursor (as an underline)
 if cursorPos >= 0 && cursorPos < inputLen {
  termbox.SetCell(inputStart+cursorPos, y, '_', termbox.ColorWhite|termbox.AttrBold, termbox.ColorRed)
 }
}