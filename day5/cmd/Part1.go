package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	End   int64
}

func Part1() {
	data, err := os.ReadFile("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	blankLineIndex := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			blankLineIndex = i
			break
		}
	}

	if blankLineIndex == -1 {
		log.Fatal("No blank line found in input")
	}

	ranges := []Range{}
	for i := 0; i < blankLineIndex; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		start, err1 := strconv.ParseInt(parts[0], 10, 64)
		end, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}

	ingredientIDs := []int64{}
	for i := blankLineIndex + 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			continue
		}
		ingredientIDs = append(ingredientIDs, id)
	}

	freshCount := 0
	for _, id := range ingredientIDs {
		isFresh := false
		for _, r := range ranges {
			if id >= r.Start && id <= r.End {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshCount++
		}
	}

	fmt.Println(freshCount)
}

func Part2() {
	data, err := os.ReadFile("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	blankLineIndex := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			blankLineIndex = i
			break
		}
	}

	if blankLineIndex == -1 {
		log.Fatal("No blank line found in input")
	}

	ranges := []Range{}
	for i := 0; i < blankLineIndex; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		start, err1 := strconv.ParseInt(parts[0], 10, 64)
		end, err2 := strconv.ParseInt(parts[1], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	if len(ranges) == 0 {
		fmt.Println(0)
		return
	}

	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		current := ranges[i]

		if current.Start <= last.End+1 {
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			merged = append(merged, current)
		}
	}

	totalCount := int64(0)
	for _, r := range merged {
		totalCount += r.End - r.Start + 1
	}

	fmt.Println(totalCount)
}