package p2

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

func walkThroughGrid(grid [][]rune, x int, y int, xLimit int, yLimit int, xMovement int, yMovement int) int {
	treesFound := 0
	if y >= yLimit {
		return treesFound
	}

	return treesFound + checkIfTree(grid[y][x]) + walkThroughGrid(grid, (x+xMovement)%xLimit, y+yMovement, xLimit, yLimit, xMovement, yMovement)
}

func driver() int {
	filename := fmt.Sprintf("p2.input")
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

	paths := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	productOfTrees := 1
	for _, path := range paths {
		productOfTrees *= walkThroughGrid(grid, 0, 0, xLimit, yLimit, path[0], path[1])
	}
	return productOfTrees
}
