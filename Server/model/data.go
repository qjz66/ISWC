package model

type WeiboNews struct {
	Content     string  `json:"content"`
	Possibility float64 `json:"possibility"`
}

type BNews struct {
	Content     string  `json:"content"`
	Possibility float64 `json:"possibility"`
}

type Weibo struct {
	Time   string `json:"time"`
	Number int    `json:"number"`
}

type B struct {
	Time   string `json:"time"`
	Number int    `json:"number"`
}

type NewsNumber struct {
	Weibo []Weibo `json:"weibo"`
	B     []B     `json:"b"`
}

type NewsRatio struct {
	Platform    string `json:"platform"`
	Possibility string `json:"possibility"`
}

type DataScreen struct {
	WeiboNews   []WeiboNews `json:"weibo_news"`
	BNews       []BNews     `json:"bNews"`
	NewsNumb    NewsNumber  `json:"news_number"`
	NewsRatio   []NewsRatio `json:"news_ratio"`
	RumourRatio float64     `json:"rumour_ratio"`
}
