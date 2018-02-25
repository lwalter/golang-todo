package data

import (
	"database/sql"
	"log"
	"strconv"

	c "github.com/lwalter/lessonshare-api/src/config"
)

var db *sql.DB

func InitDbConn() {
	connStr := "postgres://" + c.Config.Db.User + ":" + c.Config.Db.Pass + "@" + c.Config.Db.Host + ":" + strconv.Itoa(c.Config.Db.Port) + "/" + c.Config.Db.Name + "?sslmode="

	if c.Config.Db.SslMode {
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
