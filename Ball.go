package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Ball struct {
	*tl.Rectangle
	w        int
	h        int
	px       float64 // Previous x
	py       float64 // Previous y
	vx       float64 //Vector X
	vy       float64 //vector Y
	attached bool    //If it's attached to the paddle
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

	b.attached = true
	return b
}

func (b *Ball) Tick(ev tl.Event) {
	if b.attached {

		if ev.Type == tl.EventKey {
			temp_x, _ := b.Position()
			switch ev.Key {
			case tl.KeyArrowRight:
				b.SetPosition(temp_x+1, int(b.py))
			case tl.KeyArrowLeft:
				b.SetPosition(temp_x-1, int(b.py))
			case tl.KeySpace:
				b.attached = false
				b.vx = floatRange(0.15, 0.25) * float64(intRange(-1, 1))
				b.vy = -floatRange(0.15, 0.25)
			}
		}
		return
	}
	b.px += b.vx
	b.py += b.vy
	b.SetPosition(int(b.px), int(b.py))
}

func (b *Ball) Draw(s *tl.Screen) {
	b.Rectangle.Draw(s)
}

func (b *Ball) Collide(p tl.Physical) {

	if r, ok := p.(*BoardLimit); ok {
		// Collision with walls
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

		// Collision with the Paddle
	} else if paddle, ok := p.(*Paddle); ok {
		if b.attached {
			return
		}
		b.SetPosition(int(b.px), int(b.py))

		if paddle.move {
			if b.vx < 0 {
				b.vx = -floatRange(0.15, 0.25)
			} else {
				b.vx = floatRange(0.15, 0.25)
			}

			b.vy = -floatRange(0.15, 0.25)

		}

		// Collision with the Bar on top
	} else if _, ok := p.(*Bar); ok {
		b.SetPosition(int(b.px), int(b.py))
		b.vy = -b.vy
	}
}
