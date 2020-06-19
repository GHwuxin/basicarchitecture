package controllers

import (
	"fmt"

	"yangjian.com/basicarchitecture/api/services"
	"yangjian.com/basicarchitecture/config"
	"yangjian.com/basicarchitecture/storage/logs"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func InitController() (err error) {
	// check logger
	if e := logs.Error(); e != nil {
		err = fmt.Errorf("init logger error:%s\n", e.Error())
		return err
	}
	// check config
	if e := config.Error(); e != nil {
		log.Errorf("config is error:%s", e.Error())
		err = errors.Wrap(err, "config is err")
		return err
	}
	// check database
	if e := services.InitDatabase(); e != nil {
		log.Errorf("init database error:%s", e.Error())
		err = errors.Wrap(err, "init controller err")
		return err
	}
	return nil
}
