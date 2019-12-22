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

	if s.start.x <= s2.start.x && s.end.x >= s2.end.x {
		return true
	}

	return false
}

func (s *Segment) containsAlongY(s2 *Segment) bool {
	if !s.alongY() {
		return false
	}

	if s.start.y <= s2.start.y && s.end.y >= s2.end.y {
		return true
	}

	return false
}

func (s *Segment) steps() int {
	if s.alongX() {
		return int(math.Abs(float64(s.start.x) - float64(s.end.x)))
	}

	return int(math.Abs(float64(s.start.y) - float64(s.end.y)))
}

func (c *Coordinate) move(m *Movement) *Coordinate {
	return &Coordinate{
		x: c.x + m.deltaX,
		y: c.y + m.deltaY,
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
		r, _ := utf8.DecodeRuneInString(value)

		var movement *Movement

		switch string(r) {
		case "U":
			deltaY, _ := strconv.ParseInt(strings.Split(value, "U")[1], 10, 64)
			movement = &Movement{
				deltaY: int(deltaY),
			}
		case "D":
			deltaY, _ := strconv.ParseInt(strings.Split(value, "D")[1], 10, 64)
			movement = &Movement{
				deltaY: -1 * int(deltaY),
			}
		case "R":
			deltaX, _ := strconv.ParseInt(strings.Split(value, "R")[1], 10, 64)
			movement = &Movement{
				deltaX: int(deltaX),
			}
		case "L":
			deltaX, _ := strconv.ParseInt(strings.Split(value, "L")[1], 10, 64)
			movement = &Movement{
				deltaX: -1 * int(deltaX),
			}
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
}
