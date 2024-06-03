package service

import "Server/model"

// GetTopic 调用dao层同名函数
func (svc *Service) GetTopic(id int64) (string, error) {
	return svc.d.GetTopic(id)
}

// PushByTopic 调用dao层同名函数
func (svc *Service) PushByTopic(rumors *[]model.Rumor, topic string) error {
	return svc.r.PushByTopic(rumors, topic)
}

// Push 调用dao层同名函数
func (svc *Service) Push(rumors *[]model.Rumor) error {
	return svc.r.Push(rumors)
}
