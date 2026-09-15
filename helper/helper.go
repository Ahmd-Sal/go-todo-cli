package helper

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func ReadLimitedString(scanner *bufio.Scanner, maxChars int) string {
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		if utf8.RuneCountInString(text) <= maxChars {
			return text
		}

		fmt.Printf("Input excceds the %d charachers limit. Please try again: \n", maxChars)
	}
	return ""
}

func ReadTaskID(scanner *bufio.Scanner, listLength int) int {
	for scanner.Scan() {
		number := strings.TrimSpace(scanner.Text())
		convertedString, err := strconv.Atoi(number) // Convert the string into an int first

		if err != nil {
			fmt.Println("Error: the input is not a number. Please try again")
			continue
		}

		if convertedString >= 1 && convertedString <= listLength {
			return convertedString
		}
		fmt.Println("Task ID does not exist. Please try again.")
	}
	return -1
}

func ReturnStatus(num int) string {
	if num == 1 {
		return "Done"
	}
	return "Pending"

}
