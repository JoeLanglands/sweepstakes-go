package sweepstakes

import (
	"fmt"
	"math/rand"
	"sort"
)

// MakePools takes a slice of teams and a number of pools and
// organises the teams into pools based on the odds of each team.
// numPools is derived from the number of sweepstakers. If the number
// of teams is not divisible by the number of pools, the lowest ranked/odds
// teams are not allocated to a pool.
func MakePools(teams []Team, numPools int) []Pool {
	pools := make([]Pool, numPools)
	poolSize := len(teams) / numPools
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Odds < teams[j].Odds
	})

	for i := range pools {
		pools[i].ID = i + 1
		pools[i].Teams = teams[i*poolSize : (i+1)*poolSize]
	}

	return pools
}

// Allocate takes a slice of teams and a slice of sweepstakers and allocates
// teams to sweepstakers based on the odds of each team. Teams are first divided
// into pools based on the number of sweepstakers.
func Allocate(s []Sweepstaker, teams []Team) error {

	pools := MakePools(teams, len(teams)/len(s))

	for _, p := range pools {
		p.ShowPool()

		rand.Shuffle(len(p.Teams), func(i, j int) {
			p.Teams[i], p.Teams[j] = p.Teams[j], p.Teams[i]
		})
		n := 0

		for i := range s {
			if n < len(p.Teams) {
				s[i].Teams = append(s[i].Teams, p.Teams[n])
				n++
			}
		}
	}

	fmt.Println()
	fmt.Println("Allocated teams:")
	for _, s := range s {
		s.ShowSweepstaker()
	}

	return nil
}
