package services

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
	"yangjian.com/basicarchitecture/models"
)

func AddStationDic(shortName, cName string, latitude, longitude float64) (err error) {

	py := ""
	if cName != "" {
		arg := pinyin.NewArgs()
		arg.Style = pinyin.FIRST_LETTER
		pyHeader := strings.Join(pinyin.LazyPinyin(cName, arg), "")
		arg.Style = pinyin.NORMAL
		pyAll := strings.Join(pinyin.LazyPinyin(cName, arg), "")
		py = strings.Join([]string{pyHeader, pyAll}, "_")
	}
	dsStation := &models.DicSpaceStation{
		ShortName: shortName,
		CName:     cName,
		Py:        py,
		Latitude:  latitude,
		Longitude: longitude,
	}
	return dsStation.Insert()
}

func DestroyStationDic(shortName string) (err error) {

	dsStation := &models.DicSpaceStation{ShortName: shortName}
	return dsStation.Destroy()
}

func UpdateStationDic(shortName, cName string, latitude, longitude float64) (err error) {

	py := ""
	if cName != "" {
		arg := pinyin.NewArgs()
		arg.Style = pinyin.FIRST_LETTER
		pyHeader := strings.Join(pinyin.LazyPinyin(cName, arg), "")
		arg.Style = pinyin.NORMAL
		pyAll := strings.Join(pinyin.LazyPinyin(cName, arg), "")
		py = strings.Join([]string{pyHeader, pyAll}, "_")
	}
	dsStation := &models.DicSpaceStation{
		ShortName: shortName,
		CName:     cName,
		Py:        py,
		Latitude:  latitude,
		Longitude: longitude,
	}
	return dsStation.Update()
}

func SelectStationDic(shortName string) (data interface{}, err error) {

	dsStation := &models.DicSpaceStation{ShortName: shortName}
	return dsStation.Select()
}

func ListStationDic(page, limit int) (data interface{}, err error) {

	offset := (page - 1) * limit
	return (&models.DicSpaceStation{}).List(offset, limit)
}

func SearchStationDic(keyWord string, page, limit int) (data interface{}, err error) {

	offset := (page - 1) * limit
	resDic := make([]*models.DicSpaceStation, 0)
	resDic, err = (&models.DicSpaceStation{}).Search(keyWord, offset, limit, true)
	if err != nil || resDic == nil || len(resDic) == 0 {
		resDic, err = (&models.DicSpaceStation{}).Search(keyWord, offset, limit, false)
	}
	if err != nil {
		return nil, err
	}
	return resDic, err
}

