package labirynt

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	LoadTilesTemplates("tilesTemplates.json")
	var r = InitRandomTile(Point2d{1, 1})
	fmt.Println(r.ToJSON())
}

func TestExitChecks(t *testing.T) {
	var e [4]bool
	e = [4]bool{true, true, false, false}
	if exitsMatches("++++", e) {
		t.Error("Not Matches 1")
	}
	if !exitsMatches("**--", e) {
		t.Error("Not Matches 2")
	}
	if !exitsMatches("****", e) {
		t.Error("Not Matches 3")
	}
}
