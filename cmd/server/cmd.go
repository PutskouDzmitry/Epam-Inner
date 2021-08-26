package main

import (
	"github.com/PutskouDzmitry/Epam-Inner/pkg/api"
	"github.com/PutskouDzmitry/Epam-Inner/pkg/data"
	"github.com/PutskouDzmitry/Epam-Inner/pkg/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	host       = ""
	check      = os.Getenv("CHECK")
	port       = os.Getenv("DB_USERS_PORT")
	user       = os.Getenv("DB_USERS_USER")
	dbname     = os.Getenv("DB_USERS_DBNAME")
	password   = os.Getenv("DB_USERS_PASSWORD")
	sslmode    = os.Getenv("DB_USERS_SSL")
	portServer = os.Getenv("SERVER_OUT_PORT")
)

func init() {
	if check == "postgres" {
		host = "db"
	}
	if check == "timescale" {
		host = "dbup"
	}
	if check == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}
	if password == "" {
		password = "1234"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
	if portServer == "" {
		portServer = "8081"
	}
}

func main() {
	logrus.Info("App works with ", host)
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		logrus.Fatal(err)
	}
	entityData := data.NewEntityData(conn)
	router := gin.Default()
	api.ServeUserResource(*entityData, router)
	router.Run(":" + portServer)
}
