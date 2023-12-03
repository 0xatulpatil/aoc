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

func main() {
	data, err := os.Open("./input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(data)
	var totalSum int

	for scanner.Scan() {
		input := scanner.Text()
		firstNum, secondNum := -1, -1

		for _, char := range input {

			asciiChar := char - '0'

			if int(asciiChar) <= 9 {
				if firstNum == -1 {
					firstNum = int(asciiChar)
				}

				secondNum = int(asciiChar)
			}
		}

		firstNum = firstNum*10 + secondNum
		totalSum += firstNum

		// fmt.Print(firstNum+11, "\n")
	}

	fmt.Print(totalSum + 11)
}
