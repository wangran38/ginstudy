package models
import (
	"errors"
)
type Authaccess struct {
	Uid int64
	Gid int64 `json:"group_id"`
}

func (a *Authaccess) TableName() string {
	return "auth_group_access"
}
//根据用户id找用户返回数据
func SelectAdminGid(Id int64) (*Authaccess, error) {
	a := new(Authaccess)
	has, err := orm.Where("uid = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到所在组别！")
	}
	return a, nil

}