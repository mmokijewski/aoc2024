package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func checkMachine(machine []float32) int {

	x1 := machine[0]
	y1 := machine[1]
	x2 := machine[2]
	y2 := machine[3]
	sum1 := machine[4]
	sum2 := machine[5]

	x := (sum1*y2 - sum2*x2) / (x1*y2 - y1*x2)
	y := (sum1 - (x * x1)) / x2

	if x == float32(int(x)) && y == float32(int(y)) {
		return int((3 * x) + y)
	} else {
		return 0
	}
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	numRegex := regexp.MustCompile(`\d+`)
	part1Sum := 0

	i := 0
	var machine []float32
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if i == 3 {
			part1Sum += checkMachine(machine)
			machine = nil
			i = 0
		} else {
			i++
		}
		line := scanner.Text()
		numbers := numRegex.FindAllString(line, 2)
		for _, stringNum := range numbers {
			num, _ := strconv.Atoi(stringNum)
			machine = append(machine, float32(num))
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
