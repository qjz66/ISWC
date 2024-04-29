package dao

import "Server/model"

// InsertUser将用户信息插入到数据库
func (d *Dao) InsertUser(username, password string) (id uint, err error) {
	user := model.User{RumorCount: new(int64), Times: new(int64), Username: username, Password: password}
	result := d.Create(&user)
	return user.ID, result.Error
}

func (d *Dao) GetIDByName(username, password string) (id uint) {
	user := model.User{}
	result := d.Where("username = ? and password = ?", username, password).First(&user)
	if result.RowsAffected == 0 {
		return 0
	}
	return user.ID
}
