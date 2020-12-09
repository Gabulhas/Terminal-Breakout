package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"os"
)

var Config *GameConfig

func main() {
	g := tl.NewGame()
	g.SetDebugOn(true)
	g.Screen().SetFps(60)

	Config = new(GameConfig)
	Config.w = 60
	Config.h = 30

	initGame(g, Config.w, Config.h)
	g.Start()
}

func initGame(g *tl.Game, w, h int) {
	l := tl.NewBaseLevel(tl.Cell{})
	g.Screen().SetLevel(l)
	initBoard(Config.w, Config.h, l)

	l.AddEntity(NewPaddle((w+1)/2, h-1))
	l.AddEntity(NewBall(w/2, h/2))
	addBars(w, h, l)
}

func addBars(w, h int, l *tl.BaseLevel) {

	//for lines
	for a := 0; a < 4; a++ {
		for i := 1; i < w+1; i = i + BarW {
			l.AddEntity(NewBar(i, (a*BarH)+1, l))
		}
	}
}

func initBoard(w, h int, l *tl.BaseLevel) {
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

	fmt.Printf("YOU LOST")
	os.Exit(-1)
}
func Win() {

	fmt.Printf("YOU WON")
	os.Exit(-1)
}
