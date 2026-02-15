package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	task TasksModel,
) error {

	sqlQuary := `
	INSERT INTO tasks (title,description,completed,created_at)
	VALUES ($1, $2, $3, $4);
	`

	_, err := conn.Exec(ctx,
		sqlQuary,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
	)

	return err

}
