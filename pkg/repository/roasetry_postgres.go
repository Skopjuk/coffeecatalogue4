package repository

import (
	"coffeecatalogue4"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RoasteryPostgres struct {
	db *sqlx.DB
}

func NewRoasteryPostgres(db *sqlx.DB) *RoasteryPostgres {
	return &RoasteryPostgres{db: db}
}

func (r *RoasteryPostgres) Create(userId int, roastery coffeecatalogue4.Roastery) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createRoasteryQuery := fmt.Sprintf("INSERT INTO %s (name, country, address, city) VALUES ($1, $2, $3, $4) RETURNING id", roasteriesTable)
	row := tx.QueryRow(createRoasteryQuery, roastery.Name, roastery.Country, roastery.Address, roastery.City)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, roastery_id) VALUES ($1, $2)", usersTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
	}

	return id, tx.Commit()
}
