package services

import (
	"backend/models"
)

func (db *serviceProxy) GetVideos() (videos []models.Video) {
	db.mysql.Find(&videos)
	return
}