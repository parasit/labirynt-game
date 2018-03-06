package effects

import "fmt"

// List basic list of effects
type List struct {
	effects []Effect
}

// Len returns number of elements
func (el *List) Len() int {
	return len(el.effects)
}

// AppendEffect appends effect to list
func (el *List) AppendEffect(e Effect) {
	el.effects = append(el.effects, e)
}

// RemoveEffect remove effect from list
func (el *List) RemoveEffect(effect Effect) {
	var e Effect
	var index int
	for index, e = range el.effects {
		if effect == e {
			el.effects[index] = el.effects[len(el.effects)-1]
			el.effects = el.effects[:len(el.effects)-1]
		}
	}
}

func (el *List) clearList() {
	for i := len(el.effects) - 1; i >= 0; i-- {
		if el.effects[i].TTL == 0 {
			el.effects = append(el.effects[:i], el.effects[i+1:]...)
		}
	}
}

// MakeEffects iterate over effects and do them
func (el *List) MakeEffects() {
	var e Effect
	var index int
	for index, e = range el.effects {
		fmt.Println(e.ToJSON())
		el.effects[index].Tick()
	}
	el.clearList()
}
