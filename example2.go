package context

// WithValue returns a copy of parent in which the value
// associated with key is val.
func WithValue(
	parent Context, key, val interface{},
) Context

type Context interface {
	// Value returns the value associated with this
	// context for key, or nil if no value is
	// associated with key.
	Value(key interface{}) interface{}
}
