package main

import (
	"first-proj/appconfig"
	"first-proj/domain"
	"first-proj/services/connections"
	"first-proj/services/postgres"
	"first-proj/transport/http"

	"fmt"
)

var config = appconfig.MustLoad()

type DependencyContainer struct {
	noteService domain.NoteService
}

var DI = DependencyContainer{
	noteService: postgres.NewPostgres(connections.NewPool(int32(config.Storage.MaxConns), int32(config.Storage.MinConns), config.Storage.Url)),
}

func main() {
	// 	pool := connection.NewPool()
	// 	pg := postgres.NewPostgres(pool)
	// 	note := &domain.Note{Title: "govno", Content: "govno"}
	// 	fmt.Println(pg.CreateNote(context.Background(), note))
	// 	title := "GOVNO2"

	//	updateNote := &domain.UpdateNote{
	//		Title: &title,
	//	}
	//
	// fmt.Println(pg.UpdateNote(context.Background(), updateNote, "019510a3-aa07-727e-99a2-b611287eef4c"))
	// fmt.Println(pg.DeleteNote(context.Background(), "019510a3-aa07-727e-99a2-b611287eef4c"))
	config := appconfig.MustLoad()
	fmt.Println(config)
	server := http.NewServer(config.Address)
	handlers := http.New
}
