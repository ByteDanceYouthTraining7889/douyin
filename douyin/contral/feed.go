package contral

import (
	"douyin/global"
	"douyin/service"
	"log"

	"github.com/gin-gonic/gin"
)

var vedionowid = 1

func GetVideoFeed(c *gin.Context) {

	tags, err := service.GetVideoList(global.DbEngine)
	if err != nil {
		log.Println("viode list  err：", err)
		return
	}
	log.Println(tags)
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "sucess",
		"next_time":   nil,  //后面改进
		"video_list":  tags, //*************
	})
	return
}
