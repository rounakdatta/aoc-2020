package p1

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getColumn(ticket string) int {
	min := 0
	max := 7
	for _, token := range ticket[7:] {
		switch token {
		case 'L':
			max = max - int(math.Ceil(float64(max - min) / 2))
		case 'R':
			min = min + int(math.Ceil(float64(max - min) / 2))
		}
	}
	return min
}

func getRow(ticket string) int {
	min := 0
	max := 127
	for _, token := range ticket[:7] {
		switch token {
		case 'F':
			max = max - int(math.Ceil(float64(max - min) / 2))
		case 'B':
			min = min + int(math.Ceil(float64(max - min) / 2))
		}
	}
	return min
}

func getID(ticket string) int {
	return getRow(ticket) * 8 + getColumn(ticket)
}

func driver() int {
	filename := fmt.Sprintf("p1.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	maxID := -1

	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		input := bufReader.Text()
		id := getID(input)
		if id > maxID {
			maxID = id
		}
	}

	return maxID
}
