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

// var server = "localhost"
// var port = 1450
// var user = "systemweb"
// var password = "1q2w3e4r"
// var database = "app_gmapmaker"

// var db *sql.DB

type MapMakerController struct {
	*revel.Controller
}

func (mm MapMakerController) GetMakers() revel.Result {
	// Get Param
	keyword := mm.Params.Query.Get("keyword")
	fmt.Printf("keyword!" + keyword + "\n")
	function := mm.Params.Query.Get("function")
	fmt.Printf("function!" + function + "\n")
	countrycode := mm.Params.Query.Get("countrycode")
	fmt.Printf("countrycode!" + countrycode + "\n")
	locode := mm.Params.Query.Get("locode")
	fmt.Printf("locode!" + locode + "\n")

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
	// Base query
	tsql := `SELECT top 100 
		m.id,
		COALESCE(m.name_th, '') [name_th],
		COALESCE(m.name_en, '') [name_en],
		case when m.latitude is null then 0 else m.latitude end [latitude],
		case when m.longitude is null then 0 else m.longitude end [longitude],
		COALESCE(m.lo_code, '') [lo_code],
		COALESCE(m.lo_code_country, '') [lo_code_country],
		COALESCE(mc.display, '') [full_country],
		case when m.truck_amount is null then 0 else m.truck_amount end [truck_amount],
		COALESCE(m.func_1, '') [func_1],
		COALESCE(m.func_2, '') [func_2],
		COALESCE(m.func_3, '') [func_3],
		COALESCE(m.func_4, '') [func_4],
		COALESCE(m.func_5, '') [func_5],
		COALESCE(m.func_6, '') [func_6],
		COALESCE(m.func_7, '') [func_7],
		COALESCE(m.func_8, '') [func_8],
		COALESCE(m.truck_type, '') [truck_type]
	FROM map_maker m
	cross apply (
		select top 1 *
		from dbo.master_country imc
		where imc.val = m.lo_code_country
	) as mc
	WHERE`

	tsql += " ( name_th like '%%" + keyword + "%%' or name_en like '%%" + keyword + "%%' or lo_code like '%%" + keyword + "%%' )"
	// Addc Condition Function
	switch f := function; f {
	case "0":
		tsql += " and func_1 = '" + f + "'"
		fmt.Printf("function func_1.0\n")
	case "1":
		tsql += " and func_1 = '" + f + "'"
		fmt.Printf("function func_1.1\n")
	case "2":
		tsql += " and func_2 = 'Y'"
		fmt.Printf("function func_2\n")
	case "3":
		tsql += " and func_3 = 'Y'"
		fmt.Printf("function func_3\n")
	case "4":
		tsql += " and func_4 = 'Y'"
		fmt.Printf("function func_4\n")
	case "5":
		tsql += " and func_5 = 'Y'"
		fmt.Printf("function func_5\n")
	case "6":
		tsql += " and func_6 = 'Y'"
		fmt.Printf("function func_6\n")
	case "7":
		tsql += " and func_7 = 'Y'"
		fmt.Printf("function func_7\n")
	case "8":
		tsql += " and func_8 = 'Y'"
		fmt.Printf("function func_8\n")
	default:
		fmt.Printf("no function\n")
	}
	// Addc Condition lo_code_country
	if countrycode != "" {
		tsql += " and lo_code_country = '" + countrycode + "'"
	}
	// Addc Condition locode
	if locode != "" {
		tsql += " and lo_code = '" + locode + "'"
	}

	tsql += ";"
	fmt.Printf("tsql!" + tsql + "\n")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	listmakers := models.ListMakers{}
	makers := []*models.Maker{}

	for rows.Next() {
		maker := new(models.Maker)
		if err := rows.Scan(
			&maker.ID,
			&maker.NameTh,
			&maker.NameEn,
			&maker.Latitude,
			&maker.Longitude,
			&maker.LoCode,
			&maker.LoCodeCountry,
			&maker.FullCountry,
			&maker.TruckAmount,
			&maker.Func1,
			&maker.Func2,
			&maker.Func3,
			&maker.Func4,
			&maker.Func5,
			&maker.Func6,
			&maker.Func7,
			&maker.Func8,
			&maker.TruckType); err != nil {
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

func (mm MapMakerController) GetMasterCountry() revel.Result {
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

	tsql := `SELECT id, val, name, display FROM master_country order by name;`

	fmt.Printf("tsql!" + tsql + "\n")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	listmasters := models.ListMasters{}
	masters := []*models.Master{}

	for rows.Next() {
		master := new(models.Master)
		if err := rows.Scan(&master.ID, &master.Val, &master.Name, &master.Display); err != nil {
			fmt.Println(err)
		}
		masters = append(masters, master)
	}

	listmasters.Masters = masters
	return mm.RenderJSON(listmasters)
}

func (mm MapMakerController) GetMasterFunction() revel.Result {
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

	tsql := `SELECT id, val, name, display FROM master_functions;`

	fmt.Printf("tsql!" + tsql + "\n")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	listmasters := models.ListMasters{}
	masters := []*models.Master{}

	for rows.Next() {
		master := new(models.Master)
		if err := rows.Scan(&master.ID, &master.Val, &master.Name, &master.Display); err != nil {
			fmt.Println(err)
		}
		masters = append(masters, master)
	}

	listmasters.Masters = masters
	return mm.RenderJSON(listmasters)
}
