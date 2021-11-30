package local_session

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

// calcSessionKey calculates unique key for a client from IP address and
// user agent string from its request header.
func calcSessionKey(r *http.Request) string {
	return strings.Split(r.RemoteAddr, ":")[0] + ":" + r.UserAgent()
}

// CalcBs32 generates random hex with 32bit unsigned integer and return it as string
func CalcBs32() string {
	return strconv.FormatInt(int64(rand.Uint32()), 16)
}

//CalcBs160 generates random with 160bit hex series
func CalcBs160() string {
	bs1 := CalcBs32() + CalcBs32()
	bs2 := CalcBs32()
	bs3 := CalcBs32()
	bs4 := CalcBs32()
	return bs1 + "-" + bs2 + "-" + bs3 + "-" + bs4
}
