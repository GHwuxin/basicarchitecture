package controllers

import (
	log "github.com/sirupsen/logrus"
	"yangjian.com/basicarchitecture/api/services"
)

func InitController() {
	err := services.InitDatabase()

	if err != nil {
		log.Errorf("init database error:%s", err.Error())
		return
	}
}
