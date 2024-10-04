package main

import (
	"encoding/json"
	"fmt"
	"github.com/punkgoden/song_library/app/pkg/config"
	"github.com/punkgoden/song_library/app/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
)

type Song struct {
	ID         int    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(200);unique;not null" json:"name"`
	Group      string `gorm:"type:varchar(200);not null" json:"group"`
	Text       string `gorm:"type:text" json:"text"`
	Listenings int    `gorm:"type:int" json:"listenings"`
}

type Object struct {
	Songs []Song `json:"songs"`
}

func main() {
	// Init Logger
	logging.Init(false)
	log := logging.GetLogger()
	log.Infoln("Connect logger successfully!")

	// Init Config
	cfg := config.GetConfig()
	log.Infoln("Connect config successfully!")
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.AppConfig.PostgresSQL.Username,
		cfg.AppConfig.PostgresSQL.Password,
		cfg.AppConfig.PostgresSQL.Host,
		cfg.AppConfig.PostgresSQL.Port,
		cfg.AppConfig.PostgresSQL.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Error connection database: %v", err)
	}

	jsonFile, err := os.Open("test_data.json")
	if err != nil {
		log.Panicln(err)
	}
	log.Infoln("Successfully Opened test_data.json")
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(jsonFile)
	byteValue, _ := io.ReadAll(jsonFile)

	var obj Object

	if err := json.Unmarshal(byteValue, &obj); err != nil {
		log.Panicln(err)
	}

	for _, song := range obj.Songs {
		if err := db.Create(&song).Error; err != nil {
			log.Panicln(err)
		}
	}

	log.Infoln("Successfully Created songs in test_data.json")
}
