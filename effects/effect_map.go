package effects

import "fmt"

// Map of effects
type Map struct {
	effects map[string]*Effect
}

// Init initialize map of effects
func (em *Map) Init() {
	em.effects = make(map[string]*Effect)
}

// Len returns length of map
func (em *Map) Len() int {
	return len(em.effects)
}

// AppendEffect adds effect to map
func (em *Map) AppendEffect(e *Effect) {
	// ToDo: check unique ID
	em.effects[e.ID] = e
}

// RemoveEffect form map
func (em *Map) RemoveEffect(key string) {
	delete(em.effects, key)
}

// MakeEffects use effect on target
func (em *Map) MakeEffects() {
	for key, e := range em.effects {
		fmt.Println(e)
		e.Tick()
		if e.TTL == 0 {
			em.RemoveEffect(key)
			fmt.Println("Removing effect ", key)
		}
	}
}
