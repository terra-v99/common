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

func newWinCoinLogCopy1(db *gorm.DB, opts ...gen.DOOption) winCoinLogCopy1 {
	_winCoinLogCopy1 := winCoinLogCopy1{}

	_winCoinLogCopy1.winCoinLogCopy1Do.UseDB(db, opts...)
	_winCoinLogCopy1.winCoinLogCopy1Do.UseModel(&model.WinCoinLogCopy1{})

	tableName := _winCoinLogCopy1.winCoinLogCopy1Do.TableName()
	_winCoinLogCopy1.ALL = field.NewAsterisk(tableName)
	_winCoinLogCopy1.ID = field.NewInt64(tableName, "id")
	_winCoinLogCopy1.UID = field.NewInt64(tableName, "uid")
	_winCoinLogCopy1.Username = field.NewString(tableName, "username")
	_winCoinLogCopy1.Category = field.NewInt64(tableName, "category")
	_winCoinLogCopy1.ReferID = field.NewInt64(tableName, "refer_id")
	_winCoinLogCopy1.Coin = field.NewField(tableName, "coin")
	_winCoinLogCopy1.CoinReal = field.NewField(tableName, "coin_real")
	_winCoinLogCopy1.PlatID = field.NewInt64(tableName, "plat_id")
	_winCoinLogCopy1.OutIn = field.NewInt64(tableName, "out_in")
	_winCoinLogCopy1.GameID = field.NewInt64(tableName, "game_id")
	_winCoinLogCopy1.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinLogCopy1.CoinAfter = field.NewField(tableName, "coin_after")
	_winCoinLogCopy1.Status = field.NewInt64(tableName, "status")
	_winCoinLogCopy1.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinLogCopy1.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winCoinLogCopy1.fillFieldMap()

	return _winCoinLogCopy1
}

type winCoinLogCopy1 struct {
	winCoinLogCopy1Do

	ALL        field.Asterisk
	ID         field.Int64
	UID        field.Int64
	Username   field.String
	Category   field.Int64
	ReferID    field.Int64
	Coin       field.Field
	CoinReal   field.Field
	PlatID     field.Int64
	OutIn      field.Int64
	GameID     field.Int64
	CoinBefore field.Field
	CoinAfter  field.Field
	Status     field.Int64
	CreatedAt  field.Int64
	UpdatedAt  field.Int64

	fieldMap map[string]field.Expr
}

func (w winCoinLogCopy1) Table(newTableName string) *winCoinLogCopy1 {
	w.winCoinLogCopy1Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinLogCopy1) As(alias string) *winCoinLogCopy1 {
	w.winCoinLogCopy1Do.DO = *(w.winCoinLogCopy1Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinLogCopy1) updateTableName(table string) *winCoinLogCopy1 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Category = field.NewInt64(table, "category")
	w.ReferID = field.NewInt64(table, "refer_id")
	w.Coin = field.NewField(table, "coin")
	w.CoinReal = field.NewField(table, "coin_real")
	w.PlatID = field.NewInt64(table, "plat_id")
	w.OutIn = field.NewInt64(table, "out_in")
	w.GameID = field.NewInt64(table, "game_id")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.CoinAfter = field.NewField(table, "coin_after")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinLogCopy1) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinLogCopy1) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 15)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["category"] = w.Category
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["coin_real"] = w.CoinReal
	w.fieldMap["plat_id"] = w.PlatID
	w.fieldMap["out_in"] = w.OutIn
	w.fieldMap["game_id"] = w.GameID
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["coin_after"] = w.CoinAfter
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinLogCopy1) clone(db *gorm.DB) winCoinLogCopy1 {
	w.winCoinLogCopy1Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinLogCopy1) replaceDB(db *gorm.DB) winCoinLogCopy1 {
	w.winCoinLogCopy1Do.ReplaceDB(db)
	return w
}

type winCoinLogCopy1Do struct{ gen.DO }

type IWinCoinLogCopy1Do interface {
	gen.SubQuery
	Debug() IWinCoinLogCopy1Do
	WithContext(ctx context.Context) IWinCoinLogCopy1Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinLogCopy1Do
	WriteDB() IWinCoinLogCopy1Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinLogCopy1Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinLogCopy1Do
	Not(conds ...gen.Condition) IWinCoinLogCopy1Do
	Or(conds ...gen.Condition) IWinCoinLogCopy1Do
	Select(conds ...field.Expr) IWinCoinLogCopy1Do
	Where(conds ...gen.Condition) IWinCoinLogCopy1Do
	Order(conds ...field.Expr) IWinCoinLogCopy1Do
	Distinct(cols ...field.Expr) IWinCoinLogCopy1Do
	Omit(cols ...field.Expr) IWinCoinLogCopy1Do
	Join(table schema.Tabler, on ...field.Expr) IWinCoinLogCopy1Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogCopy1Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogCopy1Do
	Group(cols ...field.Expr) IWinCoinLogCopy1Do
	Having(conds ...gen.Condition) IWinCoinLogCopy1Do
	Limit(limit int) IWinCoinLogCopy1Do
	Offset(offset int) IWinCoinLogCopy1Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinLogCopy1Do
	Unscoped() IWinCoinLogCopy1Do
	Create(values ...*model.WinCoinLogCopy1) error
	CreateInBatches(values []*model.WinCoinLogCopy1, batchSize int) error
	Save(values ...*model.WinCoinLogCopy1) error
	First() (*model.WinCoinLogCopy1, error)
	Take() (*model.WinCoinLogCopy1, error)
	Last() (*model.WinCoinLogCopy1, error)
	Find() ([]*model.WinCoinLogCopy1, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinLogCopy1, err error)
	FindInBatches(result *[]*model.WinCoinLogCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinLogCopy1) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinLogCopy1Do
	Assign(attrs ...field.AssignExpr) IWinCoinLogCopy1Do
	Joins(fields ...field.RelationField) IWinCoinLogCopy1Do
	Preload(fields ...field.RelationField) IWinCoinLogCopy1Do
	FirstOrInit() (*model.WinCoinLogCopy1, error)
	FirstOrCreate() (*model.WinCoinLogCopy1, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinLogCopy1, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinLogCopy1Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinLogCopy1Do) Debug() IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Debug())
}

func (w winCoinLogCopy1Do) WithContext(ctx context.Context) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinLogCopy1Do) ReadDB() IWinCoinLogCopy1Do {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinLogCopy1Do) WriteDB() IWinCoinLogCopy1Do {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinLogCopy1Do) Session(config *gorm.Session) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinLogCopy1Do) Clauses(conds ...clause.Expression) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinLogCopy1Do) Returning(value interface{}, columns ...string) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinLogCopy1Do) Not(conds ...gen.Condition) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinLogCopy1Do) Or(conds ...gen.Condition) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinLogCopy1Do) Select(conds ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinLogCopy1Do) Where(conds ...gen.Condition) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinLogCopy1Do) Order(conds ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinLogCopy1Do) Distinct(cols ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinLogCopy1Do) Omit(cols ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinLogCopy1Do) Join(table schema.Tabler, on ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinLogCopy1Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinLogCopy1Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinLogCopy1Do) Group(cols ...field.Expr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinLogCopy1Do) Having(conds ...gen.Condition) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinLogCopy1Do) Limit(limit int) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinLogCopy1Do) Offset(offset int) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinLogCopy1Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinLogCopy1Do) Unscoped() IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinLogCopy1Do) Create(values ...*model.WinCoinLogCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinLogCopy1Do) CreateInBatches(values []*model.WinCoinLogCopy1, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinLogCopy1Do) Save(values ...*model.WinCoinLogCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinLogCopy1Do) First() (*model.WinCoinLogCopy1, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLogCopy1), nil
	}
}

func (w winCoinLogCopy1Do) Take() (*model.WinCoinLogCopy1, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLogCopy1), nil
	}
}

func (w winCoinLogCopy1Do) Last() (*model.WinCoinLogCopy1, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLogCopy1), nil
	}
}

func (w winCoinLogCopy1Do) Find() ([]*model.WinCoinLogCopy1, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinLogCopy1), err
}

func (w winCoinLogCopy1Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinLogCopy1, err error) {
	buf := make([]*model.WinCoinLogCopy1, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinLogCopy1Do) FindInBatches(result *[]*model.WinCoinLogCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinLogCopy1Do) Attrs(attrs ...field.AssignExpr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinLogCopy1Do) Assign(attrs ...field.AssignExpr) IWinCoinLogCopy1Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinLogCopy1Do) Joins(fields ...field.RelationField) IWinCoinLogCopy1Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinLogCopy1Do) Preload(fields ...field.RelationField) IWinCoinLogCopy1Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinLogCopy1Do) FirstOrInit() (*model.WinCoinLogCopy1, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLogCopy1), nil
	}
}

func (w winCoinLogCopy1Do) FirstOrCreate() (*model.WinCoinLogCopy1, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLogCopy1), nil
	}
}

func (w winCoinLogCopy1Do) FindByPage(offset int, limit int) (result []*model.WinCoinLogCopy1, count int64, err error) {
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

func (w winCoinLogCopy1Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinLogCopy1Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinLogCopy1Do) Delete(models ...*model.WinCoinLogCopy1) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinLogCopy1Do) withDO(do gen.Dao) *winCoinLogCopy1Do {
	w.DO = *do.(*gen.DO)
	return w
}
