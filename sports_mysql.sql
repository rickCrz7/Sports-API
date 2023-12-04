-- Drop database if exists
DROP SCHEMA IF EXISTS sports;

-- Create database if not exists
CREATE SCHEMA IF NOT EXISTS sports;

-- Create sports table
CREATE TABLE IF NOT EXISTS sports.sports (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sport_name VARCHAR(50) NOT NULL
);

-- Create leagues table
CREATE TABLE IF NOT EXISTS sports.leagues (
    id INT AUTO_INCREMENT PRIMARY KEY,
    league_name VARCHAR(50) NOT NULL,
    sport_id INT NOT NULL,
    FOREIGN KEY (sport_id) REFERENCES sports.sports(id) ON DELETE CASCADE
);

-- Create teams table
CREATE TABLE IF NOT EXISTS sports.teams (
    id INT AUTO_INCREMENT PRIMARY KEY,
    team_name VARCHAR(50) NOT NULL,
    sport_id INT NOT NULL,
    league_id INT,
    FOREIGN KEY (sport_id) REFERENCES sports.sports(id) ON DELETE CASCADE,
    FOREIGN KEY (league_id) REFERENCES sports.leagues(id) ON DELETE CASCADE
);

-- Create games table
CREATE TABLE IF NOT EXISTS sports.games (
    id INT AUTO_INCREMENT PRIMARY KEY,
    home_team_id INT NOT NULL,
    away_team_id INT NOT NULL,
    game_date DATE NOT NULL,
    FOREIGN KEY (home_team_id) REFERENCES sports.teams(id) ON DELETE CASCADE,
    FOREIGN KEY (away_team_id) REFERENCES sports.teams(id) ON DELETE CASCADE
);

-- Create scores table
CREATE TABLE IF NOT EXISTS sports.scores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    game_id INT NOT NULL,
    home_score INT NOT NULL,
    away_score INT NOT NULL,
    FOREIGN KEY (game_id) REFERENCES sports.games(id) ON DELETE CASCADE
);
