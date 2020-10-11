package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mic3ael/pragmaticreviews/entities"
)

type VideoRepository interface {
	Save(video entities.Video)
	Update(video entities.Video)
	Delete(video entities.Video)
	FindAll() []entities.Video
	Close() error
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&entities.Video, entities.Person)
	return &database{connection: db}
}

func (db *database) Close() error {
	err := db.connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (db *database) Save(video entities.Video){
	db.connection.Create(&video)
}
func (db *database) Update(video entities.Video){
	db.connection.Save(&video)
}
func (db *database) Delete(video entities.Video){
	db.connection.Delete(&video)
}
func (db *database) FindAll() []entities.Video{
	var videos []entities.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
