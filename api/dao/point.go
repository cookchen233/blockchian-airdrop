package dao

import (
     "time"
)

type TbPoint struct {
    Uid string `json:"uid" gorm:"type:int(11);comment:用户表 id"` 
    Utype string `json:"utype" gorm:"type:int(11);comment:币种"` 
    Dremark string `json:"dremark" gorm:"type:varchar(500);comment:备注"` 
    Jifen string `json:"jifen" gorm:"type:decimal(19,9);comment:积分个数"` 
    Ftime time.Time `json:"ftime" gorm:"type:datetime;comment:时间"` 
    Ustate string `json:"ustate" gorm:"type:int(11);comment:类型"`
}