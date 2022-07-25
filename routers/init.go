package routers

import (
	"fmt"
	"ginstudy/controllers"
	_ "ginstudy/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {

	router := gin.Default()
	router.Use(Cors())
	
	// 版本v1
	admin := router.Group("/admin")
	{
		admin.POST("/login", controllers.LoginController) //登录
		admin.POST("/logout", controllers.Loginout) //登录
		admin.POST("/add", controllers.AddAdmin)
		admin.POST("/getinfo", controllers.GetAdminInfo)
		admin.POST("/getrule", controllers.GetAdminRule)
		admin.POST("/getadminlist", controllers.GetAdminlist)
		admin.POST("/getgrouplist", controllers.Getgrouplist)
	}

	// router.POST("/admin/add", controllers.AdminController) //登录
	router.Run(":8081")
}

////// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json,multipart/form-data,application/x-www-form-urlencoded")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		
		//放行所有OPTIONS方法,防止vue握手2次的问题
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// if method == "OPTIONS" {
		// 	c.JSON(http.StatusOK, "Options Request!")
		// }
		// 处理请求
		c.Next() //  处理请求
	}
}
