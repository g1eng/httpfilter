package synthesis

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//AuthWrapper is wrapper function for httprouter.Handle.
//It recieves httprouter.Handle and returns httprouter.Handle,
//that means users can write their own rule for session handling
//inside this type of function.
type AuthWrapper func(httprouter.Handle) httprouter.Handle

//PlainAuthWrapper is wrapper function for http.HandlerFunc,
//for whom not using httprouter. It provides simple wrapper interface
//for any http.HandleFunc functions and define how they are dealt.
type PlainAuthWrapper func(handler http.HandlerFunc) http.HandlerFunc
