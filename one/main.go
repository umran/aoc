package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calcFuel(mass float64) float64 {
	fuel := math.Floor(mass/3) - 2
	return fuel
}

func calcFuel2(mass float64) float64 {
	totalFuel := float64(0)
	fuel := calcFuel(mass)

	for fuel > 0 {
		totalFuel += fuel
		fuel = calcFuel(fuel)
	}

	return totalFuel
}

func floatFromString(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func main() {
	totalFuel := float64(0)
	totalFuel2 := float64(0)

	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass := floatFromString(scanner.Text())
		totalFuel += calcFuel(mass)
		totalFuel2 += calcFuel2(mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// this is the answer to part 1
	fmt.Println(int64(totalFuel))

	// this is the answer to part 2
	fmt.Println(int64(totalFuel2))
}
