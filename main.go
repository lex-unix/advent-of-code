package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mapping = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getMatch(line string, check func(string, string) bool) string {
	for k, v := range mapping {
		if check(line, k) || check(line, v) {
			return v
		}
	}
	return ""
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	contents := string(file)
	lines := strings.Split(contents, "\n")
	acc := int64(0)

	for _, line := range lines {

		first := ""
		last := ""

		for i := 0; i < len(line); i++ {
			substring := string(line[i:])
			if first = getMatch(substring, strings.HasPrefix); first != "" {
				break
			}
		}

		for i := len(line); i >= 0; i-- {
			substring := string(line[:i])
			if last = getMatch(substring, strings.HasSuffix); last != "" {
				break
			}
		}

		num, err := strconv.ParseInt(first+last, 10, 64)

		if err == nil {
			acc += num
		}
	}
	fmt.Println(acc)
}

