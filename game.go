package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Game struct {
	Players   []Player
	available []int
	drawn     []int
	round     int
	winner    int
}

func CreateGame(numPlayers int) Game {
	players := make([]Player, numPlayers)
	for idx := range numPlayers {
		players[idx] = CreatePlayer()
	}
	balls := rand.Perm(75)
	for idx := range balls {
		balls[idx]++
	}
	return Game{
		Players:   players,
		available: balls,
		drawn:     make([]int, 75),
		winner:    -1,
	}
}

func (g *Game) NewRound() bool {
	ball := g.available[g.round]
	g.round++

	done := false
	for idx := range g.Players {
		if done = g.Players[idx].CheckNewBall(ball); done {
			g.winner = idx
			break
		}
	}
	return done
}

func (g Game) GetWinner() (Player, error) {
	if g.winner == -1 {
		return Player{}, fmt.Errorf("No Winner yet")
	}
	return g.Players[g.winner], nil
}

func (g *Game) PlayUntilEnd() {
	fallback := 0
	for {
		if done := g.NewRound(); done {
			break
		}
		if fallback > 100 {
			break
		}
		fallback++
	}
}

func (g Game) Results() (string, error) {
	if g.winner == -1 {
		return "", fmt.Errorf("No Winner yet")
	}
	win := g.Players[g.winner]
	line_win := "horizontal"
	if win.Own != 1 {
		line_win = "vertical"
	}
	return strconv.Itoa(g.round) + ", " + line_win, nil
}
