package main

import (
	"fmt"
	"time"
)

var (
	currentTurn = 1
	totalTurns  = 5
)

func main() {
	turn := make(chan bool)
	smallBreak := make(chan bool)
	longBreak := make(chan bool)
	done := make(chan bool)

	go pomodoroTurn(turn)
	go pomodoroService(turn, smallBreak, longBreak, done)

	<-done
}

func pomodoroTurn(turn chan bool) {
	tellBeginTurn()
	time.Sleep(time.Minute * 25)
	tellEndTurn()
	turn <- true
}

func tellBeginTurn() {
	fmt.Println("Pomodoro turn begins")
}

func tellEndTurn() {
	fmt.Println("Pomodoro turn ends")
}

func pomodoroBreak(chanBreak chan bool) {
	tellBeginSmallBreak()
	time.Sleep(5 * time.Minute)
	tellEndSmallBreak()
	chanBreak <- true
}

func pomodoroLongBreak(chanBreak chan bool) {
	tellBeginLongBreak()
	time.Sleep(30 * time.Minute)
	tellEndLongBreak()
	chanBreak <- true
}

func tellBeginSmallBreak() {
	fmt.Println("Have a small break")
}

func tellEndSmallBreak() {
	fmt.Println("Small break ended")
}

func tellBeginLongBreak() {
	fmt.Println("Have a long break")
}

func tellEndLongBreak() {
	fmt.Println("Long break ended")
}

func pomodoroService(turn, smallBreak, longBreak, done chan bool) {
	fmt.Println("pomodoro service started")
	fmt.Println("current turn:", currentTurn)
	for {
		select {
		case <-turn:
			if currentTurn >= totalTurns {
				go pomodoroLongBreak(longBreak)
				currentTurn = 1
			} else {
				currentTurn++
				go pomodoroBreak(smallBreak)
			}
		case <-smallBreak:
			go pomodoroTurn(turn)
		case <-longBreak:
			var needAnotherTurn string
			fmt.Printf("Another turn ?(Y/N)")
			fmt.Scanln(&needAnotherTurn)
			for needAnotherTurn != "Y" && needAnotherTurn != "N" {
				fmt.Scanln(&needAnotherTurn)
			}
			if needAnotherTurn == "Y" {
				go pomodoroTurn(turn)
			} else {
				done <- true
			}
		}
	}
}
