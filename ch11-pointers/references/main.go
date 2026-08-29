package main

/*
x := 5        // x holds the value 5
p := &x       // p holds the memory address of x
*/

import (
	"strings"
)
// *strings in the function signature means param is a an address, not direct value
// ie a pointer to a string
func removeProfanity(message *string) {
	// addyMsg := &message
	// *msg := addyMsg
	if strings.Contains(*message,"fubb") == true  {
		*message = strings.ReplaceAll(*message, "fubb", "****")

	} else if strings.Contains(*message,"shiz") == true {
		*message = strings.ReplaceAll(*message, "shiz", "****")

	} else if strings.Contains(*message,"witch") == true {
		*message = strings.ReplaceAll(*message, "witch", "*****")

	}

}

