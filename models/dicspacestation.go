package models

import (
	"strings"
	orm "yangjian.com/basicarchitecture/dao"
)

//DicSpaceStation is the dic of space station
type DicSpaceStation struct {
	ShortName string  `gorm:"PRIMARY_KEY;type:varchar(20);" json:"short_name,omitempty"`
	CName     string  `gorm:"type:varchar(100);" json:"c_name,omitempty"`
	Py        string  `json:"py,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

//TableName
func (dsStation *DicSpaceStation) TableName() string {

	return "dic_space_station"
}

//CreateTableIfNotExists
func (dsStation *DicSpaceStation) CreateTableIfNotExists() (err error) {

	if !orm.DB().HasTable(&dsStation) {
		return orm.DB().CreateTable(&dsStation).Error
	}
	return nil
}

//Insert
func (dsStation *DicSpaceStation) Insert() (err error) {

	return orm.DB().Create(&dsStation).Error
}

//Inserts
func (dsStation *DicSpaceStation) Inserts(dsStations []*DicSpaceStation) (err error) {
	tx := orm.DB().Begin()
	if err = tx.Error; err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for item := range dsStations {
		if err = tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (dsStation *DicSpaceStation) Destroy() (err error) {

	if dsStation.ShortName == "" {
		return nil
	}
	return orm.DB().Delete(&dsStation).Error
}

func (dsStation *DicSpaceStation) Update() (err error) {

	return orm.DB().Model(&DicSpaceStation{}).Updates(&dsStation).Error
}

func (dsStation *DicSpaceStation) Select() (resultStation *DicSpaceStation, err error) {

	resultStation = new(DicSpaceStation)
	err = orm.DB().Where(&dsStation).First(&resultStation).Error
	return
}

func (dsStation *DicSpaceStation) List(offset, limit int) (dsStations []*DicSpaceStation, err error) {

	tx := orm.DB().Select([]string{"short_name", "c_name", "latitude", "longitude"})
	err = tx.Offset(offset).Limit(limit).Find(&dsStations).Error
	return
}

func (dsStation *DicSpaceStation) Search(keyWord string, offset, limit int, isLeft bool) (dsStations []*DicSpaceStation, err error) {

	tx := orm.DB().Select([]string{"short_name", "c_name", "latitude", "longitude"})
	keyWordLow := strings.ToLower(keyWord)
	if keyWord != "" {
		if isLeft {
			tx = tx.Where("short_name LIKE ? OR c_name LIKE ? OR py LIKE ?", keyWordLow+"%", keyWord+"%", keyWordLow+"%")
		} else {
			tx = tx.Where("short_name LIKE ? OR c_name LIKE ? OR py LIKE ?", "%"+keyWordLow+"%", "%"+keyWord+"%", "%"+keyWordLow+"%")
		}
	}
	err = tx.Offset(offset).Limit(limit).Find(&dsStations).Error
	return
}
