package effects

import (
	"encoding/json"
	"fmt"

	"github.com/parasit/rpgforge"
)

var effectKinds = [...]string{"heal", "damage", "poison"}

// Effect basic effect struct
type Effect struct {
	ID    string `json:"effect_id"`
	Name  string `json:"effect_name"`
	Kind  string `json:"effect_kind"`
	TTL   int    `json:"effect_ttl"`
	Value string `json:"effect_value"`
}

// ToJSON export Effect struct to json
func (e Effect) ToJSON() string {
	b, err := json.Marshal(Effect(e))
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	return string(b)
}

// Tick make effect and decrease TTL
func (e *Effect) Tick() {
	fmt.Printf("Effect %s %v\n", e.Kind, rpgforge.ThrowDices(e.Value).Sum())
	e.TTL += -1
}

// String returns json string
func (e *Effect) String() string {
	return e.ToJSON()
}
