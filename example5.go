// LogKey is a context key that can be used for
// getting a log.Logger from a request.
// Don't do this.
type LogKey struct{}

// AddLogger adds a log.Logger to a request.
// No really, Don't do this.
func AddLogger(next http.Handler, l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, LogKey{}, logger)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}
