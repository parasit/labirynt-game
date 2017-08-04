package labirynt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/parasit/rpgforge"
)

// Tile default room/corridor tile
type Tile struct {
	RoomType int     `json:"roomType"`
	TypeID   string  `json:"typeId"`
	Position Point2d `json:"position"`
	Exits    [4]bool `json:"exits"`
}

var tileTypes = []string{"corridor", "room"}

// TileTemplate default tile template struct
type TileTemplate struct {
	Type  int     `json:"type"`
	ID    string  `json:"id"`
	Exits [4]bool `json:"exits"`
}

// TilesTemplatesType base
type TilesTemplatesType struct {
	Tiles []TileTemplate
}

//TilesTemplates global templates list
var TilesTemplates TilesTemplatesType

func exitsMatches(exits string, testedExits [4]bool) bool {
	result := true
	for i := 0; i < 4; i++ {
		if exits[i] == '+' && testedExits[i] == false {
			result = false
		}
		if exits[i] == '-' && testedExits[i] == true {
			result = false
		}
	}
	return result
}

func filteredTiles(exits string, roomType int) []TileTemplate {
	var result []TileTemplate
	for _, tRoom := range TilesTemplates.Tiles {
		if tRoom.Type == roomType && exitsMatches(exits, tRoom.Exits) {
			result = append(result, tRoom)
		}
	}
	return result
}

// ToJSON exports Room
func (room Tile) ToJSON() string {
	// fmt.Println(room)
	b, err := json.Marshal(Tile(room))
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	return string(b)
}

// TileFromTemplate ugly merge template and position to Tile/Room
func TileFromTemplate(template TileTemplate, position Point2d) Tile {
	var tTile Tile
	tTile.Position = position
	tTile.RoomType = template.Type
	tTile.TypeID = template.ID
	tTile.Exits = template.Exits
	return tTile
}

// InitRandomTile generates random tile
func InitRandomTile(position Point2d) Tile {
	var r = Tile{
		rpgforge.ThrowDices("1d2").Sum() - 1,
		"Random",
		position,
		[4]bool{true, true, true, true}}

	//var requredExits string
	for i := 0; i < 4; i++ {

	}
	var tNr = rand.Intn(len(TilesTemplates.Tiles))
	r.Exits = TilesTemplates.Tiles[tNr].Exits
	r.TypeID = TilesTemplates.Tiles[tNr].ID

	return r
}

// LoadTilesTemplates load tiles templates from json file
func LoadTilesTemplates(filepath string) TilesTemplatesType {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c TilesTemplatesType
	json.Unmarshal(raw, &c)
	return c
}
