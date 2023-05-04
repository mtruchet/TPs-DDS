package persistence

import "strings"

var usersLoginTable = map[string]string{
	"mauricio_truchet@hotmail.com": "123",
	"pepe@hotmail.com":             "321",
	"manolo@gmail.com":             "4",
}

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
