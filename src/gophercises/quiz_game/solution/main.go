package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

// Problem : Custom struct to handle questions and answers in quiz.
type Problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file is in the format of 'question,answer'")

	timeLimit := flag.Int("limit", 5, "time limit for the quiz in seconds.")

	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for idx, data := range lines {
		ret[idx] = Problem{data[0], data[1]}
	}
	return ret
}
