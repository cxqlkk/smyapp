package model

import "time"

type Student struct {
	ID       string    `json:"id"`
	StudentName     string    `json:"studentName"`
	PhoneNum string `json:"phoneNum"`
	Address string `json:"address"`
	BirthDay time.Time `json:"birthDay"`
	Motto string `json:"motto""` //座右铭
	Ideal string `json:"ideal"` //理想
	FavoritePerson string `json:'favoritePerson'` //最喜欢的人
	FavoriteColor string `json:"favoriteColor"` // 最喜欢的颜色
	HeroMan string `json:"heroMan"` //崇拜的人
	DefaultModel
}

type StudentVo struct {
	ID       string    `json:"id"`
	StudentName     string    `json:"studentName"`
	PhoneNum string `json:"phoneNum"`
	Address string `json:"address"`
	BirthDay string `json:"birthDay"`
	Motto string `json:"motto""` //座右铭
	Ideal string `json:"ideal"` //理想
	FavoritePerson string `json:'favoritePerson'` //最喜欢的人
	FavoriteColor string `json:"favoriteColor"` // 最喜欢的颜色
	HeroMan string `json:"heroMan"` //崇拜的人
}