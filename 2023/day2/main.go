package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

const (
	RED_MAX   = 12
	GREEN_MAX = 13
	BLUE_MAX  = 14
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	game := regexp2.MustCompile(`(\d+):(( ?(\d+) (red|green|blue),?)+;?)+`, regexp2.RE2)

	total := 0
	var totalPower int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match, err := game.FindStringMatch(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		id, power := judgeGame(match)
		total += id
		totalPower += int64(power)
	}

	println("Total:", total)
	println("Total Power:", totalPower)
	fmt.Printf("(Took %dms)", time.Since(start)/1000)
}

func judgeGame(match *regexp2.Match) (int, int) {
	valid := true
	min_red := 0
	min_grn := 0
	min_blu := 0
	id, err := strconv.Atoi(match.GroupByNumber(1).String())
	if err != nil {
		log.Fatal(err)
	}
	for _, round := range match.GroupByNumber(2).Captures {
		items := strings.Split(strings.Trim(round.String(), " ;"), ", ")
		for _, x := range items {
			parts := strings.Split(x, " ")
			val, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			switch parts[1] {
			case "red":
				if val > RED_MAX {
					valid = false
				}
				if val > min_red {
					min_red = val
				}
			case "green":
				if val > GREEN_MAX {
					valid = false
				}
				if val > min_grn {
					min_grn = val
				}
			case "blue":
				if val > BLUE_MAX {
					valid = false
				}
				if val > min_blu {
					min_blu = val
				}
			default:
				log.Fatalf(`How did we get here? Expected [red|green|blue], got "%s"`, parts[1])
			}
		}
	}

	if valid {
		return id, min_red * min_grn * min_blu
	}

	return 0, min_red * min_grn * min_blu
}
