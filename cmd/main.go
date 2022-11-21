package main

import (
	"github.com/JoeLanglands/sweepstakes-go/pkg/sweepstakes"
)

func main() {
	sweepers := []sweepstakes.Sweepstaker{
		{Name: "Joe"},
		{Name: "Dan"},
		{Name: "Dave"},
		{Name: "Harry"},
		{Name: "Sean"},
		{Name: "Jamie"},
		{Name: "Jon"},
	}

	teams := []sweepstakes.Team{
		{Name: "Brazil", 23, 3.5, 0},
		{Name: "Argentina", 1, 5.5, 0},
		{Name: "France", 14, 5.5, 0},
		{Name: "England", 14, 7.5, 0},
		{Name: "Spain", 10, 8, 0},
		{Name: "Germany", 4, 10, 0},
		{Name: "Netherlands", 54, 11, 0},
		{Name: "Portugal", 7, 12, 0},
		{Name: "Belgium", 12, 16, 0},
		{Name: "Denmark", 37, 28, 0},
		{Name: "Uruguay", 7, 40, 0},
		{Name: "Croatia", 16, 50, 0},
		{Name: "Serbia", 62, 80, 0},
		{Name: "Switzerland", 21, 80, 0},
		{Name: "Mexico", 5, 100, 0},
		{Name: "Poland", 38, 100, 0},
		{Name: "Senegal", 44, 100, 0},
		{Name: "USA", 36, 100, 0},
		{Name: "Ecudor", 6, 150, 0},
		{Name: "Wales", 18, 150, 0},
		{Name: "Japan", 13, 200, 0},
		{Name: "Cameroon", 29, 250, 0},
		{Name: "Morocco", 24, 250, 0},
		{Name: "South Korea", 17, 250, 0},
		{Name: "Australia", 6, 500, 0},
		{Name: "Canada", 18, 500, 0},
		{Name: "Ghana", 13, 500, 0},
		{Name: "Iran", 29, 500, 0},
		{Name: "Tunisia", 24, 250, 0},
		{Name: "Costa Rica", 17, 750, 0},
		{Name: "Qatar", 17, 750, 0},
		{Name: "Saudi Arabia", 17, 750, 0},
	}

	sweepstakes.Allocate(sweepers, teams)
}
