package service

// 调用dao层同名函数
func (svc *Service) InsertUser(username, password string) (id uint, err error) {
	return svc.d.InsertUser(username, password)
}

// 调用dao层同名函数
func (svc *Service) GetIDByName(username, password string) (id uint) {
	return svc.d.GetIDByName(username, password)
}
