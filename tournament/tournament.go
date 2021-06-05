//Package tournament gives us utilities for scoring sports tournaments
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

//TeamRecord is a struct used to store data about a given team and their stats in the league
type TeamRecord struct {
	name   string
	wins   int
	losses int
	draws  int
	played int
	points int
}

//Tally function take game results and tallys up the stats for each team
func Tally(reader io.Reader, writer io.Writer) error {
	//by creating your map with make you then have access to the fields inside TeamRecord
	results := make(map[string]TeamRecord)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		//checking for proper input format
		if scanner.Text() == "" || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		split := strings.Split(scanner.Text(), ";")
		if len(split) != 3 {
			return fmt.Errorf("Input fromat was invalid")
		}
		if split[2] != "win" && split[2] != "loss" && split[2] != "draw" {
			return fmt.Errorf("Input format was invalid")
		}

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
		}
		results[split[0]] = team1
		results[split[1]] = team2
	}

	ordered := sortTable(results)
	writeBoard(ordered, writer)
	return nil
}

//SortTable function takes a map of [string]TeamRecords and returns a sorted slice
func sortTable(results map[string]TeamRecord) []TeamRecord {
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
	return ordered
}

//WriteBoard function writes the table
func writeBoard(ordered []TeamRecord, writer io.Writer) {
	fmt.Fprintf(writer, "%-31s|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for i := range ordered {
		fmt.Fprintf(writer, "%-31s|%3v |%3v |%3v |%3v |%3v\n", ordered[i].name, ordered[i].played, ordered[i].wins, ordered[i].draws, ordered[i].losses, ordered[i].points)
	}
}
