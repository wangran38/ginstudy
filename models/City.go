package models

//城市后端模型
import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id        int64
	Pid       int
	Shortname string `xorm:"varchar(200)" json:"shortname"`
	Name      string `xorm:"varchar(200)" json:"name"`
	Mergename string `xorm:"varchar(200)" json:"mergename"`
	Level     int    `json:"status" xorm:"not null default 1 comment('层级 0 1 2 省市区县') TINYINT"`
	Pinyin    int    `xorm:"varchar(200)" json:"pingyin"`
	Code      string `xorm:"varchar(200)" json:"code"`
	Zip       string `xorm:"varchar(200)" json:"zip"`
	First     string `xorm:"varchar(200)" json:"first"`
	Lng       string `xorm:"varchar(200)" json:"lng"`
	Lat       string `xorm:"varchar(200)" json:"lat"`
}

func (a *City) TableName() string {
	return "area"
}

//根据用户名密码查询用户
func SelectBycityid(Id int) (*City, error) {
	a := new(City)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("暂无此条数据！")
	}
	return a, nil

}

//add
