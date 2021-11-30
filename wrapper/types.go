package wrapper

import "github.com/julienschmidt/httprouter"

type AuthWrapper func(httprouter.Handle) httprouter.Handle
