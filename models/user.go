package models

// User model with minimal fields letting MySQL handle the schema.
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // Do not expose the password in responses.
}
