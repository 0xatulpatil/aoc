package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func filterWhiteSpaces(arr []string) []string {
	var res []string

	for _, str := range(arr) {
		if strings.TrimSpace(str) != "" {
			res = append(res, str)
		}
	}
	return res
}

func main() {
	inputFile, _ := os.ReadFile("./Day6/input.txt")

	var totalProduct int = 1

	input := string(inputFile)
	inputString := strings.Split(input, "\n")

	timeVector := strings.Split(strings.TrimSpace(strings.Split(inputString[0], ":")[1]), " ")
	distanceVector := strings.Split(strings.TrimSpace(strings.Split(inputString[1], ":")[1]), " ")

	timeVector = filterWhiteSpaces(timeVector);
	distanceVector = filterWhiteSpaces(distanceVector);


	for i := 0; i < len(timeVector); i++ {
		targetDistance, _ := strconv.Atoi(distanceVector[i])
		startTime, _ := strconv.Atoi(timeVector[i])
		

		for j := 1; j <= startTime/2; j++ {
			if j*(startTime-j) > targetDistance {
				pairs := (startTime/2 - j + 1)*2

				if startTime%2 == 0 {pairs -= 1}
				totalProduct *= pairs
				break
			}
		}
	}

	// fmt.Println(timeVector)
	// fmt.Println(distanceVector)
	fmt.Println(totalProduct)
}
