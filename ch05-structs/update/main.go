package main

type User struct {
	Membership
	Name string
}

type Membership struct {
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	// var user = User{}
	// user.Name = name
	// user.Type = membershipType

	// or

	user := User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
		},
	}

	if user.Type == "premium" {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}
	return user
}

