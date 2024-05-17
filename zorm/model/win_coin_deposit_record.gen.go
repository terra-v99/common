// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinCoinDepositRecord = "win_coin_deposit_record"

// WinCoinDepositRecord mapped from table <win_coin_deposit_record>
type WinCoinDepositRecord struct {
	ID               int64           `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	OrderID          string          `gorm:"column:order_id;type:varchar(64);not null;comment:订单号(三方平台用)" json:"orderId"`                                                          // 订单号(三方平台用)
	PlatOrderID      string          `gorm:"column:plat_order_id;type:varchar(64);not null;comment:三方平台订单号" json:"platOrderId"`                                                    // 三方平台订单号
	UID              int64           `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                                                                                  // UID
	Username         string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                                                                // 用户名
	MerchantID       int64           `gorm:"column:merchant_id;type:int;not null;comment:商户id" json:"merchantId"`                                                                  // 商户id
	Code             string          `gorm:"column:code;type:varchar(50);not null;comment:支付通道编码" json:"code"`                                                                     // 支付通道编码
	PlatName         string          `gorm:"column:plat_name;type:varchar(100);not null;comment:平台名称" json:"platName"`                                                             // 平台名称
	PlatNickName     string          `gorm:"column:plat_nick_name;type:varchar(100);comment:平台自定义名称" json:"platNickName"`                                                          // 平台自定义名称
	CoinBefore       decimal.Decimal `gorm:"column:coin_before;type:decimal(15,4);not null;comment:充值前金额" json:"coinBefore"`                                                       // 充值前金额
	PayAddress       string          `gorm:"column:pay_address;type:text;comment:加密地址" json:"payAddress"`                                                                          // 加密地址
	PayAmount        decimal.Decimal `gorm:"column:pay_amount;type:decimal(15,4);not null;default:0.0000;comment:充值金额" json:"payAmount"`                                           // 充值金额
	ExchangeRate     decimal.Decimal `gorm:"column:exchange_rate;type:decimal(15,4);not null;default:0.0000;comment:汇率" json:"exchangeRate"`                                       // 汇率
	RealAmount       decimal.Decimal `gorm:"column:real_amount;type:decimal(15,4);not null;default:0.0000;comment:到账金额" json:"realAmount"`                                         // 到账金额
	Currency         string          `gorm:"column:currency;type:varchar(50);not null;comment:币种" json:"currency"`                                                                 // 币种
	DepStatus        int64           `gorm:"column:dep_status;type:int;not null;default:9;comment:充值标识:1-首充 2-二充 9-其他" json:"depStatus"`                                           // 充值标识:1-首充 2-二充 9-其他
	Category         int64           `gorm:"column:category;type:tinyint(1);not null;comment:类型:0-钱包充值 1-佣金钱包转账充值" json:"category"`                                                // 类型:0-钱包充值 1-佣金钱包转账充值
	CategoryCurrency int64           `gorm:"column:category_currency;type:tinyint(1);not null;default:1;comment:货币类型:0-数字货币 1-法币" json:"categoryCurrency"`                         // 货币类型:0-数字货币 1-法币
	CategoryTransfer int64           `gorm:"column:category_transfer;type:tinyint;not null;default:3;comment:转账类型:1-TRC-20 2-ERC-20 3-BANK 4-PIX 5-GCASH" json:"categoryTransfer"` // 转账类型:1-TRC-20 2-ERC-20 3-BANK 4-PIX 5-GCASH
	AdminUID         int64           `gorm:"column:admin_uid;type:int;not null;comment:审核ID" json:"adminUid"`                                                                      // 审核ID
	Mark             string          `gorm:"column:mark;type:varchar(150);not null;comment:备注" json:"mark"`                                                                        // 备注
	Status           int64           `gorm:"column:status;type:tinyint;not null;default:0;comment:状态: 0-申请中 1-成功 2-失败" json:"status"`                                              // 状态: 0-申请中 1-成功 2-失败
	CreatedAt        int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt        int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinCoinDepositRecord's table name
func (*WinCoinDepositRecord) TableName() string {
	return TableNameWinCoinDepositRecord
}
