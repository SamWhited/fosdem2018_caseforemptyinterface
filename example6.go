// Mechanism represents an auth mechanism
// (eg. plain, scram, or oauth2).
type Mechanism struct {
	Next func(data interface{}) (cache interface{})
}

// Negotiator is a state machine that handles
// requests and responses in the auth flow.
type Negotiator struct{
	cache interface{}
}

// Step advances the state machine.
func (c *Negotiator) Step(challenge []byte) (resp []byte)
