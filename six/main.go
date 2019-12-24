package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// OrbitMap ...
type OrbitMap struct {
	objectMap    map[string]*Object
	searchedList map[string]bool
}

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

func (o *Object) transferables() []string {
	transferables := make([]string, len(o.satellites), len(o.satellites)+1)
	copy(transferables, o.satellites)

	if o.parent != "" {
		transferables = append(transferables, o.parent)
	}

	return transferables
}

func (o *Object) transferableTo(destination string) bool {
	if o.parent == destination {
		return true
	}

	for _, sat := range o.satellites {
		if destination == sat {
			return true
		}
	}

	return false
}

func parseInput(fileName string) *OrbitMap {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	objectMap := make(map[string]*Object)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		pair := strings.Split(text, ")")

		primaryName := pair[0]
		secondaryName := pair[1]

		// upsert primary object and add secondary as a satellite
		primaryObject := objectMap[primaryName]
		switch primaryObject {
		case nil:
			objectMap[primaryName] = &Object{name: primaryName, satellites: []string{secondaryName}}
		default:
			primaryObject.addSatellite(secondaryName)
		}

		// upsert secondary object and set primary as its parent
		secondaryObject, _ := objectMap[secondaryName]
		switch secondaryObject {
		case nil:
			objectMap[secondaryName] = &Object{name: secondaryName, parent: primaryName}
		default:
			secondaryObject.setParent(primaryName)
		}
	}

	return &OrbitMap{
		objectMap: objectMap,
	}
}

func (om *OrbitMap) _initializeSearchedList() {
	om.searchedList = make(map[string]bool)
}

func (om *OrbitMap) countDirectAndIndirectOrbits() (count int) {
	for _, object := range om.objectMap {
		currentObject := object
	search:
		for {
			switch currentObject.parent {
			case "":
				break search
			default:
				count++
				currentObject = om.objectMap[currentObject.parent]
			}
		}
	}

	return count
}

// this is a stateful recursive function that should always be called after _initializeSearchedList()
func (om *OrbitMap) _enumerateNodesBetween(origin, destination string) (bool, []string) {
	var transfers []string

	object := om.objectMap[origin]
	if object.transferableTo(destination) {
		return true, transfers
	}

	// flag the current object as being searched
	om.searchedList[origin] = true

	for _, transferable := range object.transferables() {
		// if the transferable has already been searched, continue
		if _, searched := om.searchedList[transferable]; searched == true {
			continue
		}

		// otherwise, recursively search its immediate neighbours
		isTransferable, subTransfers := om._enumerateNodesBetween(transferable, destination)
		if isTransferable {
			transfers = append(transfers, transferable)
			transfers = append(transfers, subTransfers...)
			return true, transfers
		}
	}

	return false, transfers
}

func (om *OrbitMap) enumerateNodesBetween(origin, destination string) []string {
	om._initializeSearchedList()
	_, nodes := om._enumerateNodesBetween(origin, destination)
	return nodes
}

func main() {
	om := parseInput("./input.txt")

	// this is the answer to part 1
	fmt.Println(om.countDirectAndIndirectOrbits())

	nodes := om.enumerateNodesBetween("YOU", "SAN")
	fmt.Println(nodes)

	// this is the answer to part 2
	fmt.Println(len(nodes) - 1)
}
