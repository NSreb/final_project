package migration

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"go_final_project/internal/repository"
)

func Migrate(rep *repository.Repository) error {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dbFile := viper.GetString("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	dbFile = filepath.Join(filepath.Dir(appPath), dbFile)
	_, err = os.Stat(dbFile)
	if _, err := os.Stat(dbFile); err == nil {
		log.Println("База данных уже существует.")
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	install := true

	if install {
		
		if err := rep.CreateScheduler(); err != nil {
			log.Fatal("Ошибка при создании таблицы:", err)
		} else {
			log.Println("Таблица успешно создана, создаем индекс...")
			if err := rep.CreateIDXScheduler(); err != nil {
				log.Fatal("Ошибка при создании индекса:", err)
			}
			log.Println("Индекс успешно создан.")
		}
	}

	return nil
}
