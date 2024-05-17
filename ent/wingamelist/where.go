// Code generated by ent, DO NOT EDIT.

package wingamelist

import (
	"entgo.io/ent/dialect/sql"
	"github.com/terra-v99/common/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldCode, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldName, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldIcon, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldGroupID, v))
}

// PlatListID applies equality check predicate on the "plat_list_id" field. It's identical to PlatListIDEQ.
func PlatListID(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldPlatListID, v))
}

// RevenueRate applies equality check predicate on the "revenue_rate" field. It's identical to RevenueRateEQ.
func RevenueRate(v float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldRevenueRate, v))
}

// Maintenance applies equality check predicate on the "maintenance" field. It's identical to MaintenanceEQ.
func Maintenance(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldMaintenance, v))
}

// GameCount applies equality check predicate on the "game_count" field. It's identical to GameCountEQ.
func GameCount(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldGameCount, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldRemark, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldSort, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedUser applies equality check predicate on the "updated_user" field. It's identical to UpdatedUserEQ.
func UpdatedUser(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldUpdatedUser, v))
}

// OperatorName applies equality check predicate on the "operator_name" field. It's identical to OperatorNameEQ.
func OperatorName(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldOperatorName, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldCategory, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldCode, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldName, v))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldIcon, v))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldIcon, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldGroupID, v))
}

// PlatListIDEQ applies the EQ predicate on the "plat_list_id" field.
func PlatListIDEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldPlatListID, v))
}

// PlatListIDNEQ applies the NEQ predicate on the "plat_list_id" field.
func PlatListIDNEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldPlatListID, v))
}

// PlatListIDIn applies the In predicate on the "plat_list_id" field.
func PlatListIDIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldPlatListID, vs...))
}

// PlatListIDNotIn applies the NotIn predicate on the "plat_list_id" field.
func PlatListIDNotIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldPlatListID, vs...))
}

// PlatListIDGT applies the GT predicate on the "plat_list_id" field.
func PlatListIDGT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldPlatListID, v))
}

// PlatListIDGTE applies the GTE predicate on the "plat_list_id" field.
func PlatListIDGTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldPlatListID, v))
}

// PlatListIDLT applies the LT predicate on the "plat_list_id" field.
func PlatListIDLT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldPlatListID, v))
}

// PlatListIDLTE applies the LTE predicate on the "plat_list_id" field.
func PlatListIDLTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldPlatListID, v))
}

// RevenueRateEQ applies the EQ predicate on the "revenue_rate" field.
func RevenueRateEQ(v float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldRevenueRate, v))
}

// RevenueRateNEQ applies the NEQ predicate on the "revenue_rate" field.
func RevenueRateNEQ(v float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldRevenueRate, v))
}

// RevenueRateIn applies the In predicate on the "revenue_rate" field.
func RevenueRateIn(vs ...float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldRevenueRate, vs...))
}

// RevenueRateNotIn applies the NotIn predicate on the "revenue_rate" field.
func RevenueRateNotIn(vs ...float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldRevenueRate, vs...))
}

// RevenueRateGT applies the GT predicate on the "revenue_rate" field.
func RevenueRateGT(v float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldRevenueRate, v))
}

// RevenueRateGTE applies the GTE predicate on the "revenue_rate" field.
func RevenueRateGTE(v float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldRevenueRate, v))
}

// RevenueRateLT applies the LT predicate on the "revenue_rate" field.
func RevenueRateLT(v float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldRevenueRate, v))
}

// RevenueRateLTE applies the LTE predicate on the "revenue_rate" field.
func RevenueRateLTE(v float64) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldRevenueRate, v))
}

// MaintenanceEQ applies the EQ predicate on the "maintenance" field.
func MaintenanceEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldMaintenance, v))
}

// MaintenanceNEQ applies the NEQ predicate on the "maintenance" field.
func MaintenanceNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldMaintenance, v))
}

// MaintenanceIn applies the In predicate on the "maintenance" field.
func MaintenanceIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldMaintenance, vs...))
}

// MaintenanceNotIn applies the NotIn predicate on the "maintenance" field.
func MaintenanceNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldMaintenance, vs...))
}

// MaintenanceGT applies the GT predicate on the "maintenance" field.
func MaintenanceGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldMaintenance, v))
}

// MaintenanceGTE applies the GTE predicate on the "maintenance" field.
func MaintenanceGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldMaintenance, v))
}

// MaintenanceLT applies the LT predicate on the "maintenance" field.
func MaintenanceLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldMaintenance, v))
}

// MaintenanceLTE applies the LTE predicate on the "maintenance" field.
func MaintenanceLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldMaintenance, v))
}

// MaintenanceContains applies the Contains predicate on the "maintenance" field.
func MaintenanceContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldMaintenance, v))
}

// MaintenanceHasPrefix applies the HasPrefix predicate on the "maintenance" field.
func MaintenanceHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldMaintenance, v))
}

// MaintenanceHasSuffix applies the HasSuffix predicate on the "maintenance" field.
func MaintenanceHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldMaintenance, v))
}

// MaintenanceIsNil applies the IsNil predicate on the "maintenance" field.
func MaintenanceIsNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIsNull(FieldMaintenance))
}

// MaintenanceNotNil applies the NotNil predicate on the "maintenance" field.
func MaintenanceNotNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotNull(FieldMaintenance))
}

// MaintenanceEqualFold applies the EqualFold predicate on the "maintenance" field.
func MaintenanceEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldMaintenance, v))
}

// MaintenanceContainsFold applies the ContainsFold predicate on the "maintenance" field.
func MaintenanceContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldMaintenance, v))
}

// GameCountEQ applies the EQ predicate on the "game_count" field.
func GameCountEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldGameCount, v))
}

// GameCountNEQ applies the NEQ predicate on the "game_count" field.
func GameCountNEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldGameCount, v))
}

// GameCountIn applies the In predicate on the "game_count" field.
func GameCountIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldGameCount, vs...))
}

// GameCountNotIn applies the NotIn predicate on the "game_count" field.
func GameCountNotIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldGameCount, vs...))
}

// GameCountGT applies the GT predicate on the "game_count" field.
func GameCountGT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldGameCount, v))
}

// GameCountGTE applies the GTE predicate on the "game_count" field.
func GameCountGTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldGameCount, v))
}

// GameCountLT applies the LT predicate on the "game_count" field.
func GameCountLT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldGameCount, v))
}

// GameCountLTE applies the LTE predicate on the "game_count" field.
func GameCountLTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldGameCount, v))
}

// GameCountIsNil applies the IsNil predicate on the "game_count" field.
func GameCountIsNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIsNull(FieldGameCount))
}

// GameCountNotNil applies the NotNil predicate on the "game_count" field.
func GameCountNotNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotNull(FieldGameCount))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldRemark, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldSort, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int32) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedUserEQ applies the EQ predicate on the "updated_user" field.
func UpdatedUserEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldUpdatedUser, v))
}

// UpdatedUserNEQ applies the NEQ predicate on the "updated_user" field.
func UpdatedUserNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldUpdatedUser, v))
}

// UpdatedUserIn applies the In predicate on the "updated_user" field.
func UpdatedUserIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldUpdatedUser, vs...))
}

// UpdatedUserNotIn applies the NotIn predicate on the "updated_user" field.
func UpdatedUserNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldUpdatedUser, vs...))
}

// UpdatedUserGT applies the GT predicate on the "updated_user" field.
func UpdatedUserGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldUpdatedUser, v))
}

// UpdatedUserGTE applies the GTE predicate on the "updated_user" field.
func UpdatedUserGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldUpdatedUser, v))
}

// UpdatedUserLT applies the LT predicate on the "updated_user" field.
func UpdatedUserLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldUpdatedUser, v))
}

// UpdatedUserLTE applies the LTE predicate on the "updated_user" field.
func UpdatedUserLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldUpdatedUser, v))
}

// UpdatedUserContains applies the Contains predicate on the "updated_user" field.
func UpdatedUserContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldUpdatedUser, v))
}

// UpdatedUserHasPrefix applies the HasPrefix predicate on the "updated_user" field.
func UpdatedUserHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldUpdatedUser, v))
}

// UpdatedUserHasSuffix applies the HasSuffix predicate on the "updated_user" field.
func UpdatedUserHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldUpdatedUser, v))
}

// UpdatedUserIsNil applies the IsNil predicate on the "updated_user" field.
func UpdatedUserIsNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIsNull(FieldUpdatedUser))
}

// UpdatedUserNotNil applies the NotNil predicate on the "updated_user" field.
func UpdatedUserNotNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotNull(FieldUpdatedUser))
}

// UpdatedUserEqualFold applies the EqualFold predicate on the "updated_user" field.
func UpdatedUserEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldUpdatedUser, v))
}

// UpdatedUserContainsFold applies the ContainsFold predicate on the "updated_user" field.
func UpdatedUserContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldUpdatedUser, v))
}

// OperatorNameEQ applies the EQ predicate on the "operator_name" field.
func OperatorNameEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldOperatorName, v))
}

// OperatorNameNEQ applies the NEQ predicate on the "operator_name" field.
func OperatorNameNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldOperatorName, v))
}

// OperatorNameIn applies the In predicate on the "operator_name" field.
func OperatorNameIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldOperatorName, vs...))
}

// OperatorNameNotIn applies the NotIn predicate on the "operator_name" field.
func OperatorNameNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldOperatorName, vs...))
}

// OperatorNameGT applies the GT predicate on the "operator_name" field.
func OperatorNameGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldOperatorName, v))
}

// OperatorNameGTE applies the GTE predicate on the "operator_name" field.
func OperatorNameGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldOperatorName, v))
}

// OperatorNameLT applies the LT predicate on the "operator_name" field.
func OperatorNameLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldOperatorName, v))
}

// OperatorNameLTE applies the LTE predicate on the "operator_name" field.
func OperatorNameLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldOperatorName, v))
}

// OperatorNameContains applies the Contains predicate on the "operator_name" field.
func OperatorNameContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldOperatorName, v))
}

// OperatorNameHasPrefix applies the HasPrefix predicate on the "operator_name" field.
func OperatorNameHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldOperatorName, v))
}

// OperatorNameHasSuffix applies the HasSuffix predicate on the "operator_name" field.
func OperatorNameHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldOperatorName, v))
}

// OperatorNameIsNil applies the IsNil predicate on the "operator_name" field.
func OperatorNameIsNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIsNull(FieldOperatorName))
}

// OperatorNameNotNil applies the NotNil predicate on the "operator_name" field.
func OperatorNameNotNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotNull(FieldOperatorName))
}

// OperatorNameEqualFold applies the EqualFold predicate on the "operator_name" field.
func OperatorNameEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldOperatorName, v))
}

// OperatorNameContainsFold applies the ContainsFold predicate on the "operator_name" field.
func OperatorNameContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldOperatorName, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryIsNil applies the IsNil predicate on the "category" field.
func CategoryIsNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldIsNull(FieldCategory))
}

// CategoryNotNil applies the NotNil predicate on the "category" field.
func CategoryNotNil() predicate.WinGameList {
	return predicate.WinGameList(sql.FieldNotNull(FieldCategory))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.WinGameList {
	return predicate.WinGameList(sql.FieldContainsFold(FieldCategory, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WinGameList) predicate.WinGameList {
	return predicate.WinGameList(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WinGameList) predicate.WinGameList {
	return predicate.WinGameList(func(s *sql.Selector) {
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
func Not(p predicate.WinGameList) predicate.WinGameList {
	return predicate.WinGameList(func(s *sql.Selector) {
		p(s.Not())
	})
}
