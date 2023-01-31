package coffeecatalogue4

import "time"

type Coffee struct {
	Id          int64     `db:"id"          json:"id"           gorm:"primary_key"`
	Origin      string    `db:"origin"      json:"origin"`
	Variety     string    `db:"variety"     json:"variety"`
	Processing  string    `db:"processing"  json:"processing"`
	Descriptors string    `db:"descriptors" json:"descriptors"`
	CreatedAt   time.Time `db:"created_at"  json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"  json:"updated_at"`
	RoasteryId  int64     `db:"roastery_id" json:"roastery_id"`
}
