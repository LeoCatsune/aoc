package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left := []int{}
	right := []int{}

	re := regexp.MustCompile(`\s+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals := re.Split(scanner.Text(), 2)
		left = insertSorted(left, vals[0])
		right = insertSorted(right, vals[1])
	}

	// sanity check
	if len(left) != len(right) {
		log.Fatalf("Mismatched array lengths: left=%d, right=%d", len(left), len(right))
	}

	distance := 0
	similarity := 0
	for i, l := range left {
		if left[i] > right[i] {
			distance += left[i] - right[i]
		} else {
			distance += right[i] - left[i]
		}

		for _, r := range right {
			if l == r {
				similarity += l
			}
		}
	}

	println("Entries:", len(left))
	println("Distance:", distance)
	println("Similarity:", similarity)
	fmt.Printf("(Took %dms)", time.Since(start)/1000)
}

func insertSorted(slice []int, value string) []int {
	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	index, _ := slices.BinarySearch(slice, v)
	slice = slices.Insert(slice, index, v)
	return slice
}
