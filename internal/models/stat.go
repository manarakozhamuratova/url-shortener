package models

import "time"

type Stat struct {
	URLID        int64     `json:"url_id" db:"url_id"`
	VisitedCount int       `json:"visited_count" db:"visited_count"`
	VisitedAt    time.Time `json:"visited_at" db:"visited_at"`
}

type UpsertStat struct {
	URLID     int64
	VisitedAt time.Time
}
