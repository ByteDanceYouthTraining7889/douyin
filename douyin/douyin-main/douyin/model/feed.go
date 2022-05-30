package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Video struct {
	Authorid      int64  `json:"authorid"`       // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}
type Videofeed struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

func NewVideo() Video {
	return Video{}
}

func (v Video) TableName() string {
	return "video"
}
func (v Video) Get(db *gorm.DB) (Video, error) {
	db = db.Where(
		"id=?",
		v.ID,
	)
	err := db.First(&v).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("video get err:", err)
		return v, err
	}
	return v, nil

}
func VideoList(db *gorm.DB) (*[30]Videofeed, error) {
	var videolist []Video
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Video{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	// db = db.Where("state=?", t.)
	if err = db.Find(&videolist).Error; err != nil {
		log.Println("find videolist err:", err)
		return nil, err
	}
	var Feedlist [30]Videofeed
	for k, v := range videolist {
		u := User{ID: v.Authorid}
		fmt.Println("the k:", k)
		Feedlist[k] = Videofeed{Author: u, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title}
	}
	return &Feedlist, nil
}
func PublishList(db *gorm.DB) (*[30]Videofeed, error) {
	var publishlist []Video
	var err error

	db = db.Limit(30) //==============后面前

	db = db.AutoMigrate(&Video{}) //自己加的 自动绑定结构体
	// fmt.Println("gua")
	// fmt.Println(t.Tablename()) //自己加的 测试是否自动运行了 知道是哪个表
	// db = db.Where("state=?", t.)
	if err = db.Find(&publishlist).Error; err != nil {
		log.Println("find publishlist err:", err)
		return nil, err
	}
	var Feedlist [30]Videofeed
	for k, v := range publishlist {
		u := User{ID: v.Authorid}
		fmt.Println("the k:", k)
		Feedlist[k] = Videofeed{Author: u, CommentCount: v.CommentCount, FavoriteCount: v.FavoriteCount, ID: v.ID, IsFavorite: v.IsFavorite, CoverURL: v.CoverURL, PlayURL: v.PlayURL, Title: v.Title}
	}
	return &Feedlist, nil
}
func (v Video) Create(db *gorm.DB) error {
	return db.Debug().Create(&v).Error
}

// func (v Video) Update() error {
// 	db = db.Model(&User{}).Where("=?", u.Username) //model 指定db操作的实力模型 一般在没有find 等指定类型时使用 操作限制条件是传进来的tag的数据行
// 	//fmt.Println(t.Name)
// 	err := db.Select("*").Save(u).Error //用save不报错 create报错 只能先将就save 用了 配合selct 也解决了 0值不能改的问题
// 	if err != nil {
// 		fmt.Println("user 更新错误：", err)
// 		return err
// 	}
// 	return nil
// }
