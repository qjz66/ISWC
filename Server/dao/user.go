package dao

import (
	"Server/model"
	"crypto/md5"
	"encoding/hex"
)

// InsertUser 将用户信息插入到数据库
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

// UpdatePassword 修改用户密码
func (d *Dao) UpdatePassword(username, password string) error {
	md5password := md5.Sum([]byte(password))
	err := d.Model(&model.User{}).Where("username = ?", username).Update("password", hex.EncodeToString(md5password[:])).Error
	return err
}

// UpdateUsername 修改用户名称
func (d *Dao) UpdateUsername(id int64, newUsername string) error {
	err := d.Model(&model.User{}).Where("id = ?", id).Update("username", newUsername).Error
	return err
}
