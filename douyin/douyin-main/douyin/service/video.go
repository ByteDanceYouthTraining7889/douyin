package service

import (
	"douyin/model"
	"log"

	"github.com/jinzhu/gorm"
)

func GetVideo(v model.Video, db *gorm.DB) (model.Video, error) {
	video, err := v.Get(db)
	if err != nil {
		log.Println(err)
		return video, err
	}
	return video, nil
}

func CreateVideo(v model.Video, db *gorm.DB) error {
	err := v.Create(db)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func GetVideoList(db *gorm.DB) (*[30]model.Videofeed, error) {
	list, err := model.VideoList(db)
	if err != nil {
		log.Println(err)
	}
	return list, nil
}

// func UpdateVideo(v model.Video) error {
// 	err := v.Update()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return nil
// }
