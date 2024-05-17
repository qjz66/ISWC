package dao

import (
	"Server/model"
	"time"
)

// Favorite 给动态点赞
func (d *Dao) Favorite(authorId, updateId int64) (int64, error) {
	update := &model.Update{ID: authorId}
	err := d.Model(&model.Update{}).Where("update_id=?", updateId).First(&update).Error
	if err != nil {
		return 0, err
	}
	favorite := *update.Favorite + 1
	err = d.Model(&update).Update("favorite", favorite).Error
	if err != nil {
		return 0, err
	}
	return favorite, nil
}

// Update 上传动态
func (d *Dao) Update(content string, id int64) (int64, error) {
	favorite := new(int64)
	*favorite = 0
	update := &model.Update{ID: id, Content: content, Favorite: favorite, Time: time.Now()}
	err := d.Model(&model.Update{}).Create(&update).Error
	if err != nil {
		return 0, err
	}
	return update.UpdateID, nil
}

// Comment 给动态评论
func (d *Dao) Comment(comment model.Comment) error {
	err := d.Model(&model.Comment{}).Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

// GetComments 得到评论列表
func (d *Dao) GetComments(updateId int64, comments *[]model.Comment) error {
	err := d.Model(&model.Comment{}).Where("update_id=?", updateId).Find(comments).Error
	if err != nil {
		return err
	}
	return nil
}

// GetBlackList 得到黑名单列表
func (d *Dao) GetBlackList(blackList *[]model.BlackList) error {
	err := d.Model(&model.BlackList{}).Order("time").Find(blackList).Limit(1).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUpdateList 得到动态列表
func (d *Dao) GetUpdateList(updateList *[]model.Update) error {
	err := d.Model(&model.Update{}).Order("time").Find(updateList).Error
	if err != nil {
		return err
	}
	return nil
}
