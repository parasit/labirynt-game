package labirynt

import "testing"

func TestPoint(t *testing.T) {
	p := Point2d{1, 1}
	result := p.ToJSON()
	// fmt.Println(result)
	if result != "{\"posX\":1,\"posY\":1}" {
		t.Error("Point2d wrong export")
	}
}

func TestRoom(t *testing.T) {
	r := Tile{0,
		"C0",
		Point2d{0, 0},
		[4]bool{true, true, true, true}}
	result := r.ToJSON()
	// fmt.Println(result)
	if result != `{"roomType":0,"typeId":"C0","position":{"posX":0,"posY":0},"exits":[true,true,true,true]}` {
		t.Error("Room wrong export to JSON", result)
	}
}

func TestPointOperations(t *testing.T) {
	p := Point2d{0, 0}
	if p.DistanceTo(Point2d{0, 1}) != 1 {
		t.Error("Distance error 1")
	}
	if p.DistanceTo(Point2d{0, 5}) != 5 {
		t.Error("Distance error 2")
	}
	if p.DistanceTo(Point2d{5, 5}) != 5 {
		t.Error("Distance  ", p.DistanceTo(Point2d{5, 5}))
	}
}
