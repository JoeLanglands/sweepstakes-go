package main

import (
	"github.com/JoeLanglands/sweepstakes-go/pkg/sweepstakes"
)

func main() {
	sweepers := []sweepstakes.Sweepstaker{
		{"Joe"},
		{"Dan"},
		{"Dave"},
		{"Harry"},
		{"Sean"},
		{"Jamie"},
		{"Jon"},
	}

	teams := []sweepstakes.Team{
		{"Brazil", 23, 3.5},
		{"Argentina", 1, 5.5},
		{"France", 14, 5.5},
		{"England", 14, 7.5},
		{"Spain", 10, 8},
		{"Germany", 4, 10},
		{"Netherlands", 54, 11},
		{"Portugal", 7, 12},
		{"Belgium", 12, 16},
		{"Denmark", 37, 28},
		{"Uruguay", 7, 40},
		{"Croatia", 16, 50},
		{"Serbia", 62, 80},
		{"Switzerland", 21, 80},
		{"Mexico", 5, 100},
		{"Poland", 38, 100},
		{"Senegal", 44, 100},
		{"USA", 36, 100},
		{"Ecudor", 6, 150},
		{"Wales", 18, 150},
		{"Japan", 13, 200},
		{"Cameroon", 29, 250},
		{"Morocco", 24, 250},
		{"South Korea", 17, 250},
		{"Australia", 6, 500},
		{"Canada", 18, 500},
		{"Ghana", 13, 500},
		{"Iran", 29, 500},
		{"Tunisia", 24, 250},
		{"Costa Rica", 17, 750},
		{"Qatar", 17, 750},
		{"Saudi Arabia", 17, 750},
	}

	sweepstakes.Allocate(sweepers, teams)
}
