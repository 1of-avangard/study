package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(conn, ctx); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.Check(ctx, conn, 10, 0)
	if err != nil {
		panic(err)
	}

	for _, v := range tasks {
		if v.Id == 3 {
			v.Title = "покормить собаку"
			v.Description = "30 корма"
			v.Completed = true
			time := time.Now()
			v.CompletedAt = &time

			if err := simple_sql.Update(ctx, conn, v); err != nil {
				panic(err)
			}
			break
		}

	}

	fmt.Println("succeed")

}
