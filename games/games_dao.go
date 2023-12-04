package games

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/Sports-API/utils"
)

type Dao interface {
	getGames() ([]*utils.GameScore, error)
	getGameByID(id int) (*utils.GameScore, error)
	getGamesByTeamID(id int) ([]*utils.GameScore, error)
	CreateGame(game *utils.Games) error
	UpdateGame(game *utils.Games) error
	DeleteGame(id int) error
}

type DaoImpl struct {
	conn *sql.DB
}

func NewDao(conn *sql.DB) *DaoImpl {
	return &DaoImpl{conn: conn}
}

func (dao *DaoImpl) getGames() ([]*utils.GameScore, error) {
	log.Println("getGames")
	games := []*utils.GameScore{}
	rows, err := dao.conn.Query(`
	SELECT 
	g.id,
    t1.team_name AS home_team, 
    t2.team_name AS away_team, 
    g.game_date, 
    COALESCE(s.home_score, 0) AS home_score, 
    COALESCE(s.away_score, 0) AS away_score
	FROM 
    sports.games g
	JOIN 
    sports.teams t1 ON g.home_team_id = t1.id
	JOIN 
    sports.teams t2 ON g.away_team_id = t2.id
	Left JOIN 
    sports.scores s ON g.id = s.game_id	
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		game := &utils.GameScore{}
		err := rows.Scan(&game.ID, &game.HomeTeam, &game.AwayTeam, &game.GameDate, &game.HomeScore, &game.AwayScore)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (dao *DaoImpl) getGameByID(id int) (*utils.GameScore, error) {
	log.Println("getGameByID")
	game := &utils.GameScore{}
	err := dao.conn.QueryRow(`
	SELECT 
	g.id,
    t1.team_name AS home_team, 
    t2.team_name AS away_team, 
    g.game_date, 
    COALESCE(s.home_score, 0) AS home_score, 
    COALESCE(s.away_score, 0) AS away_score
	FROM 
    sports.games g
	JOIN 
    sports.teams t1 ON g.home_team_id = t1.id
	JOIN 
    sports.teams t2 ON g.away_team_id = t2.id
	Left JOIN 
    sports.scores s ON g.id = s.game_id
	where
    g.id=?`, id).Scan(&game.ID, &game.HomeTeam, &game.AwayTeam, &game.GameDate, &game.HomeScore, &game.AwayScore)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (dao *DaoImpl) getGamesByTeamID(id int) ([]*utils.GameScore, error) {
	log.Println("getGamesByTeamID")
	games := []*utils.GameScore{}
	rows, err := dao.conn.Query(`SELECT 
	g.id,
    t1.team_name AS home_team, 
    t2.team_name AS away_team, 
    g.game_date, 
    COALESCE(s.home_score, 0) AS home_score, 
    COALESCE(s.away_score, 0) AS away_score
	FROM 
    sports.games g
	JOIN 
    sports.teams t1 ON g.home_team_id = t1.id
	JOIN 
    sports.teams t2 ON g.away_team_id = t2.id
	Left JOIN 
    sports.scores s ON g.id = s.game_id
	Where 
	t1.id =? OR t2.id =?;
`, id, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		game := &utils.GameScore{}
		err := rows.Scan(&game.ID, &game.HomeTeam, &game.AwayTeam, &game.GameDate, &game.HomeScore, &game.AwayScore)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (dao *DaoImpl) CreateGame(game *utils.Games) error {
	log.Println("CreateGame")
	_, err := dao.conn.Exec("INSERT INTO games (home_team_id, away_team_id, game_date) VALUES (?, ?, ?)", game.HomeTeamID, game.AwayTeamID, game.GameDate.Format("2006-01-02"))
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) UpdateGame(game *utils.Games) error {
	log.Println("UpdateGame")
	_, err := dao.conn.Exec("UPDATE games SET home_id=?, away_id=?, game_date=? WHERE id=?", game.HomeTeamID, game.AwayTeamID, game.GameDate.Format("2006-01-02"), game.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *DaoImpl) DeleteGame(id int) error {
	log.Println("DeleteGame")
	_, err := dao.conn.Exec("DELETE FROM games WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
