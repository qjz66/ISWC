package dao

import (
	"Server/model"
	"crypto/md5"
	"encoding/hex"
)

// InsertUser将用户信息插入到数据库
func (d *Dao) InsertUser(username, password string) (id uint, err error) {
	md5password := md5.Sum([]byte(password))
	user := model.User{RumorCount: new(int64), Times: new(int64), Username: username, Password: hex.EncodeToString(md5password[:])}
	result := d.Create(&user)
	return user.ID, result.Error
}

func (d *Dao) GetIDByName(username, password string) (id uint) {
	user := model.User{}
	md5password := md5.Sum([]byte(password))
	result := d.Where("username = ? and password = ?", username, hex.EncodeToString(md5password[:])).First(&user)
	if result.RowsAffected == 0 {
		return 0
	}
	return user.ID
}
