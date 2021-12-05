package basic

import (
	"github.com/g1eng/httpfilter/session/responder"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

//RouterAuthenticate handles authentication process via Basic authentication.
//If any unauthorized access, it sends WWW-RouterAuthenticate header for the client
//and terminate the session.
func (b *Authenticator) RouterAuthenticate(handle httprouter.Handle, _ ...string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		p, err := b.getAuthPayload(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		user, plainPass, err := SplitBasicCred(p)

		//Invalid authHeader
		if err != nil {
			responder.Write400(w)
			return
		}
		for u, c := range b.userCredentials {
			if u == user {
				err = bcrypt.CompareHashAndPassword(c, []byte(plainPass))
				log.Printf("user %s, err: %v", u, err)
				if err == nil {
					handle(w, r, ps)
					return
				}
			}
		}

		responder.Write401(w)
		return
	}
}
