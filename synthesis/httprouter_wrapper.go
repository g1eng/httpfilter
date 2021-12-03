package synthesis

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
)

// AuthAND realizes function synthesis (AND pipeline) for two filter.AuthWrapper.
// Two authentication wrappers (auth1 and auth2) are applied for inbound traffic
// in the order, and finally returns httprouter.Handle for route for valid traffics.
//
func AuthAND(auth1 AuthWrapper, auth2 AuthWrapper) AuthWrapper {
	return func(handle httprouter.Handle) httprouter.Handle {
		return auth2(auth1(handle))
	}
}

// AuthOR realizes function parallel (OR pipeline) for two filter.AuthWrapper.
//
// The first evaluated handler `auth1(handle)` (httprouter.Handle) must return
// status code 200 at valid result. But, if the first function is expected to fail,
// it must not return StatusOk for the response and subsequent handler `auth2(handle)`
// should be evaluated immediately.
//
func AuthOR(auth1 AuthWrapper, auth2 AuthWrapper) AuthWrapper {
	return func(handle httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			testWriter := httptest.NewRecorder()
			auth1(handle)(testWriter, r, ps)
			if testWriter.Code == 200 {
				// Because the first function may cause mutable change for any backend state,
				// auth1 must not run again.
				_, _ = w.Write([]byte(testWriter.Body.String()))
			} else {
				auth2(handle)(w, r, ps)
			}
		}
	}
}

// AuthAll realizes two or more AuthWrapper pipeline with AND condition. (experimental)
//
func AuthAll(auth ...AuthWrapper) AuthWrapper {
	//if len(auth) == 1, then simply apply the AuthWrapper
	return func(hx httprouter.Handle) httprouter.Handle {
		for _, wrapper := range auth {
			hx = wrapper(hx)
		}
		return hx
	}
}
