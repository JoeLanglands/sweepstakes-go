package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JoeLanglands/sweepstakes-go/pkg/sweepstakes"
	"github.com/fatih/color"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sweepers := []sweepstakes.Sweepstaker{
		{Name: "Joe"},
		{Name: "Dan"},
		{Name: "Tim"},
		{Name: "Harry"},
		{Name: "Connor"},
		{Name: "James"},
		{Name: "John"},
		{Name: "Glenn"},
	}

	teams := []sweepstakes.Team{
		{Name: "Brazil", Ranking: 23, Odds: 3.5},
		{Name: "Argentina", Ranking: 1, Odds: 5.5},
		{Name: "France", Ranking: 14, Odds: 5.5},
		{Name: "England", Ranking: 14, Odds: 7.5},
		{Name: "Spain", Ranking: 10, Odds: 8},
		{Name: "Germany", Ranking: 4, Odds: 10},
		{Name: "Netherlands", Ranking: 54, Odds: 11},
		{Name: "Portugal", Ranking: 7, Odds: 12},
		{Name: "Belgium", Ranking: 12, Odds: 16},
		{Name: "Denmark", Ranking: 37, Odds: 28},
		{Name: "Uruguay", Ranking: 7, Odds: 40},
		{Name: "Croatia", Ranking: 16, Odds: 50},
		{Name: "Serbia", Ranking: 62, Odds: 80},
		{Name: "Switzerland", Ranking: 21, Odds: 80},
		{Name: "Mexico", Ranking: 5, Odds: 100},
		{Name: "Poland", Ranking: 38, Odds: 100},
		{Name: "Senegal", Ranking: 44, Odds: 100},
		{Name: "USA", Ranking: 36, Odds: 100},
		{Name: "Ecudor", Ranking: 6, Odds: 150},
		{Name: "Wales", Ranking: 18, Odds: 150},
		{Name: "Japan", Ranking: 13, Odds: 200},
		{Name: "Cameroon", Ranking: 29, Odds: 250},
		{Name: "Morocco", Ranking: 24, Odds: 250},
		{Name: "South Korea", Ranking: 17, Odds: 250},
		{Name: "Australia", Ranking: 6, Odds: 500},
		{Name: "Canada", Ranking: 18, Odds: 500},
		{Name: "Ghana", Ranking: 13, Odds: 500},
		{Name: "Iran", Ranking: 29, Odds: 500},
		{Name: "Tunisia", Ranking: 24, Odds: 250},
		{Name: "Costa Rica", Ranking: 17, Odds: 750},
		{Name: "Qatar", Ranking: 17, Odds: 750},
		{Name: "Saudi Arabia", Ranking: 17, Odds: 750},
	}

	color.Set(color.FgRed)
	fmt.Println("----------------")
	fmt.Printf("Number of sweepstakers: %d\n", len(sweepers))
	fmt.Printf("Number of teams: %d\n", len(teams))
	fmt.Printf("Number of leftover teams: %d\n", len(teams)%len(sweepers))

	color.Set(color.FgRed)
	fmt.Println("----------------")
	fmt.Println()

	s, err := sweepstakes.Allocate(sweepers, teams, sweepstakes.RankByOdds)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	color.Set(color.FgRed)
	fmt.Println("Allocated teams:")
	fmt.Println("----------------")
	for _, s := range s {
		s.ShowSweepstaker()
	}
}
