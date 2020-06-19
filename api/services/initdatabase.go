package services

import (
	"github.com/pkg/errors"
	orm "yangjian.com/basicarchitecture/dao"
)

func InitDatabase() (err error) {

	orm.Init()
	if errDB := orm.Error(); errDB != nil {
		err = errors.Wrap(errDB, "database error")
		return
	}
	return nil
}
