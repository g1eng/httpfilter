package synthesis

import (
	"net/http"
	"net/http/httptest"
)

// AuthAND realizes function synthesis (AND pipeline) for two filter.AuthWrapper.
// Two authentication wrappers (auth1 and auth2) are applied for inbound traffic
// in the order, and finally returns httprouter.Handle for route for valid traffics.
// (experimental)
//
func AuthAND(auth1 AuthWrapper, auth2 AuthWrapper) AuthWrapper {
	return func(handle http.HandlerFunc, _ ...string) http.HandlerFunc {
		return auth2(auth1(handle))
	}
}

// AuthOR (experimental)
func AuthOR(auth1 AuthWrapper, auth2 AuthWrapper) AuthWrapper {
	return func(handle http.HandlerFunc, _ ...string) http.HandlerFunc {
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

// AuthAll realizes two or more AuthWrapper pipeline with AND condition. (experimental)
func AuthAll(auth ...AuthWrapper) AuthWrapper {
	//if len(auth) == 1, then simply apply the RouterAuthWrapper
	return func(hx http.HandlerFunc, _ ...string) http.HandlerFunc {
		for _, wrapper := range auth {
			hx = wrapper(hx)
		}
		return hx
	}
}
