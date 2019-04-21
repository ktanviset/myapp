package controllers

import (
	"fmt"
	//"myapp/app"
	"myapp/app/models"
	"github.com/revel/revel"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type MapMakerController struct {
	*revel.Controller
}

func (mm MapMakerController) GetMakers() revel.Result {
	sqlsm := "SELECT latitude, longitude, name from map_maker"

	db, dberrcon := sql.Open("mysql", "root:1234@tcp(localhost:3306)/app_gmapmaker")

	if dberrcon != nil{
		fmt.Println("DB Error", dberrcon)
	}

	rows, dberrquery := db.Query(sqlsm)
	
	if dberrquery != nil{
		fmt.Println("DB Error", dberrquery)
	}

	fmt.Println(rows)

	listmakers := models.ListMakers{}
	makers := []*models.Maker{}

	for rows.Next() {
		maker := new(models.Maker)
		if err := rows.Scan(&maker.Latitude, &maker.Longitude, &maker.Name); err != nil {
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