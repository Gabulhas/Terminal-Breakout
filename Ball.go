package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Ball struct {
	*tl.Rectangle
	w  int
	h  int
	px float64 // Previous x
	py float64 // Previous y
	vx float64 //Vector X
	vy float64 //vector Y
}

func NewBall(x, y int) *Ball {

	b := new(Ball)
	b.h = 1
	b.w = 1

	b.vx = 0.2
	b.vy = 0.2

	b.px = float64(x)
	b.py = float64(y)

	b.Rectangle = tl.NewRectangle(x, y, b.w, b.h, tl.ColorCyan)
	return b
}

func (b *Ball) Tick(ev tl.Event) {
	b.px += b.vx
	b.py += b.vy
	b.SetPosition(int(b.px), int(b.py))
}

func (b *Ball) Draw(s *tl.Screen) {
	b.Rectangle.Draw(s)
}

func (b *Ball) Collide(p tl.Physical) {

	if r, ok := p.(*BoardLimit); ok {
		switch r.limitType {
		case TOP:
			b.SetPosition(int(b.px), int(b.py))
			b.vy = -b.vy
			break
		case BOT:
			Lose()
			break
		case SID:
			b.SetPosition(int(b.px), int(b.py))
			b.vx = -b.vx
			break
		}
	} else if _, ok := p.(*Paddle); ok {
		// Collision with walls
		b.SetPosition(int(b.px), int(b.py))
		b.vy = -b.vy
	} else if _, ok := p.(*Bar); ok {
		b.SetPosition(int(b.px), int(b.py))
		b.vy = -b.vy
	}
}
