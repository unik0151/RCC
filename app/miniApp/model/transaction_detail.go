package model

import "gorm.io/gorm"

type Unit int

const (
	WEI Unit = iota
	GWEI
	ETH
)

type TransactionDetail struct {
	gorm.Model
	FromAddress string `gorm:"column:from_address;type:varchar(20)" json:"from_address"`
	ToAddress   string `gorm:"column:to_address;type:varchar(20)" json:"to_address"`
	Amount      int64  `gorm:"column:amount;type:bigint" json:"amount"`
	Gas         string `gorm:"column:gas;type:bigint" json:"gas"`
}

type ReceiptDetail struct {
}

type TransferDetail struct {
	from   string
	to     string
	amount float64
	unit   Unit
}
