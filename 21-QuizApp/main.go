package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func problemPuller(fileName string) ([]problem, error) {
	if fObj, err := os.Open(fileName); err == nil {
		defer fObj.Close()
		csvR := csv.NewReader(fObj)
		if qLines, err := csvR.ReadAll(); err == nil {
			return parser(qLines), nil
		} else {
			return nil, fmt.Errorf("Error while reading data from CSV: %s, error: %v", fileName, err)
		}
	} else {
		return nil, fmt.Errorf("Error opening file: %s, error: %v", fileName, err)
	}
}

func parser(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i, line := range lines {
		r[i] = problem{q: line[0], a: line[1]}
	}
	return r
}

func main() {
	//1. Take input of file for Quiz questions
	fileName := flag.String("f", "quiz.csv", "path of csv file")
	timer := flag.Int("t", 15, "max time for Quiz completion")
	flag.Parse()

	//3. Pull the problems from file.
	questions, err := problemPuller(*fileName)
	if err != nil {
		fmt.Printf("Error while loading Problems :%v", err)
		os.Exit(1)
	}

	correctAns := 0
	timeObj := time.NewTimer(time.Duration(*timer) * time.Second)

	ansC := make(chan string)
problemLoop:
	for i, problem := range questions {
		var answer string
		fmt.Printf("Question %d: %s = ", i, problem.q)
		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-timeObj.C:
			fmt.Printf("\nTIME UP....!!\n")
			break problemLoop

		case iAns := <-ansC:
			if iAns == problem.a {
				correctAns++
			}
			if i == len(questions)-1 {
				close(ansC)
			}
		}
	}
	fmt.Printf("Your quiz result is %d out of %d\n", correctAns, len(questions))
	fmt.Printf("Press Enter to Exit..\n")
	<-ansC
}
