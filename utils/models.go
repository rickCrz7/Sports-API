package utils

import (
	"time"
)

type Sports struct {
	ID        int    `json:"id"`
	SportName string `json:"sport_name"`
}

type Leagues struct {
	ID         int    `json:"id"`
	LeagueName string `json:"league_name"`
	SportsID   int    `json:"sports_id"`
}

type Teams struct {
	ID       int    `json:"id"`
	TeamName string `json:"team_name"`
	SportsID int    `json:"sports_id"`
	LeagueID int    `json:"league_id"`
}

type Games struct {
	ID         int       `json:"id"`
	HomeTeamID int       `json:"home_team_id"`
	AwayTeamID int       `json:"away_team_id"`
	GameDate   time.Time `json:"game_date"`
}

type Scores struct {
	ID        int `json:"id"`
	GameID    int `json:"game_id"`
	HomeScore int `json:"home_score"`
	AwayScore int `json:"away_score"`
}

type GameScore struct {
	ID        int       `json:"id"`
	HomeTeam  string    `json:"home_team"`
	AwayTeam  string    `json:"away_team"`
	GameDate  time.Time `json:"game_date"`
	HomeScore int       `json:"home_score"`
	AwayScore int       `json:"away_score"`
}
