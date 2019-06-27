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

type TaxpayerController struct {
	*revel.Controller
}

func (tax TaxpayerController) GetTaxpayer() revel.Result {
	// Get Param
	nid := tax.Params.Query.Get("nid")
	fmt.Printf("nid!" + nid + "\n")

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

	tsql := `SELECT TOP 1
		COALESCE(NID, '') [NID],
		COALESCE(BranchNumber, '') [BranchNumber],
		COALESCE(BranchTitle, '') [BranchTitle],
		COALESCE(BranchName, '') [BranchName],
		COALESCE(BuildingName, '') [BuildingName],
		COALESCE(RoomNumber, '') [RoomNumber],
		COALESCE(FloorNumber, '') [FloorNumber],
		COALESCE(VillageName, '') [VillageName],
		COALESCE(HouseNumber, '') [HouseNumber],
		COALESCE(MooNumber, '') [MooNumber],
		COALESCE(SoiName, '') [SoiName],
		COALESCE(StreetName, '') [StreetName],
		COALESCE(ThumbolName, '') [ThumbolName],
		COALESCE(AmphurName, '') [AmphurName],
		COALESCE(ProvinceName, '') [ProvinceName],
		COALESCE(PostCode, '') [PostCode],
		COALESCE(BusinessFirstDate, '') [BusinessFirstDate],
		COALESCE(BranchTypeCode, '') [BranchTypeCode],
		COALESCE(BranchTypeName, '') [BranchTypeName],
		COALESCE(RegisteredCapital, '') [RegisteredCapital],
		COALESCE(Status, '') [Status]
	FROM taxpayer`

	tsql += " where NID = '" + nid + "';"

	fmt.Printf("tsql!" + tsql + "\n")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	isnil := true
	datataxpayer := new(models.Taxpayer)

	for rows.Next() {
		isnil = false
		if err := rows.Scan(
			&datataxpayer.NID,
			&datataxpayer.BranchNumber,
			&datataxpayer.BranchTitle,
			&datataxpayer.BranchName,
			&datataxpayer.BuildingName,
			&datataxpayer.RoomNumber,
			&datataxpayer.FloorNumber,
			&datataxpayer.VillageName,
			&datataxpayer.HouseNumber,
			&datataxpayer.MooNumber,
			&datataxpayer.SoiName,
			&datataxpayer.StreetName,
			&datataxpayer.ThumbolName,
			&datataxpayer.AmphurName,
			&datataxpayer.ProvinceName,
			&datataxpayer.PostCode,
			&datataxpayer.BusinessFirstDate,
			&datataxpayer.BranchTypeCode,
			&datataxpayer.BranchTypeName,
			&datataxpayer.RegisteredCapital,
			&datataxpayer.Status); err != nil {
			fmt.Println(err)
		}
	}

	if isnil {
		datataxpayer = nil
	}

	return tax.RenderJSON(datataxpayer)
}
