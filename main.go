package main

import (
	"fmt"

	"github.com/parasit/labirynt-game/tools"
	"github.com/parasit/rpgforge"

	"github.com/parasit/labirynt-game/effects"
	"github.com/parasit/labirynt-game/labirynt"
)

func main() {
	labirynt.CurrentLabirynt.InitLabiryntLevel("labirynt/tilesTemplates.json")
	//labirynt.CurrentLabirynt.AddRoomToLabirynt(labirynt.InitRandomTile(labirynt.Point2d{0, 1}))
	// labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{0, 1})
	// labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{1, 0})
	// labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{1, 1})
	fmt.Println(labirynt.CurrentLabirynt.ToJSON())

	// var efl effects.List
	var m effects.Map

	m.Init()

	ee := effects.Effect{tools.RandStringRunes(32), "Some effect", "damage", rpgforge.ThrowDices("2d6").Sum(), "3d6"}
	ee2 := effects.Effect{tools.RandStringRunes(32), "Some effect", "healing", rpgforge.ThrowDices("2d6").Sum(), "3d6"}

	m.AppendEffect(&ee)
	m.AppendEffect(&ee2)

	for m.Len() > 0 {
		m.MakeEffects()
	}

}
