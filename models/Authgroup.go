package models

import (
	"errors"
	"time"
)

type Authgroup struct {
	Id      int64     `json:"id"`
	Pid     int       `json:"pid"`
	Name    string    `json:"name" xorm:"LONGTEXT "`
	Rules   string    `json:"rules" xorm:"LONGTEXT "`
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
	Status  int       `json:"status"`
}

func (a *Authgroup) TableName() string {
	return "auth_group"
}

//根据用户id找用户返回数据
func SelectGidRule(Id int64) (*Authgroup, error) {
	a := new(Authgroup)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("组别菜单数据出错！")
	}
	return a, nil

}

//添加
func Addgroup(a *Authgroup) error {
	_, err := orm.Insert(a)
	return err
}
func GetgroupList(limit int, pagesize int, search string, order string) []*Authgroup {
	//fmt.Println("搜索关键词",search)
	//    limit,_ := strconv.Atoi(limits)
	//
	//   if pagesize-1<1 {
	page := pagesize - 1
	//   }
	listdata := []*Authgroup{}
	//拼接搜索分页查询语句
	var byorder string
	byorder = "id ASC"
	if order == "-id" {
		byorder = "id DESC"
	}
	orm.Table("auth_group").
		Where("name like ?", "%"+search+"%").
		OrderBy(byorder).
		// Orderby(byorder).
		Limit(limit, limit*page).
		Find(&listdata)
	//  orm.Where("username like ?", "%"+search+"%").Limit(limit*pagesize, pagesize).Find(&listadmin)
	//    fmt.Println(listadmin)
	return listdata
}

func Getgrouptotal(search string) int64 {
	var num int64
	num = 0
	a := new(Authgroup)
	total, err := orm.Cols("id", "name").Where("name like ?", "%"+search+"%").Count(a)
	if err == nil {
		num = total

	}
	return num
}
