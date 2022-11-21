package sweepstakes

import "fmt"

type Team struct {
	Name    string
	Ranking int
	Odds    float64
	PotID   int
}

type Sweepstaker struct {
	Name  string
	Teams []Team
}

type Pot struct {
	ID    int
	Teams []Team
}

func NewPot(id int) Pot {
	return Pot{
		ID: id,
	}
}

func (p *Pot) ShowPot() {
	fmt.Printf("Pot %d: ", p.ID)
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
