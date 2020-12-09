package main

import tl "github.com/JoelOtter/termloop"

type BoardLimit struct {
	*tl.Rectangle
	limitType LimitType
}

type LimitType int

const (
	TOP LimitType = iota
	BOT
	SID
)

func NewLimit(x, y int, limitType LimitType) *BoardLimit {

	var color tl.Attr
	color = tl.ColorWhite

	l := new(BoardLimit)

	l.limitType = limitType

	l.Rectangle = tl.NewRectangle(x, y, 1, 1, color)
	return l
}

