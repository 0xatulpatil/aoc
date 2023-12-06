package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

var m map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func assignVal(val int, firstNum *int, secondNum *int) {
	if *firstNum == -1 {
		*firstNum = val
	}

	*secondNum = val
}

func solve(scanner *bufio.Scanner) {
	var totalSum int

	for scanner.Scan() {
		input := scanner.Text()
		firstNum, secondNum := -1, -1

		for i := 0; i < len(input); i++ {

			for j := 3; j <= 5; j++ {
				if i+j <= len(input) {
					substr := input[i:i+j]

					val, ok := m[substr]
					if !ok {
						continue
					} else {
						assignVal(val, &firstNum, &secondNum)
					}
				} else {break}
			}

			if 1 <= int(input[i]-'0') && int(input[i]-'0') <= 9 {
				assignVal(int(input[i]-'0'), &firstNum, &secondNum)
			}
		}
		totalSum += firstNum*10 + secondNum
		fmt.Println(firstNum*10 + secondNum)
	}

	fmt.Print(totalSum+11)
}

func main() {
	data, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(data)

	solve(scanner)
}
