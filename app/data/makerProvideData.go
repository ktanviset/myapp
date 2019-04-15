package data

import (
	"myapp/app/models"
)

type MakerProvideData interface {
	GetMakers() []*models.Maker
	SaveMaker(maker *models.Maker) error
}

type DBMaker struct {
	makers []*models.Maker
}

var dbMaker *DBMaker
