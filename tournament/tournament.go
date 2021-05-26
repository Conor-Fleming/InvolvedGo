package tournament

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type TeamRecord struct {
	name   string
	wins   int
	losses int
	draws  int
	played int
	points int
}

//Results map holds team names
var results = map[string]TeamRecord{}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
	}

	writeBoard(results)
	return nil
}

func writeBoard(results map[string]TeamRecord) {
	writer := bytes.NewBufferString("")
	testStr := "testing writing"
	fmt.Fprint(writer, " some space in between ", testStr)
	fmt.Println(writer.String())
}
