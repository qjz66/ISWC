package service

import "Server/model"

// InsertUser 调用dao层同名函数
func (svc *Service) InsertUser(username, password string) (id int64, err error) {
	return svc.d.InsertUser(username, password)
}

// GetIDByName 调用dao层同名函数
func (svc *Service) GetIDByName(username, password string) (id int64) {
	return svc.d.GetIDByName(username, password)
}

// UpdatePassword 调用dao层同名函数
func (svc *Service) UpdatePassword(username, password string) error {
	return svc.d.UpdatePassword(username, password)
}

// UpdateUsername 调用dao层同名函数
func (svc *Service) UpdateUsername(id int64, newUsername string) error {
	return svc.d.UpdateUsername(id, newUsername)
}

// History 调用dao层同名函数
func (svc *Service) History(rumors *[]model.Info, id int64) error {
	return svc.d.History(rumors, id)
}
