package sweepstakes

import "fmt"

type Team struct {
	Name    string
	Ranking int
	Odds    float64
}

type Sweepstaker struct {
	Name  string
	Teams []Team
}

type Pool struct {
	ID    int
	Teams []Team
}

func NewPool(id int) Pool {
	return Pool{
		ID: id,
	}
}

func (p *Pool) ShowPool() {
	fmt.Printf("Pool %d: ", p.ID)
	for _, t := range p.Teams {
		fmt.Printf("%s ", t.Name)
	}
	fmt.Println()
}

func (s *Sweepstaker) ShowSweepstaker() {
	fmt.Printf("%s: ", s.Name)
	for _, t := range s.Teams {
		fmt.Printf("%s ", t.Name)
	}
	fmt.Println()
}
