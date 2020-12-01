package p2

import (
	"bufio"
	"fmt"
	"github.com/rounakdatta/aoc-2020/src/util"
	"os"
	"strconv"
)

func driver() int {
	filename := fmt.Sprintf("p2.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	var inputs []int
	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		input, ie := strconv.Atoi(bufReader.Text())
		if ie != nil {
			panic(util.INPUT_ERROR)
		}
		inputs = append(inputs, input)
	}

	for _, num1 := range inputs {
		for _, num2 := range inputs {
			for _, num3 := range inputs {
				if num1+num2+num3 == 2020 {
					return num1 * num2 * num3
				}
			}
		}
	}

	return -1
}
