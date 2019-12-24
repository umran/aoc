package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Coordinate ...
type Coordinate struct {
	x int
	y int
}

// Movement ...
type Movement struct {
	deltaX int
	deltaY int
}

// Segment ...
type Segment struct {
	start *Coordinate
	end   *Coordinate
}

// Wire ...
type Wire struct {
	segments []*Segment
}

func (c *Coordinate) move(m *Movement) *Coordinate {
	return &Coordinate{
		x: c.x + m.deltaX,
		y: c.y + m.deltaY,
	}
}

func (c *Coordinate) equals(c2 *Coordinate) bool {
	return c.x == c2.x && c.y == c2.y
}

func (s *Segment) maxX() int {
	return int(math.Max(float64(s.start.x), float64(s.end.x)))
}

func (s *Segment) minX() int {
	return int(math.Min(float64(s.start.x), float64(s.end.x)))
}

func (s *Segment) maxY() int {
	return int(math.Max(float64(s.start.y), float64(s.end.y)))
}

func (s *Segment) minY() int {
	return int(math.Min(float64(s.start.y), float64(s.end.y)))
}

func (s *Segment) alongX() bool {
	return s.start.x != s.end.x
}

func (s *Segment) alongY() bool {
	return s.start.y != s.end.y
}

func (s *Segment) containsAlongX(s2 *Segment) bool {
	if !s.alongX() {
		return false
	}

	if s.minX() <= s2.minX() && s.maxX() >= s2.maxX() {
		return true
	}

	return false
}

func (s *Segment) containsAlongY(s2 *Segment) bool {
	if !s.alongY() {
		return false
	}

	if s.minY() <= s2.minY() && s.maxY() >= s2.maxY() {
		return true
	}

	return false
}

func (s *Segment) containsCoordinate(c *Coordinate) bool {
	switch {
	case s.alongX():
		return s.minX() <= c.x && s.maxX() >= c.x && s.start.y == c.y
	case s.alongY():
		return s.minY() <= c.y && s.maxY() >= c.y && s.start.x == c.x
	default:
		return s.start.equals(c)
	}
}

func (s *Segment) steps() int {
	switch {
	case s.alongX():
		return int(math.Abs(float64(s.start.x) - float64(s.end.x)))
	case s.alongY():
		return int(math.Abs(float64(s.start.y) - float64(s.end.y)))
	default:
		return 0
	}
}

func parseInput(filename string) *Wire {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := string(bytes)
	textArray := strings.Split(text, ",")

	origin := &Coordinate{0, 0}
	wire := &Wire{
		segments: make([]*Segment, len(textArray)),
	}

	for i, value := range textArray {
		movement := new(Movement)

		r, _ := utf8.DecodeRuneInString(value)
		magnitude, _ := strconv.ParseInt(strings.Split(value, string(r))[1], 10, 64)

		switch string(r) {
		case "U":
			movement.deltaY = int(magnitude)
		case "D":
			movement.deltaY = -1 * int(magnitude)
		case "R":
			movement.deltaX = int(magnitude)
		case "L":
			movement.deltaX = -1 * int(magnitude)
		}

		switch i {
		case 0:
			wire.segments[i] = &Segment{
				start: origin,
				end:   origin.move(movement),
			}
		default:
			wire.segments[i] = &Segment{
				start: wire.segments[i-1].end,
				end:   wire.segments[i-1].end.move(movement),
			}
		}
	}

	return wire
}

func manhattanDistance(a, b *Coordinate) int {
	return int(math.Abs(float64(a.x)-float64(b.x)) + math.Abs(float64(a.y)-float64(b.y)))
}

func findIntersections(a, b *Wire) []*Coordinate {
	intersections := make([]*Coordinate, 0)

	for i := 0; i < len(a.segments); i++ {
		for j := 0; j < len(b.segments); j++ {
			segA := a.segments[i]
			segB := b.segments[j]

			switch {
			case segA.alongX() && segB.alongY():
				if segA.containsAlongX(segB) && segB.containsAlongY(segA) {
					intersections = append(intersections, &Coordinate{
						x: segB.start.x,
						y: segA.start.y,
					})
				}
			case segA.alongY() && segB.alongX():
				if segA.containsAlongY(segB) && segB.containsAlongX(segA) {
					intersections = append(intersections, &Coordinate{
						x: segA.start.x,
						y: segB.start.y,
					})
				}
			}
		}
	}

	return intersections
}

func stepsToReachIntersection(wire *Wire, intersection *Coordinate) int {
	totalSteps := 0
	for _, segment := range wire.segments {
		totalSteps += segment.steps()

		if segment.containsCoordinate(intersection) {
			// we may have overstepped so adjust total:
			switch {
			case segment.alongX():
				totalSteps -= int(math.Abs(float64(segment.end.x) - float64(intersection.x)))
			case segment.alongY():
				totalSteps -= int(math.Abs(float64(segment.end.y) - float64(intersection.y)))
			}
			break
		}
	}

	return totalSteps
}

func main() {
	wireA := parseInput("./inputA.txt")
	wireB := parseInput("./inputB.txt")

	intersections := findIntersections(wireA, wireB)

	distances := make([]int, len(intersections))
	for i, intersection := range intersections {
		//fmt.Printf("we have an intersection at: (x: %d, y: %d) \n", intersection.x, intersection.y)
		distances[i] = manhattanDistance(intersection, &Coordinate{0, 0})
		//fmt.Printf("the manhattan distance between this intersection and the central point is: %d \n", distances[i])
	}

	closestDistance := distances[0]
	for _, distance := range distances {
		if distance < closestDistance {
			closestDistance = distance
		}

		fmt.Printf("candidate distance: %d \n", distance)
	}

	// this is the answer to part 1
	fmt.Printf("closest distance: %d \n", closestDistance)

	allCombinedSteps := make([]int, 0)
	for _, intersection := range intersections {
		combinedSteps := stepsToReachIntersection(wireA, intersection) + stepsToReachIntersection(wireB, intersection)
		allCombinedSteps = append(allCombinedSteps, combinedSteps)
	}

	fewestCombinedSteps := allCombinedSteps[0]
	for _, cs := range allCombinedSteps {
		if cs < fewestCombinedSteps {
			fewestCombinedSteps = cs
		}

		fmt.Printf("candidate combined steps: %d \n", cs)
	}

	// this is the answer to part 2
	fmt.Printf("fewest combined steps: %d \n", fewestCombinedSteps)
}
