package tournament

import (
	"fmt"
	"io"
)

type Team struct {
	name   string
	wins   int
	losses int
	draws  int
	played int
	points int
}

func newTeam(name string) *Team {
	t := Team{name: name}
	t.wins = 0
	t.losses = 0
	t.draws = 0
	t.played = 0
	t.points = 0
	return &t
}

func Tally(input io.Reader, output io.Writer) error {
	AA := newTeam("AA")
	fmt.Println(AA)

	return nil
}
