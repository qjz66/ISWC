package dao

import "Server/model"

// InsertUser将用户信息插入到数据库
func (d *Dao) InsertUser(username, password string) (id string, err error) {
	user := model.User{RumorCount: new(int64), Times: new(int64), Username: username, Password: password}
	result := d.Create(&user)
	return user.ID, result.Error
}
