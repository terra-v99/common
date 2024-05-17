// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinAdminOperateLog = "win_admin_operate_log"

// WinAdminOperateLog mapped from table <win_admin_operate_log>
type WinAdminOperateLog struct {
	ID        int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	UID       int64  `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                   // UID
	Username  string `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"` // 用户名
	Title     string `gorm:"column:title;type:varchar(64);not null;comment:标题" json:"title"`        // 标题
	URL       string `gorm:"column:url;type:varchar(512);not null;comment:请求url" json:"url"`        // 请求url
	ReqParams string `gorm:"column:req_params;type:text;comment:请求参数" json:"reqParams"`             // 请求参数
	ResParams string `gorm:"column:res_params;type:text;comment:响应参数" json:"resParams"`             // 响应参数
	IP        string `gorm:"column:ip;type:varchar(128);not null;comment:IP地址" json:"ip"`           // IP地址
	CreatedAt int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinAdminOperateLog's table name
func (*WinAdminOperateLog) TableName() string {
	return TableNameWinAdminOperateLog
}
