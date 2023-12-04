package games

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/Sports-API/utils"
)

type ScoresDao interface {
	createScore(score *utils.Scores) error
	updateScore(score *utils.Scores) error
	deleteScore(id int) error
}

type ScoresDaoImpl struct {
	conn *sql.DB
}

func NewScoresDao(conn *sql.DB) *ScoresDaoImpl {
	return &ScoresDaoImpl{conn: conn}
}

func (dao *ScoresDaoImpl) createScore(score *utils.Scores) error {
	log.Println("createScore")
	_, err := dao.conn.Exec("INSERT INTO scores (game_id, home_score, away_score) VALUES (?, ?, ?)", score.GameID, score.HomeScore, score.AwayScore)
	if err != nil {
		return err
	}
	return nil
}

func (dao *ScoresDaoImpl) updateScore(score *utils.Scores) error {
	log.Println("updateScore")
	_, err := dao.conn.Exec("UPDATE scores SET home_score=?, away_score=? WHERE id=?", score.HomeScore, score.AwayScore, score.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *ScoresDaoImpl) deleteScore(id int) error {
	log.Println("deleteScore")
	_, err := dao.conn.Exec("DELETE FROM scores WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
