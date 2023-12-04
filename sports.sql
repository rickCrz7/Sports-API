DROP SCHEMA IF EXISTS sports CASCADE;
CREATE SCHEMA IF NOT EXISTS sports;

CREATE TABLE IF NOT EXISTS sports.sports (
  id SERIAL PRIMARY KEY,
  sport_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS sports.leagues (
  id SERIAL PRIMARY KEY,
  league_name VARCHAR(50) NOT NULL,
  sport_id SERIAL NOT NULL,
  FOREIGN KEY (sport_id) REFERENCES sports.sports(id)
);

CREATE TABLE IF NOT EXISTS sports.teams (
  id SERIAL PRIMARY KEY,
  team_name VARCHAR(50) NOT NULL,
  sport_id SERIAL NOT NULL,
  FOREIGN KEY (sport_id) REFERENCES sports.sports(id)
);

CREATE TABLE IF NOT EXISTS sports.games (
  id SERIAL PRIMARY KEY,
  home_team_id SERIAL NOT NULL,
  away_team_id SERIAL NOT NULL,
  game_date DATE NOT NULL,
  FOREIGN KEY (home_team_id) REFERENCES sports.teams(id),
  FOREIGN KEY (away_team_id) REFERENCES sports.teams(id)
);

CREATE TABLE IF NOT EXISTS sports.scores (
  id SERIAL PRIMARY KEY,
  game_id SERIAL NOT NULL,
  home_score INT NOT NULL,
  away_score INT NOT NULL,
  FOREIGN KEY (game_id) REFERENCES sports.games(id)
);

CREATE TABLE IF NOT EXISTS sports.game_stats (
  id SERIAL PRIMARY KEY,
  game_id SERIAL NOT NULL,
  team_id SERIAL NOT NULL,
  assists INT NOT NULL,
  rebounds INT NOT NULL,
  steals INT NOT NULL,
  blocks INT NOT NULL,
  turnovers INT NOT NULL,
  field_goal_percentage DECIMAL(4,3) NOT NULL,
  three_point_percentage DECIMAL(4,3) NOT NULL,
  free_throw_percentage DECIMAL(4,3) NOT NULL,
  FOREIGN KEY (game_id) REFERENCES sports.games(id),
  FOREIGN KEY (team_id) REFERENCES sports.teams(id)
);


-- CREATE TABLE IF NOT EXISTS sports.players (
--   id SERIAL PRIMARY KEY,
--   player_name VARCHAR(100) NOT NULL,
--   team_id SERIAL NOT NULL,
--   s VARCHAR(50) NOT NULL,
--   FOREIGN KEY (team_id) REFERENCES sports.teams(id)
-- );

-- CREATE TABLE IF NOT EXISTS sports.player_stats_nba (
--   id SERIAL PRIMARY KEY,
--   player_id SERIAL NOT NULL,
--   year VARCHAR(50) NOT NULL,
--   games_played INT NOT NULL,
--   points_per_game DECIMAL(4,3) NOT NULL,
--   assists_per_game DECIMAL(4,3) NOT NULL,
--   rebounds_per_game DECIMAL(4,3) NOT NULL,
--   steals_per_game DECIMAL(4,3) NOT NULL,
--   blocks_per_game DECIMAL(4,3) NOT NULL,
--   turnovers_per_game DECIMAL(4,3) NOT NULL,
--   field_goal_percentage DECIMAL(4,3) NOT NULL,
--   three_point_percentage DECIMAL(4,3) NOT NULL,
--   free_throw_percentage DECIMAL(4,3) NOT NULL,
--   FOREIGN KEY (player_id) REFERENCES sports.players(id)
-- );

-- -- Create picks table
-- CREATE TABLE IF NOT EXISTS sports.picks (
--   id INT PRIMARY KEY,
--   user_id INT NOT NULL,
--   nfl_score_id INT NOT NULL,
--   pick_home_team BOOLEAN NOT NULL,
--   FOREIGN KEY (user_id) REFERENCES sports.users(id),
--   FOREIGN KEY (nfl_score_id) REFERENCES sports.nfl_scores(id)
-- );

-- -- Create player_picks table
-- CREATE TABLE IF NOT EXISTS sports.player_picks (
--   id INT PRIMARY KEY,
--   user_id INT NOT NULL,
--   player_stat_id INT NOT NULL,
--   FOREIGN KEY (user_id) REFERENCES sports.users(id),
--   FOREIGN KEY (player_stat_id) REFERENCES sports.player_stats(id)
-- );

INSERT INTO sports.sport VALUES 
