package data

import (
	"database/sql"
	"log"
	"strconv"

	app "github.com/lwalter/lessonshare-api/src/app"
)

var db *sql.DB

func InitDbConn() {
	connStr := "postgres://" + app.Config.Db.User + ":" + app.Config.Db.Pass + "@" + app.Config.Db.Host + ":" + strconv.Itoa(app.Config.Db.Port) + "/" + app.Config.Db.Name + "?sslmode="

	if app.Config.Db.SslMode {
		connStr += "enable"
	} else {
		connStr += "disable"
	}

	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil || db == nil {
		log.Fatal("Unable to connect to db")
		panic("Unable to connect to db")
	}
}
