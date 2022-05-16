package dto

type LoginInput struct {
	Username string `form:"username" json:"username" comment:"用户名"  validate:"required" example:""`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:""`
	RePassword string `form:"re_password" json:"re_password" comment:"确认密码"   validate:"required,eqfield=Password" example:""`
}

type ListPageInput struct {
	PageSize int    `form:"page_size" json:"page_size" comment:"每页记录数" validate:"" example:"10"`
	Page     int    `form:"page" json:"page" comment:"页数" validate:"required" example:"1"`
	Name     string `form:"name" json:"name" comment:"姓名" validate:"" example:""`
}

type AddUserInput struct {
	Name  string `form:"name" json:"name" comment:"姓名" validate:"required"`
	Sex   int    `form:"sex" json:"sex" comment:"性别" validate:""`
	Age   int    `form:"age" json:"age" comment:"年龄" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" comment:"生日" validate:"required"`
	Addr  string `form:"addr" json:"addr" comment:"地址" validate:"required"`
}

type EditUserInput struct {
	Id    int    `form:"id" json:"id" comment:"ID" validate:"required"`
	Name  string `form:"name" json:"name" comment:"姓名" validate:"required"`
	Sex   int    `form:"sex" json:"sex" comment:"性别" validate:""`
	Age   int    `form:"age" json:"age" comment:"年龄" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" comment:"生日" validate:"required"`
	Addr  string `form:"addr" json:"addr" comment:"地址" validate:"required"`
}

type RemoveUserInput struct {
	IDS string `form:"ids" json:"ids" comment:"IDS" validate:"required"`
}
