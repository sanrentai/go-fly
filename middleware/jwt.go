package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
	"time"
)
func JwtPageMiddleware(c *gin.Context){
	//暂时不处理
	//token := c.Query("token")
	//userinfo := tools.ParseToken(token)
	//if userinfo == nil {
	//	c.Redirect(302,"/login")
	//	c.Abort()
	//}
}
func JwtApiMiddleware(c *gin.Context){
	token := c.GetHeader("token")
	userinfo := tools.ParseToken(token)
	if userinfo == nil||userinfo["name"]==nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "验证失败",
		})
		c.Abort()
	}
	createTime:=int64(userinfo["create_time"].(float64))
	var expire int64=2*60*60
	nowTime:=time.Now().Unix();
	if (nowTime-createTime) >=expire{
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "token失效",
		})
		c.Abort()
	}
	c.Set("user",userinfo["name"])
}
