package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/parasit/labirynt-game/tools"
	"github.com/parasit/rpgforge"

	"github.com/parasit/labirynt-game/effects"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	// labirynt.CurrentLabirynt.InitLabiryntLevel("labirynt/tilesTemplates.json")
	//labirynt.CurrentLabirynt.AddRoomToLabirynt(labirynt.InitRandomTile(labirynt.Point2d{0, 1}))
	// labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{0, 1})
	// labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{1, 0})
	// labirynt.CurrentLabirynt.AddRandomRoom(labirynt.Point2d{1, 1})
	// fmt.Println(labirynt.CurrentLabirynt.ToJSON())

	// var efl effects.List
	var m effects.Map

	m.Init()

	ee := effects.Effect{tools.RandStringRunes(32), "Some effect", "damage", rpgforge.ThrowDices("2d6").Sum(), "3d6"}
	ee2 := effects.Effect{tools.RandStringRunes(32), "Some effect", "healing", rpgforge.ThrowDices("2d6").Sum(), "10"}

	m.AppendEffect(&ee)
	m.AppendEffect(&ee2)

	for m.Len() > 0 {
		m.MakeEffects()
	}

	tools.Throw("d3+3d13!+4*2d4")
}
