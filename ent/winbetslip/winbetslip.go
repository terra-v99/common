// Code generated by ent, DO NOT EDIT.

package winbetslip

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the winbetslip type in the database.
	Label = "win_betslip"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoundID holds the string denoting the round_id field in the database.
	FieldRoundID = "round_id"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldGamePlatID holds the string denoting the game_plat_id field in the database.
	FieldGamePlatID = "game_plat_id"
	// FieldXbStatus holds the string denoting the xb_status field in the database.
	FieldXbStatus = "xb_status"
	// FieldXbUID holds the string denoting the xb_uid field in the database.
	FieldXbUID = "xb_uid"
	// FieldXbUsername holds the string denoting the xb_username field in the database.
	FieldXbUsername = "xb_username"
	// FieldXbProfit holds the string denoting the xb_profit field in the database.
	FieldXbProfit = "xb_profit"
	// FieldStake holds the string denoting the stake field in the database.
	FieldStake = "stake"
	// FieldValidStake holds the string denoting the valid_stake field in the database.
	FieldValidStake = "valid_stake"
	// FieldPayout holds the string denoting the payout field in the database.
	FieldPayout = "payout"
	// FieldCoinRefund holds the string denoting the coin_refund field in the database.
	FieldCoinRefund = "coin_refund"
	// FieldCoinBefore holds the string denoting the coin_before field in the database.
	FieldCoinBefore = "coin_before"
	// FieldGameID holds the string denoting the game_id field in the database.
	FieldGameID = "game_id"
	// FieldGameName holds the string denoting the game_name field in the database.
	FieldGameName = "game_name"
	// FieldAmountType holds the string denoting the amount_type field in the database.
	FieldAmountType = "amount_type"
	// FieldGameTypeID holds the string denoting the game_type_id field in the database.
	FieldGameTypeID = "game_type_id"
	// FieldGameGroupID holds the string denoting the game_group_id field in the database.
	FieldGameGroupID = "game_group_id"
	// FieldSportType holds the string denoting the sport_type field in the database.
	FieldSportType = "sport_type"
	// FieldDtStarted holds the string denoting the dt_started field in the database.
	FieldDtStarted = "dt_started"
	// FieldDtCompleted holds the string denoting the dt_completed field in the database.
	FieldDtCompleted = "dt_completed"
	// FieldWinTransactionID holds the string denoting the win_transaction_id field in the database.
	FieldWinTransactionID = "win_transaction_id"
	// FieldBetJSON holds the string denoting the bet_json field in the database.
	FieldBetJSON = "bet_json"
	// FieldRewardJSON holds the string denoting the reward_json field in the database.
	FieldRewardJSON = "reward_json"
	// FieldRefundJSON holds the string denoting the refund_json field in the database.
	FieldRefundJSON = "refund_json"
	// FieldCreateTimeStr holds the string denoting the create_time_str field in the database.
	FieldCreateTimeStr = "create_time_str"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the winbetslip in the database.
	Table = "win_betslips"
)

// Columns holds all SQL columns for winbetslip fields.
var Columns = []string{
	FieldID,
	FieldRoundID,
	FieldTransactionID,
	FieldGamePlatID,
	FieldXbStatus,
	FieldXbUID,
	FieldXbUsername,
	FieldXbProfit,
	FieldStake,
	FieldValidStake,
	FieldPayout,
	FieldCoinRefund,
	FieldCoinBefore,
	FieldGameID,
	FieldGameName,
	FieldAmountType,
	FieldGameTypeID,
	FieldGameGroupID,
	FieldSportType,
	FieldDtStarted,
	FieldDtCompleted,
	FieldWinTransactionID,
	FieldBetJSON,
	FieldRewardJSON,
	FieldRefundJSON,
	FieldCreateTimeStr,
	FieldCreatedAt,
	FieldUpdatedAt,
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

// OrderOption defines the ordering options for the WinBetslip queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRoundID orders the results by the round_id field.
func ByRoundID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoundID, opts...).ToFunc()
}

// ByTransactionID orders the results by the transaction_id field.
func ByTransactionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionID, opts...).ToFunc()
}

// ByGamePlatID orders the results by the game_plat_id field.
func ByGamePlatID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGamePlatID, opts...).ToFunc()
}

// ByXbStatus orders the results by the xb_status field.
func ByXbStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXbStatus, opts...).ToFunc()
}

// ByXbUID orders the results by the xb_uid field.
func ByXbUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXbUID, opts...).ToFunc()
}

// ByXbUsername orders the results by the xb_username field.
func ByXbUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXbUsername, opts...).ToFunc()
}

// ByXbProfit orders the results by the xb_profit field.
func ByXbProfit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXbProfit, opts...).ToFunc()
}

// ByStake orders the results by the stake field.
func ByStake(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStake, opts...).ToFunc()
}

// ByValidStake orders the results by the valid_stake field.
func ByValidStake(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidStake, opts...).ToFunc()
}

// ByPayout orders the results by the payout field.
func ByPayout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayout, opts...).ToFunc()
}

// ByCoinRefund orders the results by the coin_refund field.
func ByCoinRefund(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinRefund, opts...).ToFunc()
}

// ByCoinBefore orders the results by the coin_before field.
func ByCoinBefore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinBefore, opts...).ToFunc()
}

// ByGameID orders the results by the game_id field.
func ByGameID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameID, opts...).ToFunc()
}

// ByGameName orders the results by the game_name field.
func ByGameName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameName, opts...).ToFunc()
}

// ByAmountType orders the results by the amount_type field.
func ByAmountType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountType, opts...).ToFunc()
}

// ByGameTypeID orders the results by the game_type_id field.
func ByGameTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameTypeID, opts...).ToFunc()
}

// ByGameGroupID orders the results by the game_group_id field.
func ByGameGroupID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameGroupID, opts...).ToFunc()
}

// BySportType orders the results by the sport_type field.
func BySportType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSportType, opts...).ToFunc()
}

// ByDtStarted orders the results by the dt_started field.
func ByDtStarted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDtStarted, opts...).ToFunc()
}

// ByDtCompleted orders the results by the dt_completed field.
func ByDtCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDtCompleted, opts...).ToFunc()
}

// ByWinTransactionID orders the results by the win_transaction_id field.
func ByWinTransactionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWinTransactionID, opts...).ToFunc()
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

// ByCreateTimeStr orders the results by the create_time_str field.
func ByCreateTimeStr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTimeStr, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
