package synthesis

import (
	"net/http"
)

//AuthWrapper is wrapper function for http.HandlerFunc,
//for whom not using httprouter. It receives single http.HandlerFunc
//function and returns http.HandlerFunc. Users can define how traffics
//are evaluated inside the wrapper and what is returned.
type AuthWrapper func(http.HandlerFunc, ...string) http.HandlerFunc
