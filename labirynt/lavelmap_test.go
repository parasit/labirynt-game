package labirynt

import (
	"fmt"
	"testing"
)

func TestLevelRooms(t *testing.T) {
	var l labiryntMapType

	l.InitLabiryntLevel("tilesTemplates.json")

	if l.isTileFree(Point2d{0, 0}) {
		t.Error("Central point is free")
	}
	if l.isTileFree(Point2d{0, 1}) == false {
		t.Error("Room should be free")
	}

	if l.GetPossibileExitsAt(Point2d{0, 0}) != "****" {
		t.Error("Desired exits not matches")
	}

	r := Tile{0,
		"C0",
		Point2d{0, 1},
		[4]bool{true, true, true, true}}

	if !l.exitsMatches(r) {
		t.Error("Exits not maches")
	}
	l.AddRoomToLabirynt(r)
	fmt.Println(l.GetPossibileExitsAt(Point2d{0, 1}))

	if l.isTileFree(Point2d{0, 1}) {
		t.Error("Test point is free")
	}

	r2 := Tile{0,
		"C0",
		Point2d{0, -1},
		[4]bool{true, true, false, true}}

	if l.exitsMatches(r2) {
		t.Error("Exits matches !!!")
	}
	r3 := Tile{0,
		"C0",
		Point2d{1, 0},
		[4]bool{false, false, false, true}}
	l.AddRoomToLabirynt(r3)
	fmt.Println(l.GetPossibileExitsAt(Point2d{1, 1}))
}

func TestDesiredExitsTest(t *testing.T) {
	var l labiryntMapType
	r := Tile{0,
		"C0",
		Point2d{1, 0},
		[4]bool{true, true, true, true}}
	l.AddRoomToLabirynt(r)
	l.GetPossibileExitsAt(Point2d{1, 0})
	r2 := Tile{0,
		"C0",
		Point2d{0, 1},
		[4]bool{true, true, true, true}}
	l.GetPossibileExitsAt(Point2d{0, 1})
	l.AddRoomToLabirynt(r2)
	l.GetPossibileExitsAt(Point2d{1, 1})
	r3 := Tile{0,
		"C0",
		Point2d{1, 1},
		[4]bool{true, false, false, true}}
	l.AddRoomToLabirynt(r3)
	if l.GetPossibileExitsAt(Point2d{1, 1}) != "+**+" {
		t.Error("Desired exits error")
	}
}

func TestLevelJSON(t *testing.T) {
	var ll labiryntMapType
	ll.InitLabiryntLevel("tilesTemplates.json")
	res := ll.ToJSON()
	if len(res) <= 0 {
		t.Error("ToJSON error", res)
	}
}
