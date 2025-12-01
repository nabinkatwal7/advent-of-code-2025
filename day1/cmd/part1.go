package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
file, err := os.Open("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 50
	hitsZero := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		dir := line[0]
		numStr := line[1:]
		amount, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("invalid rotation %q: %v", line, err)
		}

		amount = amount % 100

		switch dir {
		case 'L':
			position = (position - amount) % 100
			if position < 0 {
				position += 100
			}
		case 'R':
			position = (position + amount) % 100
		default:
			log.Fatalf("unknown direction %q in line %q", string(dir), line)
		}

		if position == 0 {
			hitsZero++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(hitsZero)
}