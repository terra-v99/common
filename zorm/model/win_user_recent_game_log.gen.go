// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserRecentGameLog = "win_user_recent_game_log"

// WinUserRecentGameLog mapped from table <win_user_recent_game_log>
type WinUserRecentGameLog struct {
	ID         int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id,string"`
	UID        int64  `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                                // UID
	GameName   string `gorm:"column:game_name;type:varchar(255);not null;default:-;comment:游戏名称" json:"gameName"` // 游戏名称
	GameType   int64  `gorm:"column:game_type;type:int;not null;comment:游戏分类" json:"gameType"`                    // 游戏分类
	KeyID      string `gorm:"column:key_id;type:varchar(300);not null;comment:游戏唯一" json:"keyId"`                 // 游戏唯一
	SlotID     string `gorm:"column:slot_id;type:varchar(255);not null;comment:slot_id" json:"slotId"`            // slot_id
	GameListID int64  `gorm:"column:game_list_id;type:int;not null;comment:游戏分类id" json:"gameListId"`             // 游戏分类id
	GamePlatID int64  `gorm:"column:game_plat_id;type:int;not null;comment:平台id" json:"gamePlatId"`               // 平台id
	CreatedAt  int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                    // 创建时间
}

// TableName WinUserRecentGameLog's table name
func (*WinUserRecentGameLog) TableName() string {
	return TableNameWinUserRecentGameLog
}