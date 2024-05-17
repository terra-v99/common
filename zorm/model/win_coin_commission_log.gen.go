// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinCoinCommissionLog = "win_coin_commission_log"

// WinCoinCommissionLog mapped from table <win_coin_commission_log>
type WinCoinCommissionLog struct {
	ID         int64           `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	UID        int64           `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                                            // UID
	Username   string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                          // 用户名
	Category   int64           `gorm:"column:category;type:tinyint;not null;default:1;comment:类型:1-返佣 2- 佣金充值 3-佣金提款" json:"category"` // 类型:1-返佣 2- 佣金充值 3-佣金提款
	ReferID    int64           `gorm:"column:refer_id;type:bigint;comment:关联ID" json:"referId"`                                        // 关联ID
	Coin       decimal.Decimal `gorm:"column:coin;type:decimal(15,4);not null;default:0.0000;comment:金额" json:"coin"`                  // 金额
	OutIn      int64           `gorm:"column:out_in;type:tinyint;not null;comment:收支类型:0-支出 1-收入" json:"outIn"`                        // 收支类型:0-支出 1-收入
	CoinBefore decimal.Decimal `gorm:"column:coin_before;type:decimal(15,4);not null;default:0.0000;comment:前金额" json:"coinBefore"`    // 前金额
	CoinAfter  decimal.Decimal `gorm:"column:coin_after;type:decimal(15,4);not null;default:0.0000;comment:帐变后金额" json:"coinAfter"`    // 帐变后金额
	CreatedAt  int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt  int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinCoinCommissionLog's table name
func (*WinCoinCommissionLog) TableName() string {
	return TableNameWinCoinCommissionLog
}