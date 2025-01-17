package gamestate

import (
	"fmt"
	"os"
)

type Game struct {
	GameKey  string
	Listener func()
	Handlers map[string]Handler
	GameState
}

type Handler func()

func (g *Game) StartAGame() {
	done := make(chan os.Signal, 1)
	g.Listener = func() {
		go func() {
			for range 2 {
				fmt.Println("\n" + g.GameKey + "\n")
			}
			done <- nil
		}()
		print("\nhello\n")
		<-done
	}
	g.Listener()
}

func (g *Game) CallGame() {
}
