package service

import "Server/model"

// Favorite 调用dao层同名函数
func (svc *Service) Favorite(authorId, updateId int64) (int64, error) {
	return svc.d.Favorite(authorId, updateId)
}

// Update 调用dao层同名函数
func (svc *Service) Update(content string, id int64) (int64, error) {
	return svc.d.Update(content, id)
}

// Comment 调用dao层同名函数
func (svc *Service) Comment(comment model.Comment) error {
	return svc.d.Comment(comment)
}

// GetComments 调用dao层同名函数
func (svc *Service) GetComments(updateId int64, comments *[]model.Comment) error {
	return svc.d.GetComments(updateId, comments)
}

// GetBlackList 调用dao层同名函数
func (svc *Service) GetBlackList(blackList *[]model.BlackList) error {
	return svc.d.GetBlackList(blackList)
}

// GetUpdateList 调用dao层同名函数
func (svc *Service) GetUpdateList(updateList *[]model.Update) error {
	return svc.d.GetUpdateList(updateList)
}
