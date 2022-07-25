package controllers

import (
	"ginstudy/lib"
	"ginstudy/models"
	"net/http"
	"time"
	// "fmt"
	// "strconv"
	"github.com/gin-gonic/gin"
)

type Adminform struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Adminserch struct {
	Username string `json:"title"`
	Limit int `json:"limit"`
	Page int `json:"page"`
	Order string `json:"sort"`
}
// type AdminController struct {
// 	BaseController //继承统一判断是否登录或者是否有权限类
// 	// beego.Controller
// }
func AddAdmin(c *gin.Context) {
	var admindata Adminform
	if err := c.ShouldBind(&admindata); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"msg":     "表单未完整！",
			"msgcode": "-1",
			// "data":    reqIP,
		})
		return
	}
	// //读取数据库
	Admin := new(models.Admin)
	Admin.Username = admindata.Username
	info, _ := models.SelectUserByUserName(Admin.Username) //判断账号是否存在！
	if info != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "该用户已经存在！",
		})
		return
	}
	pwd, salt := lib.Password(4, admindata.Password) //截取四位随机盐+上这个做原始密码
	Admin.Password = pwd
	Admin.Salt = salt
	Admin.Created = time.Now()
	err := models.AddAdmin(Admin) //判断账号是否存在！
	if err != nil {
		c.JSON(200, gin.H{
			"code": "201",
			"msg":  "添加数据出错！",
			"data": err,
		})
		return
	} else {
		result := make(map[string]interface{})
		result["id"] = Admin.Id //返回当前总数
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "数据添加成功！",
			"data": result,
		})

	}
}
func GetAdminlist(c *gin.Context) {
	var searchdata Adminserch
	c.BindJSON(&searchdata)
	// //读取数据库
	// Admin := new(models.Admin)
	// var limit,pagesize int 
	// var username string
	result := make(map[string]interface{})

	// fmt.Println(username)
	// username:= c.DefaultPostForm("title", "")
	// limit1 := c.DefaultPostForm("limit", "1")
	// fmt.Println(limit1)
	// limit,_=strconv.Atoi(limit1)
	limit:= searchdata.Limit
	page:= searchdata.Page
	username:= searchdata.Username
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
		listdata := models.GetUserList(limit,page,username,order)
		listnum := models.GetUsertotal(username)
		
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