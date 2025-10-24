package tournament

import (
	"io"
	"bufio"
	"fmt"
	"strings"
	"errors"
	"sort"
)

type TeamTally struct {
	mp int
	w int
	d int
	l int
	p int
}

type Team struct {
	name string
	tally *TeamTally
}

func Tally(reader io.Reader, writer io.Writer) error {
	// create buffer scanner.
	scanner := bufio.NewScanner(reader)

	// map team names to respective tallies.
	scoresheet := map[string]*TeamTally{}

	// read input line by line.
	for scanner.Scan() {
		line := scanner.Text()
		// ignore comments and empty lines.
		if len(line) == 0 || line[0] == '#' { 
			continue
		}

		// split line into parts.
		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return errors.New("invalid row")
		}

		// check that teams are different.
		if parts[0] == parts[1] {
			return errors.New("team cannot play against itself")
		}

		// check that result is valid.
		res := parts[2]
		if res != "win" && res != "draw" && res != "loss" {
			return errors.New("invalid game result")
		}

		// get team tallies, create if needed.
		t1, ok := scoresheet[parts[0]]; if !ok {
			scoresheet[parts[0]] = &TeamTally{mp: 0, w: 0, d: 0, l: 0, p: 0}
			t1 = scoresheet[parts[0]]
		}
		t2, ok := scoresheet[parts[1]]; if !ok {
			scoresheet[parts[1]] = &TeamTally{mp: 0, w: 0, d: 0, l: 0, p: 0}
			t2 = scoresheet[parts[1]]
		}

		// increment matches played for both teams.
		t1.mp++
		t2.mp++
		
		// update team tallies depending on the result.
		switch res {
		case "win":
			t1.w += 1
			t1.p += 3

			t2.l += 1
		case "draw":
			t1.d += 1
			t1.p += 1

			t2.d += 1
			t2.p += 1
		case "loss":
			t1.l += 1

			t2.w += 1
			t2.p += 3
		}
	}

	// construct Team structs.
	teams := []Team{}
	for n, t := range scoresheet {
		teams = append(teams, Team{
			name: n,
			tally: t,
		})
	}

	// sort teams by points.
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].tally.p == teams[j].tally.p {
			return teams[i].name < teams[j].name
		}
		return teams[i].tally.p > teams[j].tally.p
	})

	// write the result.
	writer.Write([]byte(
		fmt.Sprintf("Team                           | MP |  W |  D |  L |  P\n"),
	))
	for _, t := range teams {
		writer.Write([]byte(
			fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n", t.name, t.tally.mp, t.tally.w, t.tally.d, t.tally.l, t.tally.p),
		))
	}

	return nil
}
