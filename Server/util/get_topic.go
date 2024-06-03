package util

// GetKey 用来获取redis数据库中键
import (
	"strings"
)

func GetKey(data string) (topic, uid, content string, err error) {
	// 将数据按"Key: "分割
	splitKey := strings.SplitN(data, "Key: ", 2)

	// 提取Key后面的内容，然后按冒号分割
	keyParts := strings.Split(splitKey[1], ":")

	// 提取topic和uid
	topic = keyParts[0] // 第一个元素是topic
	uid = keyParts[1]   // 第二个元素是uid

	// 提取Value部分，去除"Value: "并去除前后空白
	valuePart := strings.TrimSpace(strings.SplitN(splitKey[1], ", Value: ", 2)[1])
	content = strings.TrimSpace(valuePart)

	return topic, uid, content, nil
}
