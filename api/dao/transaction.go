package dao

import (
     "time"
)

type TbTransaction struct {
    Uid string `json:"uid" gorm:"type:int(11);comment:用户表 id"`
    To string `json:"to" gorm:"type:varchar(200);comment:收款地址"` 
    From string `json:"from" gorm:"type:varchar(200);comment:打款地址"` 
    Gas string `json:"gas" gorm:"type:varchar(100);comment:区块返回数据"` 
    Data string `json:"data" gorm:"type:varchar(200);comment:Data"` 
    GasPrice string `json:"gasPrice" gorm:"type:varchar(100);comment:GasPrice"` 
    Nonce string `json:"nonce" gorm:"type:varchar(100);comment:Nonce"` 
    Hash string `json:"hash" gorm:"type:varchar(200);comment:Hash"` 
    BlockHash string `json:"blockHash" gorm:"type:varchar(500);comment:BlockHash"` 
    BlockNumber string `json:"blockNumber" gorm:"type:varchar(100);comment:BlockNumber"` 
    ContractAddress string `json:"contractAddress" gorm:"type:varchar(200);comment:ContractAddress"` 
    IsAdd string `json:"isAdd" gorm:"type:int(11);comment:是否已检索"` 
    JYType string `json:"jYType" gorm:"type:int(11);comment:交易类型 1：交易 2：奖励"` 
    BType string `json:"bType" gorm:"type:int(11);comment:币种"` 
    ToAddress string `json:"toAddress" gorm:"type:varchar(200);comment:代币到账地址"` 
    ToValue string `json:"toValue" gorm:"type:decimal(19,9);comment:代币到账数量"` 
    CreatedTime time.Time `json:"createdTime" gorm:"type:datetime;comment:创建时间"` 
    ModifiedTime time.Time `json:"modifiedTime" gorm:"type:datetime;comment:修改时间"` 
    PtxtId string `json:"ptxtId" gorm:"type:varchar(200);comment:父区块交易id"` 
    IsAwarded string `json:"isAwarded" gorm:"type:int(11);comment:是否已发放奖励/币"` 
    IsShow string `json:"isShow" gorm:"type:int(11);comment:是否显示"` 
    IsPack string `json:"isPack" gorm:"type:int(11);comment:是否重新打包交易"` 
    Logs string `json:"logs" gorm:"type:varchar(10000);comment:Logs"` 
    Status string `json:"status" gorm:"type:varchar(100);comment:状态Success"` 
}