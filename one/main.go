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
	fmt.Printf("the fuel required for a mass of %f is %f \n", mass, fuel)

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

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		totalFuel += calcFuel(floatFromString(s))
		totalFuel2 += calcFuel2(floatFromString(s))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// this is the value for puzzle 1
	fmt.Println(int64(totalFuel))

	// this is the value for puzzle 2
	fmt.Println(int64(totalFuel2))
}
