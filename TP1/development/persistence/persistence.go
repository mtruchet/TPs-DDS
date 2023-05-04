package persistence

import (
	"strings"

	"TPs-DDS/TP1/development/entities"
)

var usersLoginTable = map[string]string{
	"mauricio_truchet@hotmail.com": "123",
	"pepe@hotmail.com":             "321",
	"manolo@gmail.com":             "4",
}

var guideTable []entities.Guide

func CheckEmail(email string) bool {
	for emailTable, _ := range usersLoginTable {
		if strings.Compare(email, emailTable) == 0 {
			return true
		}
	}
	return false
}

func CheckPassword(email, password string) bool {
	passwordTable := usersLoginTable[email]
	return strings.Compare(password, passwordTable) == 0
}

func SaveGuide(guide entities.Guide) *entities.Guide {
	usersLoginTable[guide.Email] = guide.Password
	guideTable = append(guideTable, guide)
	return &guide
}

func GetGuide(email string) *entities.Guide {
	for _, guide := range guideTable {
		if strings.Compare(email, guide.Email) == 0 {
			return &guide
		}
	}

	return nil
}
