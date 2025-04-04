package models

// Comment model.
type Comment struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
}
