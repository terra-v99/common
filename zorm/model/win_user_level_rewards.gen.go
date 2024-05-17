// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinUserLevelReward = "win_user_level_rewards"

// WinUserLevelReward mapped from table <win_user_level_rewards>
type WinUserLevelReward struct {
	ID          int64           `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	UID         int64           `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                                          // UID
	Username    string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                        // 用户名
	Coin        decimal.Decimal `gorm:"column:coin;type:decimal(15,4);not null;default:0.0000;comment:彩金金额" json:"coin"`              // 彩金金额
	CoinBefore  decimal.Decimal `gorm:"column:coin_before;type:decimal(15,4);not null;default:0.0000;comment:即时金额" json:"coinBefore"` // 即时金额
	Category    int64           `gorm:"column:category;type:int;not null;comment:福利类型:0-晋级彩金 1-生日礼金 2-周彩金3 -月彩金" json:"category"`     // 福利类型:0-晋级彩金 1-生日礼金 2-周彩金3 -月彩金
	ReceiveAt   int64           `gorm:"column:receive_at;type:int;comment:领取时间" json:"receiveAt"`                                     // 领取时间
	UserLevelID int64           `gorm:"column:user_level_id;type:int;comment:用户等级编号" json:"userLevelId"`                              // 用户等级编号
	FlowClaim   int64           `gorm:"column:flow_claim;type:int;comment:流水倍数" json:"flowClaim"`                                     // 流水倍数
	Codes       decimal.Decimal `gorm:"column:codes;type:decimal(15,4);comment:所需打码量" json:"codes"`                                   // 所需打码量
	EndedAt     int64           `gorm:"column:ended_at;type:int;comment:领取结束时间" json:"endedAt"`                                       // 领取结束时间
	Status      int64           `gorm:"column:status;type:int;not null;comment:状态:0-待领取 1-已领取 2-已过期" json:"status"`                   // 状态:0-待领取 1-已领取 2-已过期
	CreatedAt   int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt   int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinUserLevelReward's table name
func (*WinUserLevelReward) TableName() string {
	return TableNameWinUserLevelReward
}