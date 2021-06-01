package tournament

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
		team2 := results[split[1]]
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
			break
		default:
			return fmt.Errorf("input formatted incorrectly")
		}

		results[split[0]] = team1
		results[split[1]] = team2
	}

	fmt.Println(results)
	writeBoard(results)
	return nil
}

func writeBoard(results map[string]TeamRecord) {
	header := fmt.Sprintf("Team\t\t\t\t| MP | W | D | L | P\n")
	stats := ""
	for i := range results {
		stats += fmt.Sprintf("Team\t\t\t\t| %v | %v | %v | %v | %v\n", results[i].played, results[i].wins, results[i].draws, results[i].losses, results[i].points)
	}
	header += stats
	writer := bytes.NewBufferString(header)
	writer.String()
}
