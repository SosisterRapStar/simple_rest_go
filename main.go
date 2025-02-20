package main

import (
	"context"
	"first-proj/connection"
	"first-proj/domain"
	"first-proj/services/postgres"
	"fmt"
)

func main() {
	pool := connection.NewPool()
	pg := postgres.NewPostgres(pool)
	note := &domain.Note{Title: "govno", Content: "govno"}
	fmt.Println(pg.CreateNote(context.Background(), note))
	title := "GOVNO2"

	updateNote := &domain.UpdateNote{
		Title: &title,
	}
	fmt.Println(pg.UpdateNote(context.Background(), updateNote, "019510a3-aa07-727e-99a2-b611287eef4c"))
	fmt.Println(pg.DeleteNote(context.Background(), "019510a3-aa07-727e-99a2-b611287eef4c"))
}
