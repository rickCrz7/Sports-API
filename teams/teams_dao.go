package teams

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/Sports-API/utils"
)

type Dao interface {
	getAllTeams() ([]*utils.Teams, error)
	getTeamsBySportsID(sportsID int) ([]*utils.Teams, error)
	getTeamsByLeagueID(leagueID int) ([]*utils.Teams, error)
	getTeamByID(id int) (*utils.Teams, error)
	CreateTeam(team *utils.Teams) error
	UpdateTeam(team *utils.Teams) error
	DeleteTeam(id int) error
}

type DaoImpl struct {
	conn *sql.DB
}

func NewDao(conn *sql.DB) *DaoImpl {
	return &DaoImpl{conn: conn}
}

func (dao *DaoImpl) getAllTeams() ([]*utils.Teams, error) {
	log.Println("getAllTeams")
	teams := []*utils.Teams{}
	rows, err := dao.conn.Query("SELECT * FROM teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		team := &utils.Teams{}
		err := rows.Scan(&team.ID, &team.TeamName, &team.SportsID, &team.LeagueID)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (dao *DaoImpl) getTeamsBySportsID(sportsID int) ([]*utils.Teams, error) {
	log.Println("getTeamsBySportsID")
	teams := []*utils.Teams{}
	rows, err := dao.conn.Query("SELECT * FROM teams WHERE sports_id=?", sportsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		team := &utils.Teams{}
		err := rows.Scan(&team.ID, &team.TeamName, &team.SportsID, &team.LeagueID)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (dao *DaoImpl) getTeamsByLeagueID(leagueID int) ([]*utils.Teams, error) {
	log.Println("getTeamsByLeagueID")
	teams := []*utils.Teams{}
	rows, err := dao.conn.Query("SELECT * FROM teams WHERE league_id=?", leagueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		team := &utils.Teams{}
		err := rows.Scan(&team.ID, &team.TeamName, &team.SportsID, &team.LeagueID)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}


func (dao *DaoImpl) getTeamByID(id int) (*utils.Teams, error) {
	log.Println("getTeamByID")
	team := &utils.Teams{}
	err := dao.conn.QueryRow("SELECT * FROM teams WHERE id=?", id).Scan(&team.ID, &team.TeamName, &team.SportsID, &team.LeagueID)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (dao *DaoImpl) CreateTeam(team *utils.Teams) error {
	log.Println("CreateTeam")
	_, err := dao.conn.Exec("INSERT INTO teams (team_name, sport_id, league_id) VALUES (?, ?, ?)", team.TeamName, team.SportsID, team.LeagueID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) UpdateTeam(team *utils.Teams) error {
	log.Println("UpdateTeam")
	_, err := dao.conn.Exec("UPDATE teams SET team_name=?, sport_id=?, league_id=? WHERE id=?", team.TeamName, team.SportsID, team.LeagueID, team.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) DeleteTeam(id int) error {
	log.Println("DeleteTeam")
	_, err := dao.conn.Exec("DELETE FROM teams WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}