package models

type User struct {
	ID       uint64  `gorm:"autoIncrement;primaryKey" json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	NIK      *string `json:"nik"`
}
