package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		score := 50
		return score
	}
	return dm.priorityLevel
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
	
}
func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch imp := n.(type){
	case directMessage:
		return imp.senderUsername, imp.importance()
	case groupMessage:
		return imp.groupName, imp.importance()
	case systemAlert:
		return imp.alertCode, imp.importance()
	default:
		return "", 0
	}
}

