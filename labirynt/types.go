package labirynt

import (
	"encoding/json"
	"fmt"
	"math"
)

//North global
const North = 0

//East global
const East = 1

//South global
const South = 2

//West global
const West = 3

// Directions global directions
var Directions = []string{"North", "East", "South", "West"}

// Point2d simple point definition
type Point2d struct {
	PosX int `json:"posX"`
	PosY int `json:"posY"`
}

// Init Inicjalizuje pokój
func Init() {

}

// ToJSON exports point2d to JSON
func (point Point2d) ToJSON() string {
	// fmt.Println(point)
	b, err := json.Marshal(Point2d(point))
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	return string(b)
}

//GetOpositeDirection returns oposite direction eg. N-S, E-W
func GetOpositeDirection(direction int) int {
	return (direction + 2) % 4
}

//DistanceTo return distance beetween points
func (point Point2d) DistanceTo(point2 Point2d) float64 {
	return math.Sqrt(math.Pow((float64)(point.PosX-point2.PosX), 2.0) + math.Pow((float64)(point.PosY-point2.PosY), 2.0))
}
