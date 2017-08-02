package models

//Stat tracks the stats we care about. Generally for fantasy points
type Stat struct {
	PID               string `json:"pid" gorm:"primary_key"`
	Gamesplayed       int64  `json:"games_played"`
	Receptions        int64  `json:"receptions"`
	Receptionyards    int64  `json:"receptions_yards"`
	Receptiontargets  int64  `json:"receptions_targets"`
	Receptionyards10p int64  `json:"receptions_yards_10_plus"`
	Receptionyards20p int64  `json:"receptions_yards_20_plus"`
	Receptionyards30p int64  `json:"receptions_yards_30_plus"`
	Receptionyards40p int64  `json:"receptions_yards_40_plus"`
	Receptionyards50p int64  `json:"receptions_yards_50_plus"`
	Rushyards         int64  `json:"rush_yards"`
	Rushattempts      int64  `json:"rush_attempts"`
	Rushyards10p      int64  `json:"rush_yards_10_plus"`
	Rushyards20p      int64  `json:"rush_yards_20_plus"`
	Rushyards30p      int64  `json:"rush_yards_30_plus"`
	Rushyards40p      int64  `json:"rush_yards_40_plus"`
	Rushyards50p      int64  `json:"rush_yards_50_plus"`
	Passyards         int64  `json:"pass_yards"`
	Passattempts      int64  `json:"pass_attempts"`
	Passyards10p      int64  `json:"pass_yards_10_plus"`
	Passyards20p      int64  `json:"pass_yards_20_plus"`
	Passyards30p      int64  `json:"pass_yards_30_plus"`
	Passyards40p      int64  `json:"pass_yards_40_plus"`
	Passyards50p      int64  `json:"pass_yards_50_plus"`
	Touchdownpasses   int64  `json:"td_passes"`
	Touchdownrushes   int64  `json:"td_rushes"`
	Fumbles           int64  `json:"fumbles"`
}
