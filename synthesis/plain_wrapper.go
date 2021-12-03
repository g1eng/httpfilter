package synthesis

import (
	"net/http"
	"net/http/httptest"
)

// PlainAuthAND realizes function synthesis (AND pipeline) for two filter.PlainAuthWrapper.
// Two authentication wrappers (auth1 and auth2) are applied for inbound traffic
// in the order, and finally returns httprouter.Handle for route for valid traffics.
// (experimental)
//
func PlainAuthAND(auth1 PlainAuthWrapper, auth2 PlainAuthWrapper) PlainAuthWrapper {
	return func(handle http.HandlerFunc) http.HandlerFunc {
		return auth2(auth1(handle))
	}
}

// PlainAuthOR (experimental)
func PlainAuthOR(auth1 PlainAuthWrapper, auth2 PlainAuthWrapper) PlainAuthWrapper {
	return func(handle http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			testWriter := httptest.NewRecorder()
			auth1(handle)(testWriter, r)
			if testWriter.Code == 200 {
				_, _ = w.Write([]byte(testWriter.Body.String()))
			} else {
				auth2(handle)(w, r)
			}
		}
	}
}

// PlainAuthAll realizes two or more PlainAuthWrapper pipeline with AND condition. (experimental)
func PlainAuthAll(auth ...PlainAuthWrapper) PlainAuthWrapper {
	//if len(auth) == 1, then simply apply the AuthWrapper
	return func(hx http.HandlerFunc) http.HandlerFunc {
		for _, wrapper := range auth {
			hx = wrapper(hx)
		}
		return hx
	}
}
