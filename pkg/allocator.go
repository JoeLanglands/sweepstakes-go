package sweepstakes

import (
	"math/rand"
	"sort"
	"time"
)

func Allocate(s []Sweepstaker, teams []Team) error {
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Odds < teams[j].Odds
	})

	pot := NewPot(1)
	numPots := 1

	pots := []Pot{pot}

	for i, t := range teams {
		t.Pot = pot
		pot.Teams = append(pot.Teams, t)

		if i%len(s) == 0 {
			numPots++
			pot = NewPot(numPots)
			pots = append(pots, pot)
		}
	}

	rand.Seed(time.Now().UnixNano())

	for _, p := range pots {
		p.ShowPot()

		rand.Shuffle(len(p.Teams), func(i, j int) {
			s[i], s[j] = s[j], s[i]
		})
		n := 0

		for _, s := range s {
			if n < len(p.Teams) {
				s.Teams = append(s.Teams, p.Teams[n])
				n++
			}
		}
	}

	for _, s := range s {
		s.ShowSweepstaker()
	}

	return nil
}
