package coffeecatalogue4

import "time"

type User struct {
	Id        int64     `db:"id"         json:"id" gorm:"primary_key"`
	Login     string    `db:"login"      json:"login"`
	Password  string    `db:"password"   json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
