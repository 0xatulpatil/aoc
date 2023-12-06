package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(scanner *bufio.Scanner) {
	for scanner.Scan() {

		input := scanner.Text()
		gameNo := input[5:6]

		fmt.Println(gameNo)
	}
}

func main() {
	data, err := os.Open("./Day2/input.txt")
	if err != nil {
		panic(err)
	}

	inputData := bufio.NewScanner(data)

	solve(inputData)
	fmt.Print("hello")
}
