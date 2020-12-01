package p1

import (
	"bufio"
	"fmt"
	"github.com/rounakdatta/aoc-2020/src/util"
	"os"
	"strconv"
)

func reader(state map[int]bool) int {
	filename := fmt.Sprintf("p1.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		input, ie := strconv.Atoi(bufReader.Text())
		if ie != nil {
			panic(util.INPUT_ERROR)
		}

		pairingInput := 2020 - input
		if _, ok := state[pairingInput]; ok {
			return input * pairingInput
		} else {
			state[input] = true
		}
	}

	return -1
}

func driver() int {
	state := map[int]bool{}
	return reader(state)
}
