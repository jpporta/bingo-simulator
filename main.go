package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	strgames := os.Args[1]
	games, err := strconv.Atoi(strgames)
	if err != nil {
		log.Fatal(err)
	}
	strplayers := os.Args[2]
	players, err := strconv.Atoi(strplayers)
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	var wg sync.WaitGroup
	all_games := make([]Game, games)
	for idx := range games {
		wg.Add(1)
		all_games[idx] = CreateGame(players)
		go func() {
			all_games[idx].PlayUntilEnd()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("Took: ", time.Since(start))
	f, err := os.OpenFile(strgames+"_games_with_"+strplayers+"_players.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range all_games {

		res, err := g.Results()
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(res + "\n")); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
