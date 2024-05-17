// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/terra-v99/common/utils"
)

type WinPlatList struct {
	ent.Schema
}

func (WinPlatList) Fields() []ent.Field {
	return []ent.Field{field.Int32("id"), field.String("code").Unique().Comment("平台编码"), field.String("name").Unique().Comment("平台名称"), field.String("config").Optional().Comment("配置信息"), field.String("rate").Optional().Comment("费率"), field.Int8("sort").Comment("排序(从高到底、ID降序)"), field.Int8("status").Comment("状态: 1-启用 0-停用"), field.Int32("created_at"), field.Int32("updated_at")}
}
func (WinPlatList) Edges() []ent.Edge {
	return nil
}
func (WinPlatList) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "win_plat_list"}}
}

func (d WinPlatList) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		utils.CacheInterceptor,
	}
}
