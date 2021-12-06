package basic

import (
	"github.com/g1eng/httpfilter/session/responder"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (b *Authenticator) Authenticate(handler http.HandlerFunc, _ ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := b.getAuthPayload(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		user, plainPass, err := SplitBasicCred(p)

		//Invalid authHeader
		if err != nil {
			log.Println("invalid authorization header")
			responder.Write400(w)
			return
		}
		log.Println("plain password: ", plainPass)
		for u, c := range b.userCredentials {
			if u == user {
				err = bcrypt.CompareHashAndPassword(c, []byte(plainPass))
				//cost, _ := bcrypt.Cost(c)
				//log.Println("cost for ", u, cost)
				log.Printf("user %s, err: %v", u, err)
				if err == nil {
					log.Println("authenticated")
					handler(w, r)
					return
				}
			}
		}
		responder.Write401(w)
		return
	}
}
