package filter

// SetDebug permissively turn on debug flag
func (f *HttpFilter) SetDebug() {
	f.debug = true
}

// SetOrigin sets origin hostname for CORS restriction
func (f *HttpFilter) SetOrigin(origin string) {
	f.origin = origin
}

// SetCustomHeader sets a custom header for the service endpoint
func (f *HttpFilter) SetCustomHeader(key string) {
	f.customHeader = key
}
