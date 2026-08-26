package main

import "errors"

// Maps are similar to JavaScript objects, 
// Python dictionaries, and Ruby hashes. 
// Maps are a data structure that provides key->value mapping.

/*
Map values can be structs too:
type car struct {
  registration string
  model        string
}
cars := map[string]car{
  "ABC-123": {registration: "ABC-123", model: "Civic"},
}

The len() function works on a map, it returns the 
total number of key/value pairs.
*/

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers){
		lenError := errors.New("invalid sizes")
		return userMap, lenError
	}
	for i:=0;i<len(names);i++ {
		newUser := user{name:names[i], phoneNumber: phoneNumbers[i]}
		userMap[newUser.name] = newUser
	}
	return userMap, nil
	
}

type user struct {
	name        string
	phoneNumber int
}

