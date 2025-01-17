package gamestate

import (
	"fmt"
	"os"
	"seesharpsi/gobsmack/assets"
)

type Game struct {
	GameKey       string
	Listener      func()
	Actions       map[string]func()
	QueuedActions map[string]func() bool
	Players       map[string]assets.Player
	GameState
}

func (g *Game) StartGame() {
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

// donewith should be index of item in player's inventory. Could make templ insert it into the htmx request
func (g *Game) QueueAction(player string, atype string, donewith int) bool {
	cplayer := g.Players[player]
	if !cplayer.TurnOver {
		switch atype {
		case "use":
			g.QueuedActions[player] = cplayer.Items[donewith].Use
		case "pass":
			g.QueuedActions[player] = cplayer.Pass
		case "move":
			switch donewith {
			case 0:
				cplayer.Move("north")
			case 1:
				cplayer.Move("south")
			case 2:
				cplayer.Move("east")
			case 3:
				cplayer.Move("west")
			}
		}
		cplayer.TurnOver = true
		g.Players[player] = cplayer
		return true
	}
	return false
}

func Use(player string) {
	print("pew pew")
}

//action/use/gun/
