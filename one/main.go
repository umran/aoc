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

func floatFromString(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func main() {
	totalFuel := float64(0)

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalFuel += calcFuel(floatFromString(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(int64(totalFuel))
}
