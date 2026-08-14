package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	sender := mToSend.sender
	recipient := mToSend.recipient
	if sender.name == "" {
		return false
	} else if sender.number == 0 {
		return false
	}

	if recipient.name == "" {
		return false
	} else if recipient.number == 0 {
		return false
	}
	return true
}

