package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1() {
	file, err := os.Open("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(grid) == 0 {
		fmt.Println(0)
		return
	}

	rows := len(grid)
	cols := len(grid[0])

	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	accessibleCount := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != '@' {
				continue
			}

			adjacentCount := 0
			for _, dir := range dirs {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
					if grid[ni][nj] == '@' {
						adjacentCount++
					}
				}
			}

			if adjacentCount < 4 {
				accessibleCount++
			}
		}
	}

	fmt.Println(accessibleCount)
}

func Part2() {
	file, err := os.Open("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(grid) == 0 {
		fmt.Println(0)
		return
	}

	rows := len(grid)
	cols := len(grid[0])

	gridBytes := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		gridBytes[i] = []byte(grid[i])
	}

	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalRemoved := 0

	for {
		toRemove := make([][]int, 0)

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if gridBytes[i][j] != '@' {
					continue
				}

				adjacentCount := 0
				for _, dir := range dirs {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
						if gridBytes[ni][nj] == '@' {
							adjacentCount++
						}
					}
				}

				if adjacentCount < 4 {
					toRemove = append(toRemove, []int{i, j})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			gridBytes[pos[0]][pos[1]] = '.'
			totalRemoved++
		}
	}

	fmt.Println(totalRemoved)
}