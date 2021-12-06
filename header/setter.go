package header

// SetDebug permissively turn on debug flag
func (f *Filter) SetDebug() {
	f.debug = true
}

// SetOrigin sets origin hostname for CORS restriction
func (f *Filter) SetOrigin(origin string) {
	f.origin = origin
}

// SetCustomHeader sets a custom header for the service endpoint
func (f *Filter) SetCustomHeader(key string) {
	f.customHeader = key
}
