package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"go_final_project/internal/db"
	"go_final_project/internal/handler"
	"go_final_project/internal/repository"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/spf13/viper"
)

func main() {
	webDir := "./web"

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка получения рабочего каталога: %v", err)
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(filepath.Join(dir, "cmd", "service"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	//port := os.Getenv("TODO_PORT")
	port := viper.GetString("TODO_PORT")
	if port == "" {
		port = "7540" // Значение по умолчанию
	}

	db.New()
	rep := repository.New(db.New())
	migration(rep)

	handler := handler.New(rep)

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir(webDir)))
	r.Post("/api/task", handler.AddTask)
	r.Get("/api/tasks", handler.GetList)
	r.Get("/api/task", handler.GetTask)
	r.Put("/api/task", handler.EditTask)
	r.Get("/api/nextdate", handler.NextDate)
	r.Post("/api/task/done", handler.DoneTask)
	r.Delete("/api/task", handler.DeleteTask)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func migration(rep *repository.Repository) {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dbFile := viper.GetString("TODO_DBFILE")

	if dbFile == "" {
		dbFile = "scheduler.db" // Значение по умолчанию
	}

	dbFile = filepath.Join(filepath.Dir(appPath), dbFile)
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	if install {
		if err := rep.CreateScheduler(); err != nil {
			log.Fatal(err)
		} else {
			rep.CreateIDXScheduler()
		}
	}

}
