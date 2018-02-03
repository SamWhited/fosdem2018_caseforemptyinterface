
func Next(step int, data interface{}) interface{} {
	// State machine will always advance "step"
	switch step {
	case 0:
		// Do stuff
		// Return a "random" integer ID:
		return 4
	case 1:
		// We know it's an int!
		id := data.(int)
		// Do more stuff
		return nil
	}
	panic("the state machine is broken!")
}
