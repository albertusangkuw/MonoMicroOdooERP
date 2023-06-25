package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

type DBConfig struct {
	Host string
	DB string
	Username string
	Password string
}

func (d *DBConfig) toURLPath() string{
	//"root:2120@tcp(172.18.0.54:3306)/pos_db?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		 						d.Username,d.Password , d.Host, d.DB)

}

var CurrentDBConfig DBConfig

func connect() bool {
	var err error
	dbInstance, err = gorm.Open(mysql.Open(CurrentDBConfig.toURLPath()), &gorm.Config{})
	return err == nil
}

func Database() *gorm.DB {
	if dbInstance == nil {
		if !connect() {
			fmt.Print("Failed to connect DB")
			return nil
		}
	}
	return dbInstance
}
