## `HttpFilter`

[![CircleCI](https://circleci.com/gh/g1eng/httpfilter/tree/master.svg?style=svg)](https://circleci.com/gh/g1eng/httpfilter/tree/master)
[![codecov](https://codecov.io/gh/g1eng/httpfilter/branch/master/graph/badge.svg?token=EJZIHPRGNI)](https://codecov.io/gh/g1eng/httpfilter)

A set of conditional access control wrappers for golang-based web application, written in [httprouter](https://github.com/julienschmidt/httprouter) and http.HandlerFunc.

## Features

* variety of handler wrappers implemented in `AuthWrapper` helps you to protect resources for `http.HandlerFunc` and `httprouter.Handle`.
* several authentication providers and session management mechanisms are available in `auth` and `session` package for both of `http` and `httprouter` package.
* simple IP-based filtering using `jpillora/ipfilter`. (thanks @jpillora!)
* **AND** or **OR** synthetic wrapper with `httpfilter/syntesis` package, which enables you to apply two or more `AuthWrapper` for single route. (`RouterAuthWrapper` is also supported).
* Additional header management and built-in CORS support with `header` package (now only supported for `httprouter`)

## What is `AuthWrapper`?

```go
type AuthWrapper func (http.HandlerFunc, _ ...string) http.HandlerFunc
type RouterAuthWrapper func (httprouter.Handle, _ ...string) httprouter.Handle
```

AuthWrapper is the function type which receives `http.HandlerFunc` as its first argument, and returns `http.HandlerFunc`.
Its counter part for httprouter, RouterAuthWrapper is the function type which receives `httprouter.Handle` as its first argument, and returns `http.HandlerFunc`.
Both of them can receive additional string parameter for internal conditional evaluation.

If you feel `http.HandlerFunc` or `httprouter.Handle` friendly, maybe you should have been writing a number of wrappers for these handler functions. (also handler function itself). AuthWrappers (synthesys.AuthWrapper or rt_synthesis.AuthWrapper) are designed to be used as function generator which are acceptable for `http.HandleFunc` or `httprouter.Handle` like this:

```go
http.HandleFunc("/some/resource", someRouterAuthWrapper(yourHandler))
//or
router.GET("/api/path/somewhere", someRouterAuthWrapper(yourHandler))
```

Many of wrapper functions in this package are implemented in `AuthWrapper` type, including basic authentication, IP filtering, header validation, and so on.
You can apply single basic authentication for a simple but a little secured route with following snippet: 

```go
package example

import (
	"github.com/g1eng/httpfilter/auth/basic"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func yourHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//some code here
}

func herHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//any code here
}

func route() *httprouter.Router {
	router := httprouter.New()
	b := basic.NewBasicAuth("/path/to/htpasswd/or/credential/line")
	router.GET("/some/public/resource", yourHandler)
	router.GET("/some/secured/resource", b.RouterAuthenticate(herHandler))
	return router
}
```

## Function Synthesis for Access Control

Sometime, we need two or more request validators for a protected/hardened resource(s).
In such cases you can use `httpfilter/synthesis` to apply multiple authorization/validation mechanisms to users' traffic.

Two or more effects of `AuthWrapper` can be **synthesized** with `AuthAnd`, `AuthOR` or `AuthAll`.

You can write the synthesis of conditional checks with several `AuthWrappers` like this:

```go
package example

import (
	"github.com/g1eng/httpfilter/auth/basic"
	"github.com/g1eng/httpfilter/ipfilter"
	"github.com/g1eng/httpfilter/synthesis"
	"net/http"
)

func yourHandler(w http.ResponseWriter, r *http.Request) {
	//some code here
}

func herHandler(w http.ResponseWriter, r *http.Request) {
	//any code here
}

func Serve() {
	s := http.Server{
		Addr: "0.0.0.0:8080",
	}
	AND := synthesis.AuthAND

	defaultFilter := ipfilter.NewIPFilter(true, []string{"192.0.0.0/24"}).Authorize
	managedFilter := basic.NewBasicAuth("/path/to/htpasswd/or/credential/line").Authenticate

	//users must be authorized with two factor to access to protected resources for dualAuth
	dualAuth := AND(defaultFilter, managedFilter)
	
	http.HandleFunc("/o/ha", defaultFilter(yourHandler))
	http.HandleFunc("/o/con", dualAuth(herHandler))

	_ = s.ListenAndServe()
}

```

For `RouterAuthWrapper`, import `synthesis/rt_synthesis` package and use Auth* declared there:

```go
package example

import (
	"github.com/g1eng/httpfilter/auth/basic"
	"github.com/g1eng/httpfilter/ipfilter"
	"github.com/g1eng/httpfilter/synthesis/rt_synthesis"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func yourHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//some code here
}

func herHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//any code here
}

func Route() *httprouter.Router {
	router := httprouter.New()
	defaultFilter := ipfilter.NewIPFilter(true, []string{"192.0.0.0/24"}).RouterAuthorize
	managerFilter := basic.NewBasicAuth("/path/to/htpasswd/or/credential/line").Authenticate
	
	dualAuth := rt_synthesis.AuthAND(defaultFilter, managerFilter)
	
	router.GET("/some/corp/resource", defaultFilter(yourHandler))
	router.GET("/some/mgnt/resource", dualAuth(herHandler))
	return router
}
```

## Background

On my nearest experience, different project in different requirements with different stakeholders, share similar access control mechanisms that satisfy any of VIP's request within possible costs.
How do you think about such shared implementation can be reliable, full-featured, open and popular one?

This project is a proposal for generic access control wrapper mechanism for golang-based web applications.


## ToDo

* ensure basic auth has valid behavior
* hardening on local session storage
* redis token caching

## DOCUMENTATION

If you need any documentation enhancement, make issue or PR and post your request about desired additional topics!

## Contributing

You are welcomed to propose any type of commitment to this project! Contact from the issue page in open style and share your ideas about this package.

### LICENSE

Apache 2.0.
