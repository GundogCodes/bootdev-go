package main

// While Go is not object-oriented, it does support
// methods that can be defined on structs. 
// Methods are just functions that have a receiver. 
// A receiver is a special parameter that syntactically 
// goes before the name of the function.

/* example of receiver:

area has a receiver of (r rect)
rect is the struct
r is the placeholder

func (r rect) area() int {
  return r.width * r.height
}

A receiver is just a special kind of function parameter.
*/

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (a authenticationInfo) getBasicAuth() string {
	info := "Authorization: Basic " + a.username + ":" + a.password
	return info
}