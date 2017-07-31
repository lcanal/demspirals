package models

//PlayerStats tracks the stats we care about. Generally for fantasy points
type PlayerStats struct {
	PID               string `json:"pid"`
	Runs              int64  `json:"runs,omitempty"`
	Passes            int64  `json:"passes,omitempty"`
	Receptions        int64  `json:"receptions,omitempty"`
	Receptionyards    int64
	Receptiontargets  int64
	Receptionyards10p int64
	Receptionyards20p int64
	Receptionyards30p int64
	Receptionyards40p int64
	Receptionyards50p int64
	Rushyards         int64
	Rushattempts      int64
	Rushyards10p      int64
	Rushyards20p      int64
	Rushyards30p      int64
	Rushyards40p      int64
	Rushyards50p      int64
	Touchdownpasses   int64
	Touchdownrushes   int64
	Fumbles           int64
}
