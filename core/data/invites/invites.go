package invites

import (
	"log"
	database "mysql"
)

func IsValidInviteCode(inviteCode string) bool {
	statement, err := database.Db.Prepare("SELECT id FROM invite_codes WHERE code = ?")
	if err != nil {
		log.Fatal(err)
	}

	row, err := statement.Query(inviteCode)

	if err != nil {
		log.Fatal(err)
	}

	return row.Next()
}

type InvalidInviteCodeError struct{}

func (m *InvalidInviteCodeError) Error() string {
	return "invalid invite code"
}
