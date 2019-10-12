package model

import (
	"time"
)

type DefaultModel struct {
	Status     int       `json:"status"`
	Created    string    `json:"created"`
	Updated    string    `json:"updated"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type DefaultModelVo struct {
	Status     string `json:"status"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

const (
	 DELETED  =iota
	 ACTIVE
)
const(
	ETH =iota+1
	DM
	SM
	GH
	YB
	RB
	TY
)
