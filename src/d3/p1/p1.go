package p1

import (
	"bufio"
	"fmt"
	"os"
)

func checkIfTree(c rune) int {
	if c == '#' {
		return 1
	}
	return 0
}

func walkThroughGrid(grid [][]rune, x int, y int, xLimit int, yLimit int) int {
	treesFound := 0
	if y >= yLimit {
		return treesFound
	}

	return treesFound + checkIfTree(grid[y][x]) + walkThroughGrid(grid, (x+3)%xLimit, y+1, xLimit, yLimit)
}

func driver() int {
	filename := fmt.Sprintf("p1.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	var grid [][]rune

	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		inputLine := bufReader.Text()
		grid = append(grid, []rune(inputLine))
	}

	if grid == nil {
		return 0
	}

	xLimit := len(grid[0])
	yLimit := len(grid)

	return walkThroughGrid(grid, 0, 0, xLimit, yLimit)
}
