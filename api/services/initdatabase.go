package services

import (
	"github.com/pkg/errors"
	"yangjian.com/basicarchitecture/dao"
)

func InitDatabase() (err error) {

	dao.Init()
	if errDB := dao.Error(); errDB != nil {
		err = errors.Wrap(errDB, "database error")
		return
	}
	return nil
}
