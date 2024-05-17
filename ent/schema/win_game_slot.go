// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.skig.tech/zero-core/common/utils"
)

type WinGameSlot struct {
	ent.Schema
}

func (WinGameSlot) Fields() []ent.Field {
	return []ent.Field{field.String("id").Comment("ID(关联BrandGameId)"), field.Int32("game_id").Comment("游戏ID(关联game_list)"), field.Int8("game_group_id").Comment("游戏大类类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能"), field.Int32("plat_id").Comment("游戏平台id"), field.String("provider").Optional().Comment("游戏提供者"), field.String("name").Comment("简体名称"), field.String("name_zh").Comment("游戏名字(中文)"), field.String("img").Comment("英文图片"), field.String("img_new").Optional().Comment("新版游戏图片"), field.Bool("is_new").Comment("是否新游戏:1-是 0-否"), field.Int8("is_casino").Comment("是否推荐主页 0否 1是"), field.String("game_type_id").Optional().Comment("游戏类型ID(0r code)"), field.String("game_type_name").Optional().Comment("游戏类型名称"), field.Int32("favorite_star").Comment("收藏值"), field.Int32("hot_star").Comment("热度"), field.Int32("sort").Comment("排序"), field.Int8("status").Comment("状态:1-启用 0-停用"), field.Int8("device").Optional().Comment("设备:0-all 1-pc 2-h5"), field.Int32("created_at"), field.Int32("updated_at"), field.String("updated_user").Optional().Comment("最后更新人"), field.String("maintenance").Optional().Comment("维护信息"), field.String("operator_name").Optional().Comment("操作人姓名")}
}
func (WinGameSlot) Edges() []ent.Edge {
	return nil
}
func (WinGameSlot) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "win_game_slot"}}
}
func (d WinGameSlot) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		utils.CacheInterceptor,
	}
}
