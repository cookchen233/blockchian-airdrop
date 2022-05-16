package dao

type TbPlugin struct {
    Name string `json:"name" gorm:"type:varchar(50);comment:名称"` 
    Status string `json:"status" gorm:"type:int(11);comment:状态"` 
    Download string `json:"download" gorm:"type:varchar(256);comment:下载地址"`
}