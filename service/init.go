package service

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBConn() (*gorm.DB, error) {
	DBMS := "mysql"
	HOSTNAME := os.Getenv("RDS_HOSTNAME")
	USERNAME := os.Getenv("RDS_USERNAME")
	DBNAME := os.Getenv("RDS_DB_NAME")
	PASSWORD := os.Getenv("RDS_PASSWORD")
	PORT := os.Getenv("RDS_PORT")

	CONNECT := USERNAME + ":" + PASSWORD + "@(" + HOSTNAME + ":" + PORT + ")/" + DBNAME + "?parseTime=true" // ?parseTime=trueがないとdatetime型がtime.Timeに変換することができない
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil, err
	}

	return db, nil
}
