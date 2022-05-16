package dao

import "time"

type User struct{
	DataBase
	Utype string `json:"utype" gorm:"type:int(11);comment:类型"`
	Pid string `json:"pid" gorm:"type:int(11);comment:上级 id"`
	Unumber string `json:"unumber" gorm:"type:varchar(30);comment:账号"`
	Upwd string `json:"upwd" gorm:"type:varchar(360);comment:密码"`
	Upwds string `json:"upwds" gorm:"type:varchar(360);comment:Upwds"`
	Uxingming string `json:"uxingming" gorm:"type:varchar(20);comment:姓名"`
	Unicheng string `json:"unicheng" gorm:"type:varchar(20);comment:昵称"`
	Usfz string `json:"usfz" gorm:"type:varchar(20);comment:Usfz"`
	Usex string `json:"usex" gorm:"type:int(11);comment:性别"`
	Uemail string `json:"uemail" gorm:"type:varchar(60);comment:邮箱"`
	Utel string `json:"utel" gorm:"type:varchar(20);comment:手机号"`
	Utime time.Time `json:"utime" gorm:"type:datetime;comment:注册时间"`
	Uaddress string `json:"uaddress" gorm:"type:varchar(200);comment:钱包地址"`
	Ustate string `json:"ustate" gorm:"type:int(11);comment:1正常 2禁用"`
	IsGoogle string `json:"isGoogle" gorm:"type:int(11);comment:是否绑定google"`
	GoogleKey string `json:"googleKey" gorm:"type:varchar(128);comment:GoogleKey"`
	CreateGoogleKeyTime time.Time `json:"createGoogleKeyTime" gorm:"type:datetime;comment:CreateGoogleKeyTime"`
	Token string `json:"token" gorm:"type:varchar(258);comment:Token"`
}

func (user *User) GetOneById(Id uint) User {
	var userInfo User
	user.Db().Where("uid = ?", Id).First(&userInfo)
	return userInfo
}