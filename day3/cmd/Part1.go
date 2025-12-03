package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Part1() {
	file, err := os.Open("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		maxJoltage := 0

		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				joltage := int(line[i]-'0')*10 + int(line[j]-'0')
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}

		total += maxJoltage
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func Part2() {
	file, err := os.Open("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int64

	for scanner.Scan() {
		line := scanner.Text()


		stack := make([]byte, 0, 12)
		toRemove := len(line) - 12

		for i := 0; i < len(line); i++ {

			for len(stack) > 0 && toRemove > 0 && line[i] > stack[len(stack)-1] {
				remaining := len(line) - i
				needed := 12 - (len(stack) - 1)
				if remaining >= needed {
					stack = stack[:len(stack)-1]
					toRemove--
				} else {
					break
				}
			}
			stack = append(stack, line[i])
		}

		if len(stack) > 12 {
			stack = stack[:12]
		}

		resultStr := string(stack)
		joltage, err := strconv.ParseInt(resultStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		total += joltage
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}