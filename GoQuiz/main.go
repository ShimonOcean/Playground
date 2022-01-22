//Refamiliarizing with opening files, csv, channels

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

func main() {
	// Handy to give notice to users who do not have access to code
	csvFilename := flag.String("csv", "problems.csv", "csv file, format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit for quiz in seconds")
	flag.Parse()

	// Open csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided CSV file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	<-timer.C

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Congrats, you scored %d out of %d!\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("Congrats, you scored %d out of %d!\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Open and Read csv file
// f, err := os.Open("data.csv")
// if err != nil {
// 	log.Fatal(err)
// }

// defer f.Close()

// csvReader := csv.NewReader(f)
// lines, err := csvReader.Read()
// lines contains all lines of csv file
