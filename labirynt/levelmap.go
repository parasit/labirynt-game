package labirynt

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type labiryntMapType struct {
	Tiles []Tile `json:"levelTiles"`
}

// CurrentLabirynt aktualna mapa labiryntu
var CurrentLabirynt labiryntMapType

// InitLabiryntLevel create empty labirynt type
func (l *labiryntMapType) InitLabiryntLevel(path string) {
	rand.Seed(time.Now().UnixNano())
	// fmt.Println("Inicjalizacja labirytntu")
	TilesTemplates = LoadTilesTemplates(path)
	var r = Tile{
		1,
		"R1",
		Point2d{0, 0},
		[4]bool{true, true, true, true}}
	l.AddRoomToLabirynt(r)
	//l.addRandomRooms()
}

func (l labiryntMapType) ToJSON() string {
	b, err := json.Marshal(labiryntMapType(l))
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	return string(b)
}

func (l *labiryntMapType) GetPossibileExitsAt(position Point2d) string {
	var desiredExits string
	desiredExits = ""
	for i := 0; i < 4; i++ {
		room := l.getRoomAtPosition(position.getNextPosition(i))
		if room != nil {
			if room.Exits[GetOpositeDirection(i)] {
				desiredExits += "+"
			} else {
				desiredExits += "-"
			}
		} else {
			desiredExits += "*"
		}
	}
	fmt.Println(desiredExits, position)
	return desiredExits
}

// AddRandomRoom check conditions and add random room to position
func (l *labiryntMapType) AddRandomRoom(position Point2d) {
	dExits := l.GetPossibileExitsAt(position)
	fTiles := filteredTiles(dExits, 0)
	var tNr = rand.Intn(len(fTiles))
	fmt.Println(fTiles[tNr])
	room := TileFromTemplate(fTiles[tNr], position)
	l.AddRoomToLabirynt(room)
}

func (l *labiryntMapType) IsTileLegal(room Tile) bool {
	if l.isTileFree(room.Position) && l.exitsMatches(room) {
		return true
	}
	return false
}

// AddRoomToLabirynt add new room
func (l *labiryntMapType) AddRoomToLabirynt(room Tile) bool {
	if l.IsTileLegal(room) {
		l.Tiles = append(l.Tiles, room)
		return true
	}
	fmt.Println("Illegal room !", room)
	return false
}

func (l labiryntMapType) getRoomAtPosition(Position Point2d) *Tile {
	for _, element := range l.Tiles {
		if element.Position == Position {
			return &element
		}
	}
	return nil
}

func (p Point2d) getNextPosition(direction int) Point2d {
	var tPos Point2d
	switch direction {
	case North:
		tPos = Point2d{p.PosX, p.PosY - 1}
	case East:
		tPos = Point2d{p.PosX + 1, p.PosY}
	case South:
		tPos = Point2d{p.PosX, p.PosY + 1}
	case West:
		tPos = Point2d{p.PosX - 1, p.PosY}
	}
	return tPos
}

func (l *labiryntMapType) isTileFree(pos Point2d) bool {
	if l.getRoomAtPosition(pos) == nil {
		return true
	}
	return false
}

func (l labiryntMapType) testExit(r Tile, direction int) bool {
	tPos := r.Position.getNextPosition(direction)
	if l.isTileFree(tPos) {
		// fmt.Println("Room not exist", tPos)
		return true
	}
	r2 := l.getRoomAtPosition(tPos)
	if r2.Exits[GetOpositeDirection(direction)] == r.Exits[direction] {
		return true
	}
	return false
}

func (l labiryntMapType) exitsMatches(newRoom Tile) bool {
	for i := 0; i < 4; i++ {
		if !l.testExit(newRoom, i) {
			return false
		}
	}
	return true
}
