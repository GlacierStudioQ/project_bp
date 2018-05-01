package entity

import (
	"time"
)

const(
	TAB_FIELD_KEY = "form"
)

type User struct {
	Uid          string `json:"uid" form:"f_uid"`
	Pwd          string `json:"pwd" form:"f_pwd"`
	Name         string `json:"name" form:"f_name"`
}

type Favorite struct {
	Fid          string    `json:"fid" form:"f_fid"`
	Uid          string    `json:"uid" form:"f_uid"`
	Vid          string    `json:"vid" form:"f_vid"`
	CreateTime   time.Time `json:"createTime" form:"f_create_time"`
}

// 这里字段名首字母必须大写，才能被外面访问到
type Volume struct {
	Vid          string    `json:"vid" form:"f_vid"`
	Name         string    `json:"name" form:"f_name"`
	Cover        string    `json:"cover" form:"f_cover"`
	Desc         string    `json:"desc" form:"f_desc"`
	Uid          string    `json:"uid" form:"f_uid"`
	Status       int       `json:"status" form:"f_status"`
	CreateTime   time.Time `json:"createTime" form:"f_create_time"`
}

type VolumeSheet struct {
	Vsid         string    `json:"vsid" form:"f_vsid"`
	Vid          string    `json:"vid" form:"f_vid"`
	Sid          string    `json:"sid" form:"f_sid"`
	CreateTime   time.Time `json:"createTime" form:"f_create_time"`
}