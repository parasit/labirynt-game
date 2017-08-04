package main

import (
	"fmt"

	"github.com/parasit/labirynt-game/labirynt"
)

func main() {
	labirynt.CurrentLabirynt.InitLabiryntLevel("labirynt/tilesTemplates.json")
	//labirynt.CurrentLabirynt.AddRoomToLabirynt(labirynt.InitRandomTile(labirynt.Point2d{0, 1}))
	labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{0, 1})
	labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{1, 0})
	labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{1, 1})
	fmt.Println(labirynt.CurrentLabirynt.ToJSON())
}
