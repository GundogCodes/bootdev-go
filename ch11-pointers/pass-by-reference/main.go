package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?

func analyzeMessage(analytics *Analytics, message Message) {
	isSuccess := message.Success
	if isSuccess {
		
		(*analytics).MessagesTotal++
		(*analytics).MessagesSucceeded++
	} else {
		(*analytics).MessagesTotal++
		(*analytics).MessagesFailed++
	}
}
