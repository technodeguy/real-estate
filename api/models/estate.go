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

func (e *Estate) CreateEstate(db *sql.DB) (uint32, error) {
	var lastEstateId uint32

	err := db.QueryRow("CALL create_estate(?, ?, ?, ?, ?)", e.Title, e.Price, e.IsRent, e.CityId, e.UserId).Scan(&lastEstateId)

	if err != nil {
		return 0, err
	}

	return lastEstateId, nil
}

func (e *Estate) SetSellingDate(db *sql.DB, estateId uint32) {
	db.QueryRow("UPDATE estate SET selling_date = NOW(), update_dt = NOW() WHERE id = ?", estateId)
}

func (e *Estate) GetEstatesByCityId(db *sql.DB, cityId int) ([]Estate, error) {
	rows, err := db.Query("CALL sel_estates_by_city_id(?)", cityId)

	if err != nil {
		log.Println("Error while retrieving all local estates")
		return nil, err
	}

	estates := make([]Estate, 0)

	for rows.Next() {
		estate := Estate{}

		err = rows.Scan(&estate.Id, &estate.Price, &estate.IsRent, &estate.CityId, &estate.UserId, &estate.SellingDate, &estate.InsertDt)

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
