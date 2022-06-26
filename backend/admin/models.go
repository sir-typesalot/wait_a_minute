package adminModel

import (
	"wait_a_minute/backend/util"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Admin_ID int
	Username string
	Password string
	Email string
	Auth_level int
}

func CreateAdmin(username, password, email string, authLevel int) int {
	hash, _ := hashPassword(password)

	db := util.GetConnection()
	query := `INSERT INTO admin 
		(username, password, email, auth_level) 
		VALUES (?, ?, ?, ?)`
	result, _ := db.Exec(query, username, hash, email, authLevel)

	rowsAff, _ := result.RowsAffected()
	if rowsAff == 1 {
		return 200
	} else {
		return 500
	}
}

func LogIn(username, password string) (string, int) {
	hash := "pass hash from DB"
	checkPassword := checkPasswordHash(password, hash)
	if checkPassword {
		return "token", 200
	} else {
		return "Invalid", 404
	}
}

func LogOut() {

}

func IsAdmin() {

}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

