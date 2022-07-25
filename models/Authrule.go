package models

import (
	"fmt"
	"strings"
	"time"
	// "reflect"
)

type Authrule struct {
	Id         int64
	Pid        int64
	Type       string `json:"type"`
	Icon       string `json:"icon"`
	Pathname   string `json:"pathname"`
	Title      string `json:"title"`
	Remark     string
	Ismenu     int       `json:"ismenu" xorm:"not null default 1 comment('是否启用 默认1 菜单 0 文件') TINYINT"`
	Created    time.Time `json:"createtime" xorm:"created int"`
	Updated    time.Time `json:"updatetime" xorm:"updated int"`
	Deletetime int       `json:"deletetime"`
	Weigh      int
	Status     string `json:"status" xorm:"varchar(40)"`
	Component  string `json:"component"`
	// Children   []*Authrule
}
type Treerule struct {
	Id        int64
	Pid       int64
	Type      string `json:"type"`
	Icon      string `json:"icon"`
	Pathname  string `json:"pathname"`
	Component string `json:"component"`
	Title     string `json:"title"`
	Remark    string
	Ismenu    int `json:"ismenu"`
	Weigh     int
	Status    string `json:"status"`
	Children  []*Treerule
}

func (a *Authrule) TableName() string {
	return "auth_rule"
}

//获取树状数据
func Getruletree() []*Treerule {
	m := new(Authrule)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treelist(0)

}

//全部菜单
func (m *Authrule) Treelist(pid int64) []*Treerule {
	// menus := new(Authrule)
	// 	var a []Authrule
	var menus []Authrule
	orm.Where("pid = ?", pid).Where("deletetime is null").Find(&menus)
	treelist := []*Treerule{}
	for _, v := range menus {
		child := v.Treelist(v.Id)
		node := &Treerule{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Icon:      v.Icon,
			Pathname:  v.Pathname,
			Component: v.Component,
			Title:     v.Title,
			Remark:    v.Remark,
			Ismenu:    v.Ismenu,
			Weigh:     v.Weigh,
			Status:    v.Status,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}

//获取用户树状数据
func Getruleadmintree(Rules string) []*Treerule {
	m := new(Authrule)
	//不new一个新的，采用结构体，外部无法访问()getruletreee() []*tt这样子，只能new一个，然后去访问
	return m.Treelistgroup(0, Rules)

}

//传参得某个权限的菜单集合
func (m *Authrule) Treelistgroup(pid int64, Rules string) []*Treerule {
	fmt.Println(Rules)
	//    a := new(Authrule)
	// // 	var a []Authrule
	var menus []*Authrule
	ids := strings.Split(Rules, ",") //转成数组用orm in
	// 	// ids:= string.Join(Rules,",")
	// Where("pid = ?", v.Id)
	orm.Where("pid = ?", pid).Where("status = ?", "normal").In("id", ids).Find(&menus)
	treelist := []*Treerule{}
	for _, v := range menus {
		child := v.Treelistgroup(v.Id, Rules)
		node := &Treerule{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Icon:      v.Icon,
			Pathname:  v.Pathname,
			Component: v.Component,
			Title:     v.Title,
			Remark:    v.Remark,
			Ismenu:    v.Ismenu,
			Weigh:     v.Weigh,
			Status:    v.Status,
		}
		node.Children = child
		treelist = append(treelist, node)
	}
	return treelist

}

// func (menus []*Authrule)tree(pid int64) []*Authrule {
// 	var nodes []*Authrule
// 	if reflect.ValueOf(menus).IsValid() {
// 	for k,v:=range menus {
// 	orm.Where("pid = ?", v.Id).Find(&nodes)
// 	    //判断是否存在数据,存在进行树状图重构

// 		for kk,_:=range nodes{
// 			menus[k].Children=append(menus[k].Children,nodes[kk])
// 			tree(nodes)
// 		}

// 		//
// 	}

// 	}
// 	return menus
// }
// func SelectAllRules() ([]Authrule) {
// 	// a := new(Authrule)
// 	var a []Authrule
// 	// a := make([]*Authrule,0)
// 	has, err := orm.Find(&a)
// 	if err != nil {
// 		return nil
// 	}
// 	if !has {
// 		return nil
// 	}
// 	return a

// }
// //根据用户名密码查询用户
// func SelectAllRules(Id int64) []Authrule {
// 	// a := new(Authrule)
// 	everyone := make([]Authrule, 0)
// 	orm.Where("Pid = ?", Id).Find(&everyone)
// 	for _,v:=range everyone {
// 				// var cc make([]rule,0)
// 				everyone.Children=orm.Where("Pid = ?", v.Id).Find(&everyone)
// 				// fmt.Println([]*rule.Children)
// 	}
// 	// if err != nil {
// 	// 	return nil, errors.New("没有数据！")
// 	// }

// 	return everyone

// }

//列表查询
// func RulesPageList(params *NoticeParam) ([]*Authrule, int64) {
// 	query := orm.NewOrm().QueryTable(NoticeTBName())
// 	data := make([]*Authrule, 0)
// 	//默认排序
// 	sortorder := "-id" //定义索引
// 	if len(params.Search.Value) > 0 {
// 		query = query.Filter("notice_id", params.Search.Value)
// 	}

// 	total, _ := query.Count()
// 	query.OrderBy(sortorder).Limit(params.Length, params.Start).All(&data)
// 	return data, total
// }
