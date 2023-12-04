package sports

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/Sports-API/utils"
)

type Dao interface {
	getAllSports() ([]*utils.Sports, error)
	getSportByID(id int) (*utils.Sports, error)
	CreateSport(sport *utils.Sports) error
	UpdateSport(sport *utils.Sports) error
	DeleteSport(id int) error
}

type DaoImpl struct {
	conn *sql.DB
}

func NewDao(conn *sql.DB) *DaoImpl {
	return &DaoImpl{conn: conn}
}

func (dao *DaoImpl) getAllSports() ([]*utils.Sports, error) {
	log.Println("getAllSports")
	sports := []*utils.Sports{}
	rows, err := dao.conn.Query("SELECT * FROM sports")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		sport := &utils.Sports{}
		err := rows.Scan(&sport.ID, &sport.SportName)
		if err != nil {
			return nil, err
		}
		sports = append(sports, sport)
	}
	return sports, nil
}

func (dao *DaoImpl) getSportByID(id int) (*utils.Sports, error) {
	log.Println("getSportByID")
	sport := &utils.Sports{}
	err := dao.conn.QueryRow("SELECT * FROM sports WHERE id=?", id).Scan(&sport.ID, &sport.SportName)
	if err != nil {
		return nil, err
	}
	return sport, nil
}

func (dao *DaoImpl) CreateSport(sport *utils.Sports) error {
	log.Println("CreateSport")
	_, err := dao.conn.Exec("INSERT INTO sports (sport_name) VALUES (?)", sport.SportName)
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) UpdateSport(sport *utils.Sports) error {
	log.Println("UpdateSport")
	_, err := dao.conn.Exec("UPDATE sports SET sport_name=? WHERE id=?", sport.SportName, sport.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) DeleteSport(id int) error {
	log.Println("DeleteSport")
	_, err := dao.conn.Exec("DELETE FROM sports WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
