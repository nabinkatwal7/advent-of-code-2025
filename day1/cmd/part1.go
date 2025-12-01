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

func Part2() {
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
		if amount < 0 {
			log.Fatalf("negative rotation %q not supported", line)
		}

		switch dir {
		case 'L', 'R':
			hitsThisMove := 0
			if amount > 0 {
				var dist int
				if position == 0 {
					dist = 100
				} else {
					if dir == 'R' {
						dist = 100 - position
					} else { // 'L'
						dist = position
					}
				}

				if amount >= dist {
					hitsThisMove = 1 + (amount-dist)/100
				}
			}
			hitsZero += hitsThisMove

			if dir == 'R' {
				position = (position + amount) % 100
			} else {
				position = (position - amount) % 100
				if position < 0 {
					position += 100
				}
			}

		default:
			log.Fatalf("unknown direction %q in line %q", string(dir), line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(hitsZero)
}