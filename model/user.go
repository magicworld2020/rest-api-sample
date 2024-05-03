package model

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" form:"id" json:"id"`
	UserID   string `gorm:"type:varchar(255) binary" form:"user_id" json:"user_id"`
	Password string `gorm:"type:varchar(255) binary" form:"password" json:"password"`
	Nickname string `gorm:"type:varchar(255) binary" form:"nickname" json:"nickname"`
	Comment  string `gorm:"type:varchar(255) binary" form:"comment" json:"comment"`
}
