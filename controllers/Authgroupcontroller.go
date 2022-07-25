package controllers

import (
	// "fmt"
	"ginstudy/models"
    // "ginstudy/utils"
	"github.com/gin-gonic/gin"
)
type Groupserch struct {
	Name string `json:"name"`
	Limit int `json:"limit"`
	Page int `json:"page"`
	Order string `json:"sort"`
}
//获取当前用户信息
func Getgrouplist(c *gin.Context) {
	//从header中获取到token
	var searchdata Groupserch
	c.BindJSON(&searchdata)
	// //读取数据库
	// Admin := new(models.Admin)
	// var limit,pagesize int 
	// var username string
	result := make(map[string]interface{})
    // name:=""
	// fmt.Println(username)
	// username:= c.DefaultPostForm("title", "")
	// limit1 := c.DefaultPostForm("limit", "1")
	// fmt.Println(limit1)
	// limit,_=strconv.Atoi(limit1)
	limit:= searchdata.Limit
	page:= searchdata.Page
	name:= searchdata.Name
	order:= searchdata.Order
	// pagesize1 := c.DefaultPostForm("page", "1")
	// pagesize,_= strconv.Atoi(pagesize1)
	// page:=pagesize-1
	// limit,_=strconv.Atoi(limit1)
	// if c.PostForm("limit")!="" { 
	// 	limit1:=c.PostForm("limit")
	// 	limit,_=strconv.Atoi(limit1)
	// 	if limit<1 {
	// 		limit=10
	// 	}
	// 	} else {
	// 		limit=10 
	// 		}
	// if c.PostForm("pagenum")!="" {
	// 	 pagesize1:=c.PostForm("pagenum") 
	// 	 pagesize,_= strconv.Atoi(pagesize1)
	// 	 } else {
	// 		 pagesize=0 
	// 	}
	// 	//判断是否是超级用户

		// limit=0
		// pagesize=10
		listdata := models.GetgroupList(limit,page,name,order)
		listnum := models.Getgrouptotal(name)
		
		result["page"] = page
		result["totalnum"] = listnum
		result["limit"] = limit
		if listdata== nil {
			c.JSON(200, gin.H{
			"code":    201,
			"message": "获取菜单失败1",
			"data": "",
		})
		return
			} else {
result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功1",
			"data": result,
		})
		return
			}
}

//获取当前用户 所有的权限控制
