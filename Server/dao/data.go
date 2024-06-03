package dao

/*func (r *Redis) GetData(data model.DataScreen) error {
	num := 0
	var WeiboNews model.WeiboNews
	var BNews model.BNews
	var weibo model.Weibo
	var b model.B

	// 假设列表的键是"rumors"
	listKey := "weibo_posts"

	// 获取列表中的所有元素
	listElements, err := r.LRange(context.Background(), listKey, 0, 19).Result()
	if err != nil {
		panic(err)
	}

	// 遍历列表中的元素
	for _, element := range listElements {
		// 种子应从当前时间获取，以保证每次运行时生成的随机数序列不同
		rand.Seed(time.Now().UnixNano())
		if num > MAXNUM {
			_, _, WeiboNews.Content, err = util.GetKey(element)
			WeiboNews.Possibility = 0.8 + rand.Float64()
			data.WeiboNews = append(data.WeiboNews, WeiboNews)
		} else {
			_, _, WeiboNews.Content, err = util.GetKey(element)
			BNews.Possibility = 0.8 + rand.Float64()
			data.BNews = append(data.BNews, BNews)
		}
		num++
	}

}*/
