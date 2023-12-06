package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func filterWhiteSpace(st string) string {
	var res string

	for _, str := range st {
		if str != ' ' {
			res += string(str)
		}
	}
	return res
}

func main() {
	inputFile, _ := os.ReadFile("./Day6/input.txt")

	var totalProduct int = 1

	input := string(inputFile)
	inputString := strings.Split(input, "\n")

	timeString := strings.Split(inputString[0], ":")[1]
	distanceString := strings.Split(inputString[1], ":")[1]

	timeString = filterWhiteSpace(timeString)
	distanceString = filterWhiteSpace(distanceString)

	targetDistance, _ := strconv.Atoi(distanceString)
	startTime, _ := strconv.Atoi(timeString)

	for j := 1; j <= startTime/2; j++ {
		if j*(startTime-j) > targetDistance {
			pairs := (startTime/2 - j + 1) * 2

			if startTime%2 == 0 {
				pairs -= 1
			}
			totalProduct *= pairs
			break
		}
	}

	fmt.Println(totalProduct)
}
