package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	var err = errors.New("unsupported plan")
	if plan == "pro" {
		return messages[0:3], nil
	} else if plan == "free" {
		return messages[0:2], nil
	} else {
		var nilSlice []string
		return nilSlice, err
	}

}

