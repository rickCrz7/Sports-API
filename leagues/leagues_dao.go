package leagues

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/Sports-API/utils"
)

type Dao interface {
	getAllLeagues() ([]*utils.Leagues, error)
	getLeagueBySportID(id int) ([]*utils.Leagues, error)
	getLeagueByID(id int) (*utils.Leagues, error)
	CreateLeague(league *utils.Leagues) error
	UpdateLeague(league *utils.Leagues) error
	DeleteLeague(id int) error
}

type DaoImpl struct {
	conn *sql.DB
}

func NewDao(conn *sql.DB) *DaoImpl {
	return &DaoImpl{
		conn: conn,
	}
}

func (dao *DaoImpl) getAllLeagues() ([]*utils.Leagues, error) {
	log.Println("getAllLeagues")
	leagues := []*utils.Leagues{}
	rows, err := dao.conn.Query("SELECT * FROM leagues")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		league := &utils.Leagues{}
		err := rows.Scan(&league.ID, &league.LeagueName, &league.SportsID)
		if err != nil {
			return nil, err
		}
		leagues = append(leagues, league)
	}
	return leagues, nil
}

func (dao *DaoImpl) getLeagueBySportID(id int) ([]*utils.Leagues, error) {
	log.Println("getLeagueBySportID")
	leagues := []*utils.Leagues{}
	rows, err := dao.conn.Query("SELECT * FROM leagues WHERE sport_id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		league := &utils.Leagues{}
		err := rows.Scan(&league.ID, &league.LeagueName, &league.SportsID)
		if err != nil {
			return nil, err
		}
		leagues = append(leagues, league)
	}
	return leagues, nil
}

func (dao *DaoImpl) getLeagueByID(id int) (*utils.Leagues, error) {
	log.Println("getLeagueByID")
	league := &utils.Leagues{}
	err := dao.conn.QueryRow("SELECT * FROM leagues WHERE id=?", id).Scan(&league.ID, &league.LeagueName, &league.SportsID)
	if err != nil {
		return nil, err
	}
	return league, nil
}

func (dao *DaoImpl) CreateLeague(league *utils.Leagues) error {
	log.Println("CreateLeague")
	_, err := dao.conn.Exec("INSERT INTO leagues (league_name, sport_id) VALUES (?, ?)", league.LeagueName, league.SportsID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) UpdateLeague(league *utils.Leagues) error {
	log.Println("UpdateLeague")
	_, err := dao.conn.Exec("UPDATE leagues SET league_name=?, sport_id=? WHERE id=?", league.LeagueName, league.SportsID, league.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) DeleteLeague(id int) error {
	log.Println("DeleteLeague")
	_, err := dao.conn.Exec("DELETE FROM leagues WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
