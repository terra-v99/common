// Code generated by ent, DO NOT EDIT.

package windictitem

import (
	"entgo.io/ent/dialect/sql"
	"github.com/terra-v99/common/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldCode, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldTitle, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldRemark, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldSort, v))
}

// ReferID applies equality check predicate on the "refer_id" field. It's identical to ReferIDEQ.
func ReferID(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldReferID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldStatus, v))
}

// IsShow applies equality check predicate on the "is_show" field. It's identical to IsShowEQ.
func IsShow(v int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldIsShow, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldContainsFold(FieldCode, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldContainsFold(FieldTitle, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldContainsFold(FieldRemark, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldSort, v))
}

// ReferIDEQ applies the EQ predicate on the "refer_id" field.
func ReferIDEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldReferID, v))
}

// ReferIDNEQ applies the NEQ predicate on the "refer_id" field.
func ReferIDNEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldReferID, v))
}

// ReferIDIn applies the In predicate on the "refer_id" field.
func ReferIDIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldReferID, vs...))
}

// ReferIDNotIn applies the NotIn predicate on the "refer_id" field.
func ReferIDNotIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldReferID, vs...))
}

// ReferIDGT applies the GT predicate on the "refer_id" field.
func ReferIDGT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldReferID, v))
}

// ReferIDGTE applies the GTE predicate on the "refer_id" field.
func ReferIDGTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldReferID, v))
}

// ReferIDLT applies the LT predicate on the "refer_id" field.
func ReferIDLT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldReferID, v))
}

// ReferIDLTE applies the LTE predicate on the "refer_id" field.
func ReferIDLTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldReferID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldStatus, v))
}

// IsShowEQ applies the EQ predicate on the "is_show" field.
func IsShowEQ(v int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldIsShow, v))
}

// IsShowNEQ applies the NEQ predicate on the "is_show" field.
func IsShowNEQ(v int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldIsShow, v))
}

// IsShowIn applies the In predicate on the "is_show" field.
func IsShowIn(vs ...int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldIsShow, vs...))
}

// IsShowNotIn applies the NotIn predicate on the "is_show" field.
func IsShowNotIn(vs ...int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldIsShow, vs...))
}

// IsShowGT applies the GT predicate on the "is_show" field.
func IsShowGT(v int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldIsShow, v))
}

// IsShowGTE applies the GTE predicate on the "is_show" field.
func IsShowGTE(v int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldIsShow, v))
}

// IsShowLT applies the LT predicate on the "is_show" field.
func IsShowLT(v int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldIsShow, v))
}

// IsShowLTE applies the LTE predicate on the "is_show" field.
func IsShowLTE(v int8) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldIsShow, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int32) predicate.WinDictItem {
	return predicate.WinDictItem(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WinDictItem) predicate.WinDictItem {
	return predicate.WinDictItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WinDictItem) predicate.WinDictItem {
	return predicate.WinDictItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WinDictItem) predicate.WinDictItem {
	return predicate.WinDictItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
