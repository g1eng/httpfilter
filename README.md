## `HttpFilter`

A set of conditional access control wrappers for golang web application based on [httprouter](https://github.com/julienschmidt/httprouter).

### Status

Unstable

## Features

* additional acceptable custom header with SetCustomHeader which appends a new atom in `Access-Control-Allow-Headers`
* most common hardening header and built-in CORS support, powered by httprouter.Router.Lookup
* **AND** or **OR** synthetic wrapper with `httpfilter/wrapper` package, which enables you to apply two or more `AuthWrapper` for single route.

#### Experimental ToDo

Other extending and developing fractions in this package are:

* builtin IP filtering and basic authentication support
* ~~passing any authentication backend for an authorization~~ (that's enough to give httprouter.Handle for user-defined AuthWrapper. `AuthAnd` makes multi-factor authentication easier.)
* redis backend

## DOCUMENTATION

WIP

### LICENSE

Apache 2.0