package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamRecord struct {
	name   string
	wins   int
	losses int
	draws  int
	played int
	points int
}

func Tally(reader io.Reader, writer io.Writer) error {
	//by creating your map with make you then have access to the fields inside TeamRecord
	results := make(map[string]TeamRecord)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if scanner.Text() == "" || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		//fmt.Println(scanner.Text())
		split := strings.Split(scanner.Text(), ";")

		team1 := results[split[0]]
		team1.name = split[0]
		team2 := results[split[1]]
		team2.name = split[1]
		//switch statement handling the tallying of points and games based off results of matches
		switch split[2] {
		case "win":
			team1.wins++
			team1.played++
			team2.played++
			team2.losses++
			team1.points += 3
			break
		case "loss":
			team1.losses++
			team1.played++
			team2.played++
			team2.wins++
			team2.points += 3
			break
		case "draw":
			team1.played++
			team2.played++
			team1.points++
			team2.points++
			team1.draws++
			team2.draws++
			break
		default:
			return fmt.Errorf("input formatted incorrectly")
		}

		results[split[0]] = team1
		results[split[1]] = team2
	}

	ordered := make([]TeamRecord, 0, len(results))
	for _, v := range results {
		ordered = append(ordered, v)
	}
	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].points == ordered[j].points {
			return ordered[i].name < ordered[j].name
		}
		return ordered[i].points > ordered[j].points
	})

	writeBoard(ordered, writer)
	return nil
}

func writeBoard(ordered []TeamRecord, writer io.Writer) {
	fmt.Fprintf(writer, "%-31s|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for i := range ordered {
		fmt.Fprintf(writer, "%-31s|%3v |%3v |%3v |%3v |%3v\n", ordered[i].name, ordered[i].played, ordered[i].wins, ordered[i].draws, ordered[i].losses, ordered[i].points)
	}
}

/*func sortTable(results map[string]TeamRecord) {
	order := make([]string, 0, len(results))
	for k, v := range results {
		v.name = k
		order = append(order, k)
	}
	sort.Strings(order)
}*/
