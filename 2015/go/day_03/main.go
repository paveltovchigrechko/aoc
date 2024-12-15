package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input_03.txt"
	north     = "^"
	south     = "v"
	east      = ">"
	west      = "<"
	sep       = ""
)

type Coordinates struct {
	Latitude  int
	Longitude int
}

func (c Coordinates) Print() {
	fmt.Printf("Latitude: %d, longitude: %d\n", c.Latitude, c.Longitude)
}

func (c *Coordinates) Move(direction string) {
	switch direction {
	case north:
		c.Longitude += 1
	case south:
		c.Longitude -= 1
	case east:
		c.Latitude += 1
	case west:
		c.Latitude -= 1
	default:
		log.Printf("Unknown direction: %s\n", direction)
	}
}

func trackVisitedHomes(route string) int {
	visitedLocations := make(map[Coordinates]bool)
	currentPos := Coordinates{
		Latitude:  0,
		Longitude: 0,
	}

	visitedHomes := 1
	visitedLocations[currentPos] = true

	for _, arrow := range route {
		direction := string(arrow)
		currentPos.Move(direction)
		if !visitedLocations[currentPos] {
			visitedLocations[currentPos] = true
			visitedHomes += 1
		}
	}
	return visitedHomes
}

func trackVisitedHomesForSantaAndRoboSanta(santaRoute, roboSantaRoute string) int {
	visitedLocations := make(map[Coordinates]bool)
	currentSantaPos := Coordinates{
		Latitude:  0,
		Longitude: 0,
	}

	currentRoboSantaPos := Coordinates{
		Latitude:  0,
		Longitude: 0,
	}

	visitedHomes := 1
	visitedLocations[currentSantaPos] = true

	for i := 0; i < len(santaRoute); i++ {
		santaDirection := string(santaRoute[i])
		currentSantaPos.Move(santaDirection)
		if !visitedLocations[currentSantaPos] {
			visitedLocations[currentSantaPos] = true
			visitedHomes += 1
		}

		roboSantaDirection := string(roboSantaRoute[i])
		currentRoboSantaPos.Move(roboSantaDirection)
		if !visitedLocations[currentRoboSantaPos] {
			visitedLocations[currentRoboSantaPos] = true
			visitedHomes += 1
		}
	}

	return visitedHomes
}

func makeSantaAndRoboSantaRoutes(route string) (string, string) {
	santaArrows := make([]string, 0, len(route)/2)
	roboSantaArrows := make([]string, 0, len(route)/2)
	for i, arrow := range route {
		if i%2 == 0 {
			santaArrows = append(santaArrows, string(arrow))
		} else {
			roboSantaArrows = append(roboSantaArrows, string(arrow))
		}
	}

	santaRoute := strings.Join(santaArrows, sep)
	roboSantaRoute := strings.Join(roboSantaArrows, sep)
	return santaRoute, roboSantaRoute
}

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	fmt.Printf("%d houses received at least one present\n", trackVisitedHomes(input))

	santa, roboSanta := makeSantaAndRoboSantaRoutes(input)
	fmt.Printf("%d houses received at least one present the next year when Santa and Robo-Santa delivered them\n", trackVisitedHomesForSantaAndRoboSanta(santa, roboSanta))
}
