package coffeecatalogue4

import "time"

type Roastery struct {
	Id         int64     `db:"id"         json:"id"            gorm:"primary_key"`
	Name       string    `db:"name"       json:"name"`
	Country    string    `db:"country"    json:"country"`
	Address    string    `db:"address"    json:"address"`
	City       string    `db:"city"       json:"city"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updatedAt"`
}

type Count struct {
	Quantity int64 `json:"count"`
}
