package models

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	// Avatar      string `json:"avatar"`
}

func (u *User) hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *User) VerifyAndCompare(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}

	return true
}

func (u *User) FindAllUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query("SELECT id, nickname, phone_number FROM user")

	if err != nil {
		log.Println("Error while retrieving all users")
		return nil, err
	}

	// var users []User
	users := make([]User, 0)

	for rows.Next() {
		user := User{}

		err = rows.Scan(&user.Id, &user.Nickname, &user.PhoneNumber)

		if err != nil {
			log.Println("Error by parsing - findAllUsers")
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) CreateUser(db *sql.DB) (uint32, error) {

	hashedPassword, err := u.hash(u.Password)

	if err != nil {
		return 0, err
	}

	u.Password = hashedPassword

	var lastUserId uint32

	err = db.QueryRow("CALL create_user(?, ?, ?)", u.Nickname, u.Password, u.PhoneNumber).Scan(&lastUserId)

	if err != nil {
		return 0, err
	}

	return lastUserId, nil
}

func (u *User) SaveUserAvatar(db *sql.DB, id int, avatar string) {
	db.QueryRow("UPDATE user SET avatar = ?, update_dt = NOW() WHERE id = ?", avatar, id)
}

func (u *User) FindUserByNickname(db *sql.DB, nickname string) (*User, error) {

	err := db.QueryRow("SELECT id, password FROM user WHERE nickname = ?", nickname).Scan(&u.Id, &u.Nickname)

	if err != nil {
		return &User{}, err
	}

	return u, nil
}
