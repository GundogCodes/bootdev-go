package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	mtc, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	mts, er := sendSMS(msgToSpouse)
	if er != nil {
		return 0, er
	}
	return  mtc + mts, er
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

