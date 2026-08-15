package main

func (e email) cost() int {
	cost := 0
	if !e.isSubscribed {
		cost = len(e.body)*5
	} else if e.isSubscribed {
		cost = len(e.body)*2
	}
	return cost
}

func (e email) format() string {
	if !e.isSubscribed {
		return "'" +e.body +"'" + " | " + "Not Subscribed"
	} 
	
	return "'" +e.body +"'" + " | " + "Subscribed"
	
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

