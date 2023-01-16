package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"
)

type Record struct {
	mp, wins, draws, losses, points int
}

type Contenders map[string]Record

func Tally(in io.Reader, out io.Writer) error {
	// reads in match data and formats the output

	readMatches := bufio.NewScanner(in)
	matches := []string{}
	teams := Contenders{}

	// read each line from input as a 'match'
	for readMatches.Scan() {
		matches = append(matches, readMatches.Text())
	}

	for _, m := range matches {
		// skip new lines
		if m == "" {
			continue
		}

		// skip "comments"
		if strings.HasPrefix(m, "#") {
			continue
		}

		parts := strings.Split(m, ";")

		// Bail if match data is no good
		if len(parts) != 3 {
			return errors.New(fmt.Sprintf("Malformed match data: %v", m))
		}

		// Check the first two string parts and add to team list
		for i := 0; i < 2; i++ {
			teams.addTeam(parts[i])
		}

		outcome, err := tallyResult(parts[2])
		if err != nil {
			return err
		}

		homeTeam := teams[parts[0]] // "left team"
		visitingTeam := teams[parts[1]] // "right team"

		homeTeam.mp += 1
		visitingTeam.mp += 1
		switch outcome {
		case 3:
			homeTeam.wins += 1
			homeTeam.points += 3

			visitingTeam.losses += 1
		case 1:
			homeTeam.draws += 1
			visitingTeam.draws += 1

			homeTeam.points += 1
			visitingTeam.points += 1
		case 0:
			visitingTeam.wins += 1
			visitingTeam.points += 3

			homeTeam.losses += 1
		}

		teams[parts[0]] = homeTeam
		teams[parts[1]] = visitingTeam
	}

	rank := teams.rank()

	w := tabwriter.NewWriter(out, 0, 0,1, ' ', tabwriter.Debug)

	_,_ = fmt.Fprintln(w, "Team\t MP\t  W\t  D\t  L\t  P")
	for _, team := range rank {
		_,_ = fmt.Fprintf(w, "%v       \t  %v\t  %v\t  %v\t  %v\t  %v\n",
			team,
			teams[team].mp,
			teams[team].wins,
			teams[team].draws,
			teams[team].losses,
			teams[team].points)
	}

	w.Flush()
	return nil
}

func tallyResult(s string) (int, error) {
	// tally the results

	switch s {
	case "win":
		return 3, nil
	case "draw":
		return 1, nil
	case "loss":
		return 0, nil
	default:
		return -1, errors.New(fmt.Sprintf("Unable to determine outcome: %v", s))
	}
}

func (c Contenders) addTeam (s string) {
	// if s matches what would normally be considered a team
	// add it to the list of passed in Contenders with append()
	// if it doesn't already exist else return error

	_, present := c[s]
	if present {
		return
	}

	c[s] = Record{0,0,0 ,0, 0}
}

func (c Contenders) rank() []string {
	var rank []string

	for k := range c {
		rank = append(rank, k)
	}

	sort.Slice(rank, func (i, j int) bool {
		lpoints := c[rank[i]].points
		rpoints := c[rank[j]].points

		if lpoints == rpoints {
			return sort.StringsAreSorted([]string{rank[i], rank[j]})
		}

		return  lpoints > rpoints
	})

	return rank
}
