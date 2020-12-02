package p1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isPolicySatisfied(min int32, max int32, c int32, pwd string) bool {
	count := int32(0)
	for _, ch := range pwd {
		switch ch {
		case c:
			count++
			if count > max {
				return false
			}
		}
	}
	if count < min {
		return false
	}
	return true
}

func driver() int {
	filename := fmt.Sprintf("p1.input")
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

		if isPolicySatisfied(int32(min), int32(max), c, pwd) {
			valid++
		}
	}

	return valid
}
