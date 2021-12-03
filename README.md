## `HttpFilter`

A set of conditional access control wrappers for golang web application based on [httprouter](https://github.com/julienschmidt/httprouter) and http.HandlerFunc.

### Status

[![CircleCI](https://circleci.com/gh/g1eng/httpfilter/tree/master.svg?style=svg)](https://circleci.com/gh/g1eng/httpfilter/tree/master)
[![codecov](https://codecov.io/gh/g1eng/httpfilter/branch/master/graph/badge.svg?token=EJZIHPRGNI)](https://codecov.io/gh/g1eng/httpfilter)

(Unstable)

## Features

* additional acceptable custom header with SetCustomHeader which appends a new atom in `Access-Control-Allow-Headers`
* most common hardening header and built-in CORS support, powered by httprouter.Router.Lookup
* **AND** or **OR** synthetic wrapper with `httpfilter/wrapper` package, which enables you to apply two or more `AuthWrapper` for single route.

#### ToDo

Other extending and developing fractions in this package are:

* basic authentication support <- doing
* redis backend

## DOCUMENTATION

WIP

### LICENSE

Apache 2.0
