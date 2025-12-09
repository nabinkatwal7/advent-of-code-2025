package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	data, err := os.ReadFile("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	var nonEmptyLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	if len(nonEmptyLines) < 2 {
		log.Fatal("Not enough lines in input")
	}

	operationLine := nonEmptyLines[len(nonEmptyLines)-1]
	numberLines := nonEmptyLines[:len(nonEmptyLines)-1]

	operationTokens := strings.Fields(operationLine)

	var numberTokens [][]string
	maxCols := 0
	for _, line := range numberLines {
		tokens := strings.Fields(line)
		numberTokens = append(numberTokens, tokens)
		if len(tokens) > maxCols {
			maxCols = len(tokens)
		}
	}

	if len(operationTokens) < maxCols {
		log.Fatal("Operations line doesn't have enough tokens")
	}

	grandTotal := 0
	for col := 0; col < maxCols; col++ {
		var numbers []int
		for _, tokens := range numberTokens {
			if col < len(tokens) && tokens[col] != "" {
				num, err := strconv.Atoi(tokens[col])
				if err == nil {
					numbers = append(numbers, num)
				}
			}
		}

		if len(numbers) == 0 {
			continue
		}

		op := strings.TrimSpace(operationTokens[col])

		var result int
		if op == "*" {
			result = 1
			for _, num := range numbers {
				result *= num
			}
		} else if op == "+" {
			result = 0
			for _, num := range numbers {
				result += num
			}
		} else {
			continue
		}

		grandTotal += result
	}

	fmt.Println(grandTotal)
}

func Part2() {
	data, err := os.ReadFile("cmd/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	var nonEmptyLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	if len(nonEmptyLines) < 2 {
		log.Fatal("Not enough lines in input")
	}

	operationLine := nonEmptyLines[len(nonEmptyLines)-1]
	numberLines := nonEmptyLines[:len(nonEmptyLines)-1]

	maxWidth := 0
	for _, line := range nonEmptyLines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	for i := range numberLines {
		for len(numberLines[i]) < maxWidth {
			numberLines[i] += " "
		}
	}
	for len(operationLine) < maxWidth {
		operationLine += " "
	}

	grandTotal := 0
	col := maxWidth - 1

	for col >= 0 {
		isSeparator := true
		for _, line := range numberLines {
			if col < len(line) && line[col] != ' ' {
				isSeparator = false
				break
			}
		}

		if isSeparator {
			col--
			continue
		}

		problemStart := col
		problemEnd := col
		for problemEnd >= 0 {
			isSep := true
			for _, line := range numberLines {
				if problemEnd < len(line) && line[problemEnd] != ' ' {
					isSep = false
					break
				}
			}
			if isSep {
				break
			}
			problemEnd--
		}
		problemEnd++ // problemEnd is now the leftmost column of the problem

		// Get the operation from the bottom row
		var op byte = ' '
		for c := problemStart; c >= problemEnd; c-- {
			if c < len(operationLine) {
				char := operationLine[c]
				if char == '+' || char == '*' {
					op = char
					break
				}
			}
		}

		if op == ' ' {
			col = problemEnd - 1
			continue
		}


		var numbers []int

		for c := problemStart; c >= problemEnd; c-- {
			hasDigits := false
			columnChars := []byte{}
			for _, line := range numberLines {
				if c < len(line) {
					char := line[c]
					if char >= '0' && char <= '9' {
						hasDigits = true
						columnChars = append(columnChars, char)
					} else {
						columnChars = append(columnChars, ' ')
					}
				} else {
					columnChars = append(columnChars, ' ')
				}
			}

			if hasDigits {
				numStr := ""
				for _, char := range columnChars {
					if char >= '0' && char <= '9' {
						numStr += string(char)
					}
				}

				numStr = strings.TrimSpace(numStr)
				if numStr != "" {
					num, err := strconv.Atoi(numStr)
					if err == nil {
						numbers = append(numbers, num)
					}
				}
			}
		}

		if len(numbers) > 0 {
			var result int
			if op == '*' {
				result = 1
				for _, num := range numbers {
					result *= num
				}
			} else if op == '+' {
				result = 0
				for _, num := range numbers {
					result += num
				}
			}
			grandTotal += result
		}

		col = problemEnd - 1
	}

	fmt.Println(grandTotal)
}