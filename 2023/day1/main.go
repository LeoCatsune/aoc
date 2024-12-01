package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/dlclark/regexp2"
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// because fuck you, I'm not using a loop.
	num := regexp2.MustCompile(`(?=(\d|one|two|three|four|five|six|seven|eight|nine))`, regexp2.RE2)
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := regexp2FindAllString(num, scanner.Text())
		first := nums[0]
		last := nums[len(nums)-1]

		fv := slices.Index(words, first)
		if fv == -1 {
			fv, err = strconv.Atoi(first)
			if err != nil {
				log.Fatal(err)
			}
		}

		lv := slices.Index(words, last)
		if lv == -1 {
			lv, err = strconv.Atoi(last)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("%s (%s) (%s) => %d%d\n", scanner.Text(), first, last, fv, lv)

		sum += (fv * 10) + lv
	}

	println("Sum:", sum)
	fmt.Printf("(Took %dms)", time.Since(start)/1000)
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.GroupByNumber(1).String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
