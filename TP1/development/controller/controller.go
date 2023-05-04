package controller

import "TPs-DDS/TP1/development/persistence"

func CheckEmailAndPassword(email, password string) bool {
	checkEmail := persistence.CheckEmail(email)

	if !checkEmail {
		return false
	}

	return persistence.CheckPassword(email, password)
}
