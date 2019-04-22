package controllers

import (
	"context"
	"fmt"
	"log"

	"database/sql"
	"myapp/app/models"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/revel/revel"
)

var server = "localhost"
var port = 1450
var user = "systemweb"
var password = "1q2w3e4r"
var database = "app_gmapmaker"

var db *sql.DB

type MapMakerController struct {
	*revel.Controller
}

func (mm MapMakerController) GetMakers() revel.Result {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Execute query
	tsql := fmt.Sprintf("SELECT id, name_th, name_en, latitude, longitude FROM map_maker;")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	listmakers := models.ListMakers{}
	makers := []*models.Maker{}

	for rows.Next() {
		maker := new(models.Maker)
		if err := rows.Scan(&maker.ID, &maker.NameTh, &maker.NameEn, &maker.Latitude, &maker.Longitude); err != nil {
			fmt.Println(err)
		}
		makers = append(makers, maker)
	}

	// maker := new(models.Maker)
	// maker.Name = "LINE Village Bangkok"
	// maker.Latitude = 13.7395
	// maker.Longitude = 100.549

	// makers = append(makers, maker)
	listmakers.Makers = makers
	return mm.RenderJSON(listmakers)
}

func (mm MapMakerController) SaveMaker() revel.Result {

	return nil
}
