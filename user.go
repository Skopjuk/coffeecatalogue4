package coffeecatalogue4

import "time"

type User struct {
	Id        int64     `db:"id"         json:"id" db:"id"`
	Login     string    `db:"login"      json:"login" binding:"required"`
	Password  string    `db:"password"   json:"password" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
