package main

import (
	"fmt"
	"os"

	tl "github.com/JoelOtter/termloop"
)

var Config *GameConfig
var l *tl.BaseLevel
var paddle *Paddle
var ball *Ball

func main() {
	g := tl.NewGame()
	g.SetDebugOn(true)
	g.Screen().SetFps(50)

	Config = new(GameConfig)
	Config.w = 60
	Config.h = 30
	Config.lives = 3

	initGame(g, Config.w, Config.h)
	g.Start()
}

func initGame(g *tl.Game, w, h int) {
	l = tl.NewBaseLevel(tl.Cell{})
	g.Screen().SetLevel(l)
	initBoard(Config.w, Config.h)
	initPlayer(w, h)
	addBars(w, h)
}

func initPlayer(w, h int) {
	paddle = NewPaddle(w/2, h-1)
	l.AddEntity(paddle)
	ball = NewBall((w+6)/2, h-2)
	l.AddEntity(ball)
}

func addBars(w, h int) {

	//for lines
	for a := 0; a < 4; a++ {
		for i := 1; i < w+1; i = i + BarW {
			l.AddEntity(NewBar(i, (a*BarH)+1, l))
		}
	}
}

func initBoard(w, h int) {
	// + 2 because of borders
	for i := 0; i < w+2; i++ {

		for j := 0; j < h+2; j++ {
			if j == h+1 {
				l.AddEntity(NewLimit(i, j, BOT))
			} else if j == 0 {
				l.AddEntity(NewLimit(i, j, TOP))
			} else if i == w+1 || i == 0 {
				l.AddEntity(NewLimit(i, j, SID))
			}
		}
	}

}

func Lose() {
	if Config.lives == 0 {
		fmt.Printf("YOU LOST")
		os.Exit(-1)
	} else {
		Config.lives = Config.lives - 1
		l.RemoveEntity(ball)
		l.RemoveEntity(paddle)
		initPlayer(Config.w, Config.h)
	}

}
func Win() {

	fmt.Printf("YOU WON")
	os.Exit(-1)
}
