package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	first := regexp.MustCompile(`^.*?(\d)`)
	last := regexp.MustCompile(`.*(\d).*?$`)

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, err := strconv.Atoi(first.FindStringSubmatch(scanner.Text())[1])
		if err != nil {
			log.Fatal(err)
		}

		l, err := strconv.Atoi(last.FindStringSubmatch(scanner.Text())[1])
		if err != nil {
			log.Fatal(err)
		}

		sum += (f * 10) + l
	}

	println("Sum:", sum)
	fmt.Printf("(Took %dms)", time.Since(start)/1000)
}
