package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Update(ctx context.Context, conn *pgx.Conn, task TasksModel) error {
	sqlQuary := `
UPDATE tasks
SET title=$1, description=$2, completed=$3, created_at=$4, completed_at=$5
WHERE id = $6;
`

	_, err := conn.Exec(ctx,
		sqlQuary,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.Id,
	)
	return err

}
