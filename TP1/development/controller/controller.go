package controller

import (
	"strings"

	"TPs-DDS/TP1/development/entities"
	"TPs-DDS/TP1/development/externalapis"
	"TPs-DDS/TP1/development/persistence"
)

func CheckEmailAndPassword(email, password string) bool {
	checkEmail := persistence.CheckEmail(email)

	if !checkEmail {
		return false
	}

	return persistence.CheckPassword(email, password)
}

func CheckData(cuit, name, lastName string) bool {
	contributor := externalapis.GetContributor(cuit)

	if contributor == nil {
		return false
	}
	return strings.ToUpper(name) == strings.ToUpper(contributor.Name) && strings.ToUpper(lastName) == strings.ToUpper(contributor.LastName)
}

func CreateGuide(cuit, name, lastName, email, password string) *entities.Guide {
	guideInDB := persistence.GetGuide(email)

	if guideInDB != nil {
		return nil
	}

	userInDB := persistence.CheckEmail(email)

	if userInDB {
		return nil
	}

	guide := entities.Guide{
		CUIT:     cuit,
		Name:     name,
		LastName: lastName,
		Email:    email,
		Password: password,
	}
	return persistence.SaveGuide(guide)
}
