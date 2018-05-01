package entity

type User struct {
	Uid          string `json:"uid" form:"f_uid"`
	Pwd          string `json:"pwd" form:"f_pwd"`
	Name         string `json:"name" form:"f_name"`
}

type Favorite struct {
	Fid          string `json:"fid" form:"f_fid"`
	Uid          string `json:"uid" form:"f_uid"`
	Vid          string `json:"vid" form:"f_vid"`
	CreateTime   string `json:"createTime" form:"f_create_time"`
}

type Volume struct {
	Fid          string `json:"fid" form:"f_fid"`
	Name         string `json:"name" form:"f_name"`
	Cover        string `json:"cover" form:"f_cover"`
	Desc         string `json:"desc" form:"f_desc"`
	Uid          string `json:"uid" form:"f_uid"`
	Status       string `json:"status" form:"f_status"`
	CreateTime   string `json:"createTime" form:"f_create_time"`
}

type VolumeSheet struct {
	Vsid         string `json:"vsid" form:"f_vsid"`
	Vid          string `json:"vid" form:"f_vid"`
	Sid          string `json:"sid" form:"f_sid"`
	CreateTime   string `json:"createTime" form:"f_create_time"`
}