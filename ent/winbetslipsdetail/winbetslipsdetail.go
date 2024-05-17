// Code generated by ent, DO NOT EDIT.

package winbetslipsdetail

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the winbetslipsdetail type in the database.
	Label = "win_betslips_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldXbUID holds the string denoting the xb_uid field in the database.
	FieldXbUID = "xb_uid"
	// FieldXbUsername holds the string denoting the xb_username field in the database.
	FieldXbUsername = "xb_username"
	// FieldBetJSON holds the string denoting the bet_json field in the database.
	FieldBetJSON = "bet_json"
	// FieldRewardJSON holds the string denoting the reward_json field in the database.
	FieldRewardJSON = "reward_json"
	// FieldRefundJSON holds the string denoting the refund_json field in the database.
	FieldRefundJSON = "refund_json"
	// Table holds the table name of the winbetslipsdetail in the database.
	Table = "win_betslips_details"
)

// Columns holds all SQL columns for winbetslipsdetail fields.
var Columns = []string{
	FieldID,
	FieldXbUID,
	FieldXbUsername,
	FieldBetJSON,
	FieldRewardJSON,
	FieldRefundJSON,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the WinBetslipsDetail queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByXbUID orders the results by the xb_uid field.
func ByXbUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXbUID, opts...).ToFunc()
}

// ByXbUsername orders the results by the xb_username field.
func ByXbUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXbUsername, opts...).ToFunc()
}

// ByBetJSON orders the results by the bet_json field.
func ByBetJSON(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBetJSON, opts...).ToFunc()
}

// ByRewardJSON orders the results by the reward_json field.
func ByRewardJSON(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRewardJSON, opts...).ToFunc()
}

// ByRefundJSON orders the results by the refund_json field.
func ByRefundJSON(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefundJSON, opts...).ToFunc()
}
