package p1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPassportValid(input string) bool {
	inputFields := strings.Split(input, " ")
	points := 0
	for _, f := range inputFields {
		fieldName := strings.Split(f, ":")[0]
		switch fieldName {
		case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid":
			points += 2
		case "cid":
			points += 1
		}
	}

	return points >= 14
}

func driver() int {
	filename := fmt.Sprintf("p1.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	count := 0

	bufReader := bufio.NewScanner(fp)
	inputInBuffer := ""
	for bufReader.Scan() {
		readInput := bufReader.Text()
		if readInput != "" {
			inputInBuffer = fmt.Sprintf("%s %s", inputInBuffer, readInput)
		} else {
			if isPassportValid(inputInBuffer) {
				count += 1
			}
			inputInBuffer = ""
		}
	}
	return count
}
