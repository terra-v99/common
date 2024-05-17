// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserFacebookLogin = "win_user_facebook_login"

// WinUserFacebookLogin mapped from table <win_user_facebook_login>
type WinUserFacebookLogin struct {
	ID             int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键" json:"id,string"`                 // 主键
	UserID         int64  `gorm:"column:user_id;type:int;not null;comment:客户表ID" json:"userId"`                                 // 客户表ID
	Username       string `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                        // 用户名
	FacebookUserID string `gorm:"column:facebook_user_id;type:varchar(16);not null;comment:Facebook用户ID" json:"facebookUserId"` // Facebook用户ID
	CreatedAt      int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                              // 创建时间
}

// TableName WinUserFacebookLogin's table name
func (*WinUserFacebookLogin) TableName() string {
	return TableNameWinUserFacebookLogin
}