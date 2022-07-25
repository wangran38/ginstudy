package controllers

import (
	"ginstudy/models"

	// "ginstudy/utils"
	"github.com/gin-gonic/gin"
)

type Rulesserch struct {
	Title     string `json:"title"`
	Pathname  string `json:"pathname"`
	Component string `json:"component"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Order     string `json:"order"`
}

//获取当前用户信息
func Getruleslist(c *gin.Context) {
	//从header中获取到token
	var searchdata Rulesserch
	c.BindJSON(&searchdata)
	// //读取数据库
	limit := searchdata.Limit
	page := searchdata.Page
	order := searchdata.Order
	result := make(map[string]interface{})
	// name:=""
	// fmt.Println(username)
	search := &models.Authrule{
		Title:     searchdata.Title,
		Pathname:  searchdata.Pathname,
		Component: searchdata.Component,
	}
	// fmt.Println(search.Title)
	listdata := models.GetRulesList(limit, page, search, order)
	listnum := models.GetRulestotal(search)

	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取菜单失败1",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功1",
			"data":    result,
		})
		return
	}
}

//获取当前用户 所有的权限控制
