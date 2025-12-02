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
	scanner.Scan()
	line := scanner.Text()

	ranges := strings.Split(line, ",")
	var sum int64

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		start, err1 := strconv.ParseInt(parts[0], 10, 64)
		end, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}

		for id := start; id <= end; id++ {
			if isInvalidID(id) {
				sum += id
			}
		}
	}

	fmt.Println(sum)
}

func isInvalidID(id int64) bool {
	idStr := strconv.FormatInt(id, 10)

	if len(idStr)%2 != 0 {
		return false
	}

	halfLen := len(idStr) / 2
	firstHalf := idStr[:halfLen]
	secondHalf := idStr[halfLen:]

	return firstHalf == secondHalf
}

func Part2() {
	file, err := os.Open("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	ranges := strings.Split(line, ",")
	var sum int64

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		start, err1 := strconv.ParseInt(parts[0], 10, 64)
		end, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}

		for id := start; id <= end; id++ {
			if isInvalidIDPart2(id) {
				sum += id
			}
		}
	}

	fmt.Println(sum)
}

func isInvalidIDPart2(id int64) bool {
	idStr := strconv.FormatInt(id, 10)
	strLen := len(idStr)

	for segmentLen := 1; segmentLen <= strLen/2; segmentLen++ {
		if strLen%segmentLen != 0 {
			continue
		}

		firstSegment := idStr[:segmentLen]
		numRepetitions := strLen / segmentLen

		allMatch := true
		for i := 1; i < numRepetitions; i++ {
			segment := idStr[i*segmentLen : (i+1)*segmentLen]
			if segment != firstSegment {
				allMatch = false
				break
			}
		}

		if allMatch && numRepetitions >= 2 {
			return true
		}
	}

	return false
}