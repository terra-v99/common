// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinUserWallet = "win_user_wallet"

// WinUserWallet mapped from table <win_user_wallet>
type WinUserWallet struct {
	ID            int64           `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	UID           int64           `gorm:"column:uid;type:int(11);not null;comment:用户id" json:"uid"`                                    // 用户id
	Username      string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                       // 用户名
	MerchantID    int64           `gorm:"column:merchant_id;type:int(11);not null;comment:商户id" json:"merchantId"`                     // 商户id
	Category      int64           `gorm:"column:category;type:tinyint(4);not null;default:1;comment:钱包类型:支付/游戏/活动/佣金" json:"category"` // 钱包类型:支付/游戏/活动/佣金
	Currency      int64           `gorm:"column:currency;type:tinyint(4);not null;default:1;comment:币种" json:"currency"`               // 币种
	Coin          decimal.Decimal `gorm:"column:coin;type:decimal(15,4);not null;default:0.0000;comment:账户余额" json:"coin"`             // 账户余额
	CodeAmount    decimal.Decimal `gorm:"column:code_amount;type:decimal(15,4);not null;default:0.0000;comment:打码量" json:"codeAmount"` // 打码量
	Frozen        decimal.Decimal `gorm:"column:frozen;type:decimal(32,4);not null;default:0.0000;comment:冻结金额" json:"frozen"`         // 冻结金额
	Version       int64           `gorm:"column:version;type:int(11);default:1;comment:版本号" json:"version"`                            // 版本号
	ModifyAt      int64           `gorm:"column:modify_at;type:bigint(20);not null;comment:13位时间戳" json:"modifyAt"`                    // 13位时间戳
	CodeUpdatedAt int64           `gorm:"column:code_updated_at;type:int(11);not null;comment:打码量-更新时间" json:"codeUpdatedAt"`          // 打码量-更新时间
	CreatedAt     int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt     int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinUserWallet's table name
func (*WinUserWallet) TableName() string {
	return TableNameWinUserWallet
}
