package service

import "Server/model"

// InsertIntoInfo 调用dao层同名函数
func (svc *Service) InsertIntoInfo(info model.Info, topic string) error {
	return svc.d.InsertIntoInfo(info, topic)
}
