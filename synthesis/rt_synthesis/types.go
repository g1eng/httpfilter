package rt_synthesis

import (
	"github.com/julienschmidt/httprouter"
)

//AuthWrapper is wrapper function for httprouter.Handle.
//It receives httprouter.Handle and returns httprouter.Handle,
//that means users can write their own rule for session handling
//inside this type of function.
type AuthWrapper func(httprouter.Handle, ...string) httprouter.Handle
