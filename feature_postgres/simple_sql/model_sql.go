package simple_sql

import "time"

type TasksModel struct {
	Id          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
