package controllers

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "localhost"
var port = 1450
var user = "systemweb"
var password = "1q2w3e4r"
var database = "app_gmapmaker"

var db *sql.DB
