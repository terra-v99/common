// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type WinDict struct {
	ent.Schema
}

func (WinDict) Fields() []ent.Field {
	return []ent.Field{field.Int32("id"), field.String("title").Comment("名称"), field.String("category").Comment("种类"), field.Bool("status").Comment("状态:1-启用 0-禁用"), field.Int32("created_at"), field.Int32("updated_at")}
}
func (WinDict) Edges() []ent.Edge {
	return nil
}
func (WinDict) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "win_dict"}}
}