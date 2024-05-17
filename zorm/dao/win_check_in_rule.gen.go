// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/terra-v99/common/zorm/model"
)

func newWinCheckInRule(db *gorm.DB, opts ...gen.DOOption) winCheckInRule {
	_winCheckInRule := winCheckInRule{}

	_winCheckInRule.winCheckInRuleDo.UseDB(db, opts...)
	_winCheckInRule.winCheckInRuleDo.UseModel(&model.WinCheckInRule{})

	tableName := _winCheckInRule.winCheckInRuleDo.TableName()
	_winCheckInRule.ALL = field.NewAsterisk(tableName)
	_winCheckInRule.ID = field.NewInt64(tableName, "id")
	_winCheckInRule.PromotionID = field.NewInt64(tableName, "promotion_id")
	_winCheckInRule.Days = field.NewInt64(tableName, "days")
	_winCheckInRule.IconImg = field.NewString(tableName, "icon_img")
	_winCheckInRule.AwardAmount = field.NewField(tableName, "award_amount")
	_winCheckInRule.AuditTimes = field.NewInt64(tableName, "audit_times")
	_winCheckInRule.CreatedAt = field.NewInt64(tableName, "created_at")

	_winCheckInRule.fillFieldMap()

	return _winCheckInRule
}

type winCheckInRule struct {
	winCheckInRuleDo

	ALL         field.Asterisk
	ID          field.Int64  // 主键
	PromotionID field.Int64  // 签到活动ID
	Days        field.Int64  // 连续签到天数
	IconImg     field.String // 签到图标
	AwardAmount field.Field  // 签到奖励
	AuditTimes  field.Int64  // 稽核倍数
	CreatedAt   field.Int64  // 创建时间

	fieldMap map[string]field.Expr
}

func (w winCheckInRule) Table(newTableName string) *winCheckInRule {
	w.winCheckInRuleDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCheckInRule) As(alias string) *winCheckInRule {
	w.winCheckInRuleDo.DO = *(w.winCheckInRuleDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCheckInRule) updateTableName(table string) *winCheckInRule {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.PromotionID = field.NewInt64(table, "promotion_id")
	w.Days = field.NewInt64(table, "days")
	w.IconImg = field.NewString(table, "icon_img")
	w.AwardAmount = field.NewField(table, "award_amount")
	w.AuditTimes = field.NewInt64(table, "audit_times")
	w.CreatedAt = field.NewInt64(table, "created_at")

	w.fillFieldMap()

	return w
}

func (w *winCheckInRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCheckInRule) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.ID
	w.fieldMap["promotion_id"] = w.PromotionID
	w.fieldMap["days"] = w.Days
	w.fieldMap["icon_img"] = w.IconImg
	w.fieldMap["award_amount"] = w.AwardAmount
	w.fieldMap["audit_times"] = w.AuditTimes
	w.fieldMap["created_at"] = w.CreatedAt
}

func (w winCheckInRule) clone(db *gorm.DB) winCheckInRule {
	w.winCheckInRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCheckInRule) replaceDB(db *gorm.DB) winCheckInRule {
	w.winCheckInRuleDo.ReplaceDB(db)
	return w
}

type winCheckInRuleDo struct{ gen.DO }

type IWinCheckInRuleDo interface {
	gen.SubQuery
	Debug() IWinCheckInRuleDo
	WithContext(ctx context.Context) IWinCheckInRuleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCheckInRuleDo
	WriteDB() IWinCheckInRuleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCheckInRuleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCheckInRuleDo
	Not(conds ...gen.Condition) IWinCheckInRuleDo
	Or(conds ...gen.Condition) IWinCheckInRuleDo
	Select(conds ...field.Expr) IWinCheckInRuleDo
	Where(conds ...gen.Condition) IWinCheckInRuleDo
	Order(conds ...field.Expr) IWinCheckInRuleDo
	Distinct(cols ...field.Expr) IWinCheckInRuleDo
	Omit(cols ...field.Expr) IWinCheckInRuleDo
	Join(table schema.Tabler, on ...field.Expr) IWinCheckInRuleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCheckInRuleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCheckInRuleDo
	Group(cols ...field.Expr) IWinCheckInRuleDo
	Having(conds ...gen.Condition) IWinCheckInRuleDo
	Limit(limit int) IWinCheckInRuleDo
	Offset(offset int) IWinCheckInRuleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCheckInRuleDo
	Unscoped() IWinCheckInRuleDo
	Create(values ...*model.WinCheckInRule) error
	CreateInBatches(values []*model.WinCheckInRule, batchSize int) error
	Save(values ...*model.WinCheckInRule) error
	First() (*model.WinCheckInRule, error)
	Take() (*model.WinCheckInRule, error)
	Last() (*model.WinCheckInRule, error)
	Find() ([]*model.WinCheckInRule, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCheckInRule, err error)
	FindInBatches(result *[]*model.WinCheckInRule, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCheckInRule) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCheckInRuleDo
	Assign(attrs ...field.AssignExpr) IWinCheckInRuleDo
	Joins(fields ...field.RelationField) IWinCheckInRuleDo
	Preload(fields ...field.RelationField) IWinCheckInRuleDo
	FirstOrInit() (*model.WinCheckInRule, error)
	FirstOrCreate() (*model.WinCheckInRule, error)
	FindByPage(offset int, limit int) (result []*model.WinCheckInRule, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCheckInRuleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCheckInRuleDo) Debug() IWinCheckInRuleDo {
	return w.withDO(w.DO.Debug())
}

func (w winCheckInRuleDo) WithContext(ctx context.Context) IWinCheckInRuleDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCheckInRuleDo) ReadDB() IWinCheckInRuleDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCheckInRuleDo) WriteDB() IWinCheckInRuleDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCheckInRuleDo) Session(config *gorm.Session) IWinCheckInRuleDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCheckInRuleDo) Clauses(conds ...clause.Expression) IWinCheckInRuleDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCheckInRuleDo) Returning(value interface{}, columns ...string) IWinCheckInRuleDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCheckInRuleDo) Not(conds ...gen.Condition) IWinCheckInRuleDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCheckInRuleDo) Or(conds ...gen.Condition) IWinCheckInRuleDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCheckInRuleDo) Select(conds ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCheckInRuleDo) Where(conds ...gen.Condition) IWinCheckInRuleDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCheckInRuleDo) Order(conds ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCheckInRuleDo) Distinct(cols ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCheckInRuleDo) Omit(cols ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCheckInRuleDo) Join(table schema.Tabler, on ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCheckInRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCheckInRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCheckInRuleDo) Group(cols ...field.Expr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCheckInRuleDo) Having(conds ...gen.Condition) IWinCheckInRuleDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCheckInRuleDo) Limit(limit int) IWinCheckInRuleDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCheckInRuleDo) Offset(offset int) IWinCheckInRuleDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCheckInRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCheckInRuleDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCheckInRuleDo) Unscoped() IWinCheckInRuleDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCheckInRuleDo) Create(values ...*model.WinCheckInRule) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCheckInRuleDo) CreateInBatches(values []*model.WinCheckInRule, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCheckInRuleDo) Save(values ...*model.WinCheckInRule) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCheckInRuleDo) First() (*model.WinCheckInRule, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCheckInRule), nil
	}
}

func (w winCheckInRuleDo) Take() (*model.WinCheckInRule, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCheckInRule), nil
	}
}

func (w winCheckInRuleDo) Last() (*model.WinCheckInRule, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCheckInRule), nil
	}
}

func (w winCheckInRuleDo) Find() ([]*model.WinCheckInRule, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCheckInRule), err
}

func (w winCheckInRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCheckInRule, err error) {
	buf := make([]*model.WinCheckInRule, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCheckInRuleDo) FindInBatches(result *[]*model.WinCheckInRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCheckInRuleDo) Attrs(attrs ...field.AssignExpr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCheckInRuleDo) Assign(attrs ...field.AssignExpr) IWinCheckInRuleDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCheckInRuleDo) Joins(fields ...field.RelationField) IWinCheckInRuleDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCheckInRuleDo) Preload(fields ...field.RelationField) IWinCheckInRuleDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCheckInRuleDo) FirstOrInit() (*model.WinCheckInRule, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCheckInRule), nil
	}
}

func (w winCheckInRuleDo) FirstOrCreate() (*model.WinCheckInRule, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCheckInRule), nil
	}
}

func (w winCheckInRuleDo) FindByPage(offset int, limit int) (result []*model.WinCheckInRule, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w winCheckInRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCheckInRuleDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCheckInRuleDo) Delete(models ...*model.WinCheckInRule) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCheckInRuleDo) withDO(do gen.Dao) *winCheckInRuleDo {
	w.DO = *do.(*gen.DO)
	return w
}
