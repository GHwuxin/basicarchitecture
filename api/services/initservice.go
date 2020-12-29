package services

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	orm "yangjian.com/basicarchitecture/dao"
	"yangjian.com/basicarchitecture/models"
)

func InitServices() (err error) {

	if err := InitDataBase(); err != nil {
		return err
	}
	return nil
}

func InitDataBase() (err error) {

	orm.Init()
	if errDB := orm.Error(); errDB != nil {
		err = errors.Wrap(errDB, "database error")
		return
	}
	// init dic space station table
	if err = (&models.DicSpaceStation{}).CreateTableIfNotExists(); err != nil {
		return fmt.Errorf("%s table init error:%s", (&models.DicSpaceStation{}).TableName(), err.Error())
	}
	mode := viper.GetString("server.mode")
	if mode == "debug" {
		orm.DB().LogMode(true)
	}
	orm.DB().SetLogger(log.StandardLogger())
	return nil
}
