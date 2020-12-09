package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

type Bar struct {
	*tl.Rectangle
	level *tl.BaseLevel
}

var BarW = 5
var BarH = 2
var TotalBars = 0

func NewBar(x, y int, level *tl.BaseLevel) *Bar {
	colors := [...]tl.Attr{tl.ColorYellow, tl.ColorRed, tl.ColorWhite, tl.ColorCyan, tl.ColorMagenta, tl.ColorGreen, tl.ColorBlue}

	randomColor := rand.Intn(7)

	b := new(Bar)
	b.level = level
	b.Rectangle = tl.NewRectangle(x, y, 5, 2, colors[randomColor])

	TotalBars++

	return b

}

func (bar *Bar) Draw(s *tl.Screen) {
	bar.Rectangle.Draw(s)
}

func (bar *Bar) Collide(p tl.Physical) {

	if _, ok := p.(*Ball); ok {
		TotalBars--
		bar.level.RemoveEntity(bar)
		if TotalBars == 0 {
			Win()
		}
	}
}
