package dao

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	dbError  error
	once     sync.Once
	instance *gorm.DB
)

type MySQLHelper struct {
	Host        string
	Port        int
	DBName      string
	UsrName     string
	UsrPassword string
}

func (msHelper *MySQLHelper) CheckConfig() bool {

	result := false
	if strings.TrimSpace(msHelper.Host) != "" &&
		strings.TrimSpace(msHelper.DBName) != "" &&
		strings.TrimSpace(msHelper.UsrName) != "" &&
		strings.TrimSpace(msHelper.UsrPassword) != "" &&
		msHelper.Port != 0 {
		result = true
	}
	return result
}
func Init() {
	once.Do(func() {
		msHelper := &MySQLHelper{
			Host:        viper.GetString("database.mysql_host"),
			Port:        viper.GetInt("database.mysql_port"),
			DBName:      viper.GetString("database.mysql_dbname"),
			UsrName:     viper.GetString("database.mysql_user"),
			UsrPassword: viper.GetString("database.mysql_password"),
		}
		if !msHelper.CheckConfig() {
			dbError = errors.New("config is empty")
			return
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=50ms", msHelper.UsrName, msHelper.UsrPassword, msHelper.Host, msHelper.Port, msHelper.DBName)
		instance, dbError = gorm.Open("mysql", connStr)
	})
}

func DB() *gorm.DB {
	return instance
}

func Error() error {
	return dbError
}
