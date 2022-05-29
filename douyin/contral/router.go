package contral

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./public") //将url上对应的static 改为public
	r.POST("/douyin/user/register/", CreatUserInDenglu)
	r.GET("/douyin/feed", GetVideoFeed)
	r.POST("/douyin/user/login/", UserInlogin)
	r.GET("/douyin/user/", GetUserMsg)
	r.POST("/douyin/publish/action/", FabuVideo)
	r.GET("/douyin/publish/list/", PublishList)
	return r
}
