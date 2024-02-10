package main

import (
	"time"

	keyboard "github.com/sago35/tinygo-keyboard"
)

func main() {
	d := keyboard.New()

	for {
		d.Mouse.Move(100, 0)
		time.Sleep(1 * time.Second)
		d.Mouse.Move(0, 100)
		time.Sleep(1 * time.Second)
		d.Mouse.Move(-100, 0)
		time.Sleep(1 * time.Second)
		d.Mouse.Move(0, -100)
		time.Sleep(1 * time.Second)
	}
}
