package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	// *timeLimit is an int
	// time.Second is a time.Duration
	// So *timeLimit needs to be converted into time.Duration
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// waiting for a message from that channel. code will block until a message from channel is received.
	// <-timer.C

	correct := 0
	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		default:
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Wrong answer!")
			}
		}
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
