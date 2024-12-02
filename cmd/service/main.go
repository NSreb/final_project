package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"go_final_project/internal/db"
	"go_final_project/internal/db/migration"
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

	db, err := db.New()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Ошибка при закрытии соединения:", err)
		}
	}()

	rep := repository.New(db)

	if err := migration.Migrate(rep); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
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

	log.Printf("Приложение запущено на порту: %s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
