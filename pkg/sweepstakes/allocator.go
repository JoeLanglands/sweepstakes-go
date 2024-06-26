package sweepstakes

import (
	"math/rand"
	"sort"
)

type RankBy int

const (
	RankByOdds RankBy = iota
	RankByRanking
)

// MakePools takes a slice of teams and a number of pools and
// organises the teams into pools based on the odds of each team.
// numPools is derived from the number of sweepstakers. If the number
// of teams is not divisible by the number of pools, the lowest ranked/odds
// teams are not allocated to a pool.
func MakePools(teams []Team, numPools int, rankBy RankBy) ([]Pool, error) {
	pools := make([]Pool, numPools)
	poolSize := len(teams) / numPools
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Odds < teams[j].Odds
	})
	teamsSorted := sortTeams(teams, rankBy)

	for i := range pools {
		pools[i].ID = i + 1
		pools[i].Teams = teamsSorted[i*poolSize : (i+1)*poolSize]
	}

	return pools, nil
}

func sortTeams(teams []Team, rankBy RankBy) []Team {
	if rankBy == RankByOdds {
		sort.Slice(teams, func(i, j int) bool {
			return teams[i].Odds < teams[j].Odds
		})
	} else {
		sort.Slice(teams, func(i, j int) bool {
			return teams[i].Ranking < teams[j].Ranking
		})
	}
	return teams
}

// Allocate takes a slice of teams and a slice of sweepstakers and allocates
// teams to sweepstakers based on the odds of each team. Teams are first divided
// into pools based on the number of sweepstakers.
func Allocate(s []Sweepstaker, teams []Team, rankBy RankBy) ([]Sweepstaker, error) {

	pools, err := MakePools(teams, len(teams)/len(s), rankBy)
	if err != nil {
		return nil, err
	}

	if len(pools) == 1 {
		initSinglePool(&pools, len(s))
	}

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

	return s, nil
}

// initSinglePool will truncate the teams of the single pool to match the number of players
// this ensures that every sweeper gets a team from the top most rankings.
func initSinglePool(pools *[]Pool, numSweepstakers int) {
	(*pools)[0].Teams = (*pools)[0].Teams[:numSweepstakers]
}
