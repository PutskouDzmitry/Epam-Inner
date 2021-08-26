package main

import (
	"github.com/PutskouDzmitry/JWT/pkg/api"
	"github.com/PutskouDzmitry/JWT/pkg/data"
	"github.com/PutskouDzmitry/JWT/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	hostUp     = os.Getenv("DBUP_USERS_HOST")
	host       = os.Getenv("DB_USERS_HOST_POSTGRES")
	port       = os.Getenv("DB_USERS_PORT")
	user       = os.Getenv("DB_USERS_USER")
	dbname     = os.Getenv("DB_USERS_DBNAME")
	password   = os.Getenv("DB_USERS_PASSWORD")
	sslmode    = os.Getenv("DB_USERS_SSL")
	portServer = os.Getenv("SERVER_OUT_PORT")
)

func init() {
	if host == "" {
		host = "db"
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
	time.Sleep(20 * time.Second)
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		logrus.Fatal(err)
	}
	entityData := data.NewEntityData(conn)
	router := gin.Default()
	api.ServeUserResource(*entityData, router)
	router.Run(":8081")
}
