package models

import (
	"database/sql"
	"log"
)

type Estate struct {
	Id          int
	Title       string
	Price       int
	IsRent      bool
	CityId      int
	UserId      int
	SellingDate string
	InsertDt    string
	UpdateDt    string
}

func (e *Estate) FindEstatesByCityId(db *sql.DB) ([]Estate, error) {
	rows, err := db.Query("CALL sel_estates_by_city_id(?)", e.CityId)

	if err != nil {
		log.Println("Error while retrieving all local estates", err)
		return nil, err
	}

	estates := make([]Estate, 0)

	for rows.Next() {
		estate := Estate{}

		err = rows.Scan(&estate.Id, &estate.Title, &estate.Price, &estate.IsRent)

		if err != nil {
			log.Println("Error by parsing - getAllEstates")
			return nil, err
		}

		estates = append(estates, estate)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return estates, nil
}

func (e *Estate) CreateEstate(db *sql.DB) (uint32, error) {
	var lastEstateId uint32

	err := db.QueryRow("CALL create_estate(?, ?, ?, ?, ?)", e.Title, e.Price, e.IsRent, e.CityId, e.UserId).Scan(&lastEstateId)

	if err != nil {
		return 0, err
	}

	return lastEstateId, nil
}

func (e *Estate) SetSellingDate(db *sql.DB) {
	db.QueryRow("UPDATE estate SET selling_date = NOW() WHERE id = ?", e.Id)
}
