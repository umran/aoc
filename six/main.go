package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// OrbitMap ...
type OrbitMap map[string]*Object

// Object ...
type Object struct {
	name       string
	parent     string
	satellites []string
}

func (o *Object) addSatellite(name string) {
	o.satellites = append(o.satellites, name)
}

func (o *Object) setParent(name string) {
	o.parent = name
}

func parseInput(fileName string) OrbitMap {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	objectMap := make(OrbitMap)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		pair := strings.Split(text, ")")

		primaryName := pair[0]
		secondaryName := pair[1]

		// upsert primary object and add secondary as a satellite
		primaryObject, _ := objectMap[primaryName]
		switch primaryObject {
		case nil:
			objectMap[primaryName] = &Object{name: primaryName}
		default:
			primaryObject.addSatellite(secondaryName)
		}

		// upsert secondary object
		secondaryObject, _ := objectMap[secondaryName]
		switch secondaryObject {
		case nil:
			objectMap[secondaryName] = &Object{name: secondaryName, parent: primaryName}
		default:
			secondaryObject.setParent(primaryName)
		}
	}

	return objectMap
}

func (om OrbitMap) computeOrder() (count int) {
	for _, object := range om {
		currentObject := object
	search:
		for {
			switch currentObject.parent {
			case "":
				break search
			default:
				count++
				currentObject = om[currentObject.parent]
			}
		}
	}

	return count
}

func main() {
	objectMap := parseInput("./input.txt")

	// this is the answer to part 1
	fmt.Println(objectMap.computeOrder())
}
