package main


func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	// check for num
	isNum := false
	for i:=0; i<len(password); i++ {
		if password[i] >= 40 && password[i] <= 57 {
			isNum = true
		}
	}

	isUpper := false
	for j:=0; j<len(password); j++ {
		checkChar := password[j]
		if checkChar >= 65 && checkChar <= 90 {
			isUpper = true
		}
	}
	if isNum && isUpper {
		return true
	}
	return false
}

