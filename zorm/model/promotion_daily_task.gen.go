// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNamePromotionDailyTask = "promotion_daily_task"

// PromotionDailyTask mapped from table <promotion_daily_task>
type PromotionDailyTask struct {
	ID              int64           `gorm:"column:id;type:bigint(20) unsigned;primaryKey;comment:主键ID" json:"id,string"`                       // 主键ID
	PromotionsID    int64           `gorm:"column:promotions_id;type:int(11);not null;comment:优惠活动ID" json:"promotionsId"`                     // 优惠活动ID
	UserID          int64           `gorm:"column:user_id;type:int(11);not null;comment:用户ID" json:"userId"`                                   // 用户ID
	TaskTitle       string          `gorm:"column:task_title;type:varchar(65);not null;comment:任务标题: task1, task2" json:"taskTitle"`           // 任务标题: task1, task2
	BetAmount       int64           `gorm:"column:bet_amount;type:int(11);not null;comment:有效投注金额" json:"betAmount"`                           // 有效投注金额
	BetAmountFinish int64           `gorm:"column:bet_amount_finish;type:int(11);not null;comment:已投注金额" json:"betAmountFinish"`               // 已投注金额
	Bonus           int64           `gorm:"column:bonus;type:int(11);not null;comment:彩金金额" json:"bonus"`                                      // 彩金金额
	CodeMultiple    int64           `gorm:"column:code_multiple;type:int(11);not null;comment:打码倍数" json:"codeMultiple"`                       // 打码倍数
	RotatingNumber  int64           `gorm:"column:rotating_number;type:int(11);not null;comment:旋转数量" json:"rotatingNumber"`                   // 旋转数量
	PerBetAmount    decimal.Decimal `gorm:"column:per_bet_amount;type:decimal(15,2);not null;default:0.00;comment:单次投注金额" json:"perBetAmount"` // 单次投注金额
	ValidDays       int64           `gorm:"column:valid_days;type:int(11);not null;comment:有效期天数" json:"validDays"`                            // 有效期天数
	SortID          int64           `gorm:"column:sort_id;type:tinyint(4);not null;default:1;comment:任务序号ID: 1-6" json:"sortId"`               // 任务序号ID: 1-6
	Status          int64           `gorm:"column:status;type:tinyint(4);not null;comment:任务状态: 0-未完成, 1-已完成" json:"status"`                   // 任务状态: 0-未完成, 1-已完成
	TaskDate        int64           `gorm:"column:task_date;type:int(11);not null;comment:任务日期" json:"taskDate"`                               // 任务日期
	UpdateAt        int64           `gorm:"column:update_at;type:int(11);not null;comment:修改时间" json:"updateAt"`                               // 修改时间
}

// TableName PromotionDailyTask's table name
func (*PromotionDailyTask) TableName() string {
	return TableNamePromotionDailyTask
}
