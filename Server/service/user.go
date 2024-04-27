package service

// 调用dao层同名函数
func (svc *Service) InsertUser(username, password string) (id uint, err error) {
	return svc.d.InsertUser(username, password)
}
