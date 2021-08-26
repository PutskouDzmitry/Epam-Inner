package db

import (
	"fmt"
	"time"

	"github.com/cenkalti/backoff"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//GetConnection it's return a new connection in db
func GetConnection(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	back := config()
	var connection *gorm.DB
	var err error
	for {
		timeWait := back.NextBackOff()
		if timeWait.Seconds() >= 20  {
			return nil, fmt.Errorf("error with connection to db")
		}
		connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			//logrus.Debug("we wait connect to postgres, time: ", timeWait)
			//logrus.Error(err)
		} else {
			break
		}
	}
	return connection, nil
}

func config() *backoff.ExponentialBackOff {
	back := backoff.NewExponentialBackOff()
	back.MaxInterval = 20 * time.Second
	back.Multiplier = 2
	return back
}