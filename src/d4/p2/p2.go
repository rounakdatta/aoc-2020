package p2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func regexValidator(pattern string, input string) int {
	matched, _ := regexp.MatchString(pattern, input)
	if matched {
		return 1
	} else {
		return 0
	}
}

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

	if points < 14 {
		return false
	}

	validationPoints := 0
	for _, f := range inputFields {
		if f == "" {
			continue
		}
		fSplit := strings.Split(f, ":")
		fieldName := fSplit[0]
		fieldValue := fSplit[1]
		switch fieldName {
		case "byr":
			validationPoints += regexValidator("^(19[2-9][0-9]|200[0-2])$", fieldValue)
		case "iyr":
			validationPoints += regexValidator("^(201[0-9]|2020)$", fieldValue)
		case "eyr":
			validationPoints += regexValidator("^(202[0-9]|2030)$", fieldValue)
		case "hgt":
			validationPoints += regexValidator("^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$", fieldValue)
		case "hcl":
			validationPoints += regexValidator("^#([0-9]|[a-f]){6}$", fieldValue)
		case "ecl":
			validationPoints += regexValidator("^(amb|blu|brn|gry|grn|hzl|oth)$", fieldValue)
		case "pid":
			validationPoints += regexValidator("^([0-9]{9})$", fieldValue)
		}
	}

	return validationPoints == 7
}

func driver() int {
	filename := fmt.Sprintf("p2.input")
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
