package command

import termbox "github.com/nsf/termbox-go"

type CommandSet struct {
	Cmd Command
	Chr rune
}

type Command int

const (
	QuitApp Command = 0
	DmyCmd  Command = 1
	NilCmd  Command = 2
)

func NewCommandSet() *CommandSet {
	return &CommandSet{}
}

func (c *CommandSet) Parse(event termbox.Event) error {
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyEsc:
			c.Cmd = QuitApp
		case termbox.KeyEnter:
		case termbox.KeyArrowUp:
		case termbox.KeyArrowDown:
		case termbox.KeyArrowLeft:
		case termbox.KeyArrowRight:
		case termbox.KeyBackspace, termbox.KeyBackspace2:
		case termbox.KeyCtrlS:
		default:
			if event.Ch != 0 {
			}
		}
	case termbox.EventError:
		return event.Err
	}
	return nil
}