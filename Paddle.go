package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Paddle struct {
	*tl.Rectangle
	w    int
	h    int
	px   int // Previous x
	py   int // Previous y
	move bool
}

func NewPaddle(x, y int) *Paddle {
	s := new(Paddle)
	s.h = 1
	s.w = 10
	s.Rectangle = tl.NewRectangle(x, y, s.w, s.h, tl.ColorWhite)
	return s
}

func (m *Paddle) Tick(ev tl.Event) {

	if ev.Type == tl.EventKey {
		temp_x, _ := m.Position()
		if temp_x != m.px {
			m.move = true
		} else {
			m.move = false
		}
		m.px, m.py = m.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			m.SetPosition(m.px+1, m.py)
		case tl.KeyArrowLeft:
			m.SetPosition(m.px-1, m.py)
		}
	}
}

func (m *Paddle) Draw(s *tl.Screen) {
	m.Rectangle.Draw(s)
}

func (m *Paddle) Collide(p tl.Physical) {

	if r, ok := p.(*BoardLimit); ok {
		if r.limitType == SID {
			m.SetPosition(m.px, m.py)
		}
	}
}
