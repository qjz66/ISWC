package dao

import "Server/model"

// InsertIntoInfo 将检测信息插入数据库
func (d *Dao) InsertIntoInfo(info model.Info, topic string) error {
	var userTopic model.UserTopic
	err := d.Model(&model.Info{}).Create(&info).Error
	if err != nil {
		return err
	}
	err = d.Model(&model.UserTopic{}).Where("id=? and topic=?", info.ID, topic).First(&userTopic).Error
	if len(userTopic.Topic) == 0 {
		userTopic.Topic = topic
		userTopic.ID = info.ID
		userTopic.Num = 1
		err = d.Model(&model.UserTopic{}).Create(&userTopic).Error
		if err != nil {
			return err
		}
	} else {
		//将user_topics表更新
		err = d.Model(&userTopic).Update("num", userTopic.Num+1).Error
	}
	return err
}
