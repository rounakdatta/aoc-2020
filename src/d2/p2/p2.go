package p2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isPolicySatisfied(min int32, max int32, c int32, pwd string) bool {
	prop1 := int32(pwd[min]) == c
	prop2 := int32(pwd[max]) == c

	return prop1 != prop2
}

func driver() int {
	filename := fmt.Sprintf("p2.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	valid := 0

	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		input := bufReader.Text()
		matcher := regexp.MustCompile(`([0-9]+)-([0-9]+)\ ([a-z]):\ ([a-z]+)`)
		matchedComponents := matcher.FindStringSubmatch(input)[1:]

		min, _ := strconv.Atoi(matchedComponents[0])
		max, _ := strconv.Atoi(matchedComponents[1])
		c := []rune(matchedComponents[2])[0]
		pwd := matchedComponents[3]

		if isPolicySatisfied(int32(min)-1, int32(max)-1, c, pwd) {
			valid++
		}
	}

	return valid
}
