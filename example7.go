
func Next(step int, data interface{}) (cache interface{}, done bool) {
	// State machine will always advance "step"
	switch step {
	case 0:
		// Do stuff
		// Return a "random" integer ID:
		return 4, false
	case 1:
		// We know the ID was an int, so type assert:
		id := data.(int)
		// Do more stuff
		return nil, true
	}
	panic("the state machine is broken!")
}
