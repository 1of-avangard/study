package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Check(ctx context.Context,
	conn *pgx.Conn,
	colvStrok int,
	offset int,

) ([]TasksModel, error) {

	sqlQuary := `
SELECT * FROM tasks
ORDER BY id ASC
LIMIT $1
OFFSET $2
`
	rows, err := conn.Query(ctx, sqlQuary, colvStrok, offset)
	if err != nil {
		return nil, err
	}
	//conn.Exec()
	defer rows.Close()
	tasks := make([]TasksModel, 0)
	for rows.Next() {
		var task TasksModel

		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		//PrintTasks(task)
	}
	return tasks, nil
}

func PrintTasks(task TasksModel) {
	fmt.Println("----------------------------------")
	fmt.Println("id:", task.Id)
	fmt.Println("title:", task.Title)
	fmt.Println("description:", task.Description)
	fmt.Println("completed:", task.Completed)
	fmt.Println("created_at:", task.CreatedAt)
	fmt.Println("completed_at:", task.CompletedAt)
}
