package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFile, err := os.Open("/home/ven/problems.csv")
	if err != nil {
		panic(err)
	}
	csvData := csv.NewReader(bufio.NewReader(csvFile))
	qna, _ := csvData.ReadAll()
	noOfQuestions := len(qna)
	counter := 0

	fmt.Println("Press Enter to start the quiz")
	fmt.Println("-----------------------------")
	fmt.Scanln()

	timer1 := time.NewTimer(30 * time.Second)

	answerByClient := make(chan bool)
	done := make(chan bool)

	go func(answerByClient chan bool, done chan bool) {
		for idx, data := range qna {
			fmt.Printf("Question %d): %s\nAnswer: ", idx+1, data[0])
			var answered string
			fmt.Scanln(&answered)
			if answered == data[1] {
				counter++
				answerByClient <- true
			}
		}
		// if all questions are taken before time, exit and show results.
		done <- true
	}(answerByClient, done)

	go func(done chan bool) {
		for {
			select {
			case <-answerByClient:
			case <-timer1.C:
				// fmt.Println("timed out")
				done <- true
			}
		}
	}(done)
	<-done
	fmt.Println("you scored:", counter, "out of", noOfQuestions)
}
