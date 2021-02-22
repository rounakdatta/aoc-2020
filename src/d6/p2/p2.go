package p2

import (
	"bufio"
	"fmt"
	"os"
)

const ElementCount = 26

func clear(elements *[ElementCount]int) {
	for i, _ := range *elements {
		(*elements)[i] = 0
	}
}

func count(elements *[ElementCount]int, groupCount int) int {
	count := 0
	for i, _ := range *elements {
		if (*elements)[i] == groupCount {
			count += 1
		}
	}
	return count
}

func set(response []rune, elements *[ElementCount]int) {
	for _, c := range response {
		(*elements)[c - 97] += 1
	}
}

func driver() int {
	filename := fmt.Sprintf("p2.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	var checker [26]int
	groupCount := 0
	countSum := 0

	clear(&checker)
	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		singleResponse := bufReader.Text()
		if singleResponse == "" {
			countSum += count(&checker, groupCount)
			clear(&checker)
			groupCount = 0
		} else {
			set([]rune(singleResponse), &checker)
			groupCount += 1
		}
	}
	countSum += count(&checker, groupCount)
	return countSum
}
