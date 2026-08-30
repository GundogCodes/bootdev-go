package main

import (
	"fmt"
)

func test(text string) {
	fmt.Println(text)
}

func main() {
	test("starting Textio server")
	test("stopping Textio server")
}

/*
Go programs are organized into packages. 
A package is a directory of Go code that's 
all compiled together. Functions, types, 
variables, and constants defined in one source
file are visible to all other source files 
within the same package (directory).

A repository contains one or more modules. 
A module is a collection of Go packages 
that are released together.

Directory Structure
You will have many git repositories on your machine (typically one per project).
Each repository is typically a single module.
Each module contains one or more packages
Each package consists of one or more Go source files in a single directory.


You have to explicitly import them, unlike Swift where files in the same module 
can see each other automatically.

example:
import "github.com/GundogCodes/bootdev-go/ch04-functions"

And the function has to be exported (capitalized) to be usable from another package:

// ch04-functions/main.go
func Concat(s1, s2 string) string {  // capital C = exported
    return s1 + s2
}
*/