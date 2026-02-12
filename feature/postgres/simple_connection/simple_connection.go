package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	con, err := pgx.Connect(ctx, "postgres://postgres:arina31072008@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	err = con.Ping(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Успешно подключился")
}
