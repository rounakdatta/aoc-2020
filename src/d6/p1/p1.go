package p1

import (
	"bufio"
	"fmt"
	"os"
)

const ElementCount = 26

func clear(elements *[ElementCount]bool) {
	for i, _ := range *elements {
		(*elements)[i] = false
	}
}

func count(elements *[ElementCount]bool) int {
	count := 0
	for i, _ := range *elements {
		if (*elements)[i] {
			count += 1
		}
	}
	return count
}

func set(response []rune, elements *[ElementCount]bool) {
	for _, c := range response {
		(*elements)[c - 97] = true
	}
}

func driver() int {
	filename := fmt.Sprintf("p1.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	var checker [26]bool
	countSum := 0

	clear(&checker)
	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		singleResponse := bufReader.Text()
		if singleResponse == "" {
			countSum += count(&checker)
			clear(&checker)
		} else {
			set([]rune(singleResponse), &checker)
		}
	}
	countSum += count(&checker)
	return countSum
}
