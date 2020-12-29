package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"yangjian.com/basicarchitecture/api/services"
	"yangjian.com/basicarchitecture/config"
	"yangjian.com/basicarchitecture/storage/logs"
	"yangjian.com/basicarchitecture/utils"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const bannerPath = "./config/banner.txt"

func InitController() (err error) {
	// check logger
	if e := logs.Error(); e != nil {
		err = fmt.Errorf("init logger error:%s\n", e.Error())
		return err
	}
	log.Infoln("log init success")
	// check config
	if e := config.Error(); e != nil {
		log.Errorf("config is error:%s", e.Error())
		err = errors.Wrap(err, "config is err")
		return err
	}
	log.Infoln("config init success")
	// check database
	if e := services.InitServices(); e != nil {
		log.Errorf("init database error:%s", e.Error())
		err = errors.Wrap(err, "init controller err")
		return err
	}
	log.Infoln("database init success")
	printBanner()
	return nil
}

func printBanner() {

	if utils.Exists(bannerPath) {
		file, err := os.Open(bannerPath)
		if err == nil {
			reader := bufio.NewReader(file)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				fmt.Println(strings.Replace(line, "\r\n", "", -1))
			}
		}
	}
}
