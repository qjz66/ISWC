package dao

import (
	"Server/model"
	"Server/util"
	"context"
)

const MAXNUM = 10

// GetTopic 获得用户最常用的topic
func (d *Dao) GetTopic(id int64) (string, error) {
	var result []model.UserTopic
	var topic string
	var maxTopic int64 = 0
	err := d.Model(&model.UserTopic{}).Where("id = ?", id).Find(&result).Error
	if err != nil {
		return "", err
	}
	//该用户未检测过谣言
	if len(result) == 0 {
		return "", nil
	}
	for _, v := range result {
		if v.Num > maxTopic {
			maxTopic = v.Num
			topic = v.Topic
		}
	}
	return topic, nil
}

// PushByTopic 根据检测最多的topic推送
func (r *Redis) PushByTopic(rumors *[]model.Rumor, topic string) error {
	num := 0
	// 假设列表的键是"rumors"
	listKey := "rumors"
	var tmpRumor model.Rumor

	// 获取列表中的所有元素
	listElements, err := r.LRange(context.Background(), listKey, 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// 遍历列表中的元素
	for _, element := range listElements {
		if num > MAXNUM {
			break
		}
		//将列表反序列化获得topic、uid、content
		tmpRumor.Topic, tmpRumor.AuthorId, tmpRumor.Content, err = util.GetKey(element)
		tmpRumor.Platform = "weibo"
		if tmpRumor.Topic == topic {
			*rumors = append(*rumors, tmpRumor)
			num++
		}
	}
	return nil
}

// Push 直接将redis库中前十条推送给用户
func (r *Redis) Push(rumors *[]model.Rumor) error {
	num := 0
	// 假设列表的键是"rumors"
	listKey := "rumors"
	var tmpRumor model.Rumor

	// 获取列表中的所有元素
	listElements, err := r.LRange(context.Background(), listKey, 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// 遍历列表中的元素
	for _, element := range listElements {
		if num > MAXNUM {
			break
		}
		//将列表反序列化获得topic、uid、content
		tmpRumor.Topic, tmpRumor.AuthorId, tmpRumor.Content, err = util.GetKey(element)
		tmpRumor.Platform = "weibo"
		*rumors = append(*rumors, tmpRumor)
		num++
	}
	return nil
}
