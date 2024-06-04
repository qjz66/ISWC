package dao

import (
	"Server/model"
)

// Favorite 给动态点赞
func (d *Dao) Favorite(updateId int64) (int64, error) {
	//update := &model.Update{ID: authorId}
	var update model.Update
	update.UpdateID = updateId
	err := d.Model(&model.Update{}).Where("id=?", updateId).First(&update).Error
	if err != nil {
		return 0, err
	}
	favorite := *update.Favorite + 1
	err = d.Model(&update).Update("likeNum", favorite).Error
	if err != nil {
		return 0, err
	}
	return favorite, nil
}

// Update 上传动态
func (d *Dao) Update(update model.Update) (int64, error) {

	err := d.Model(&model.Update{}).Create(&update).Error
	if err != nil {
		return 0, err
	}
	return update.UpdateID, nil
}

// Comment 给动态评论
func (d *Dao) Comment(comment model.Comment) (int64, error) {
	err := d.Model(&model.Comment{}).Create(&comment).Error
	if err != nil {
		return 0, err
	}
	return comment.ID, nil
}

// GetComments 得到评论列表
func (d *Dao) GetComments(updateId int64, comments *[]model.Comment) error {
	err := d.Model(&model.Comment{}).Where("updateId=?", updateId).Find(comments).Error
	if err != nil {
		return err
	}
	return nil
}

// GetBlackList 得到黑名单列表
func (d *Dao) GetBlackList(blackList *[]model.BlackList) error {
	err := d.Model(&model.BlackList{}).Find(blackList).Limit(1).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUpdateList 得到动态列表
func (d *Dao) GetUpdateList(updateList *[]model.Update) error {
	err := d.Model(&model.Update{}).Find(updateList).Error
	if err != nil {
		return err
	}
	for i, update := range *updateList {
		// 为每个更新的Comment分配一个新的切片
		(*updateList)[i].Comment = []model.Comment{}
		var comments []model.Comment // 注意这里使用的是新的切片变量
		err = d.Model(&model.Comment{}).Where("updateId = ?", update.UpdateID).Find(&comments).Error
		// 将查询到的评论赋值给update的Comment字段
		(*updateList)[i].Comment = comments // 确保更新了updateList中的元素
		if err != nil {
			return err
		}

	}

	return nil
}
