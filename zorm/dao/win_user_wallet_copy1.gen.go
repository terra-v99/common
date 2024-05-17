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

func newWinUserWalletCopy1(db *gorm.DB, opts ...gen.DOOption) winUserWalletCopy1 {
	_winUserWalletCopy1 := winUserWalletCopy1{}

	_winUserWalletCopy1.winUserWalletCopy1Do.UseDB(db, opts...)
	_winUserWalletCopy1.winUserWalletCopy1Do.UseModel(&model.WinUserWalletCopy1{})

	tableName := _winUserWalletCopy1.winUserWalletCopy1Do.TableName()
	_winUserWalletCopy1.ALL = field.NewAsterisk(tableName)
	_winUserWalletCopy1.ID = field.NewInt64(tableName, "id")
	_winUserWalletCopy1.Username = field.NewString(tableName, "username")
	_winUserWalletCopy1.Coin = field.NewField(tableName, "coin")
	_winUserWalletCopy1.Version = field.NewInt64(tableName, "version")
	_winUserWalletCopy1.ModifyAt = field.NewInt64(tableName, "modify_at")
	_winUserWalletCopy1.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserWalletCopy1.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winUserWalletCopy1.fillFieldMap()

	return _winUserWalletCopy1
}

type winUserWalletCopy1 struct {
	winUserWalletCopy1Do

	ALL       field.Asterisk
	ID        field.Int64
	Username  field.String // 用户名
	Coin      field.Field  // 账户余额
	Version   field.Int64  // 版本号
	ModifyAt  field.Int64  // 13位时间戳
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winUserWalletCopy1) Table(newTableName string) *winUserWalletCopy1 {
	w.winUserWalletCopy1Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserWalletCopy1) As(alias string) *winUserWalletCopy1 {
	w.winUserWalletCopy1Do.DO = *(w.winUserWalletCopy1Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserWalletCopy1) updateTableName(table string) *winUserWalletCopy1 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Username = field.NewString(table, "username")
	w.Coin = field.NewField(table, "coin")
	w.Version = field.NewInt64(table, "version")
	w.ModifyAt = field.NewInt64(table, "modify_at")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winUserWalletCopy1) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserWalletCopy1) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.ID
	w.fieldMap["username"] = w.Username
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["version"] = w.Version
	w.fieldMap["modify_at"] = w.ModifyAt
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winUserWalletCopy1) clone(db *gorm.DB) winUserWalletCopy1 {
	w.winUserWalletCopy1Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserWalletCopy1) replaceDB(db *gorm.DB) winUserWalletCopy1 {
	w.winUserWalletCopy1Do.ReplaceDB(db)
	return w
}

type winUserWalletCopy1Do struct{ gen.DO }

type IWinUserWalletCopy1Do interface {
	gen.SubQuery
	Debug() IWinUserWalletCopy1Do
	WithContext(ctx context.Context) IWinUserWalletCopy1Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserWalletCopy1Do
	WriteDB() IWinUserWalletCopy1Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserWalletCopy1Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserWalletCopy1Do
	Not(conds ...gen.Condition) IWinUserWalletCopy1Do
	Or(conds ...gen.Condition) IWinUserWalletCopy1Do
	Select(conds ...field.Expr) IWinUserWalletCopy1Do
	Where(conds ...gen.Condition) IWinUserWalletCopy1Do
	Order(conds ...field.Expr) IWinUserWalletCopy1Do
	Distinct(cols ...field.Expr) IWinUserWalletCopy1Do
	Omit(cols ...field.Expr) IWinUserWalletCopy1Do
	Join(table schema.Tabler, on ...field.Expr) IWinUserWalletCopy1Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserWalletCopy1Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserWalletCopy1Do
	Group(cols ...field.Expr) IWinUserWalletCopy1Do
	Having(conds ...gen.Condition) IWinUserWalletCopy1Do
	Limit(limit int) IWinUserWalletCopy1Do
	Offset(offset int) IWinUserWalletCopy1Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserWalletCopy1Do
	Unscoped() IWinUserWalletCopy1Do
	Create(values ...*model.WinUserWalletCopy1) error
	CreateInBatches(values []*model.WinUserWalletCopy1, batchSize int) error
	Save(values ...*model.WinUserWalletCopy1) error
	First() (*model.WinUserWalletCopy1, error)
	Take() (*model.WinUserWalletCopy1, error)
	Last() (*model.WinUserWalletCopy1, error)
	Find() ([]*model.WinUserWalletCopy1, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserWalletCopy1, err error)
	FindInBatches(result *[]*model.WinUserWalletCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserWalletCopy1) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserWalletCopy1Do
	Assign(attrs ...field.AssignExpr) IWinUserWalletCopy1Do
	Joins(fields ...field.RelationField) IWinUserWalletCopy1Do
	Preload(fields ...field.RelationField) IWinUserWalletCopy1Do
	FirstOrInit() (*model.WinUserWalletCopy1, error)
	FirstOrCreate() (*model.WinUserWalletCopy1, error)
	FindByPage(offset int, limit int) (result []*model.WinUserWalletCopy1, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserWalletCopy1Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserWalletCopy1Do) Debug() IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Debug())
}

func (w winUserWalletCopy1Do) WithContext(ctx context.Context) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserWalletCopy1Do) ReadDB() IWinUserWalletCopy1Do {
	return w.Clauses(dbresolver.Read)
}

func (w winUserWalletCopy1Do) WriteDB() IWinUserWalletCopy1Do {
	return w.Clauses(dbresolver.Write)
}

func (w winUserWalletCopy1Do) Session(config *gorm.Session) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Session(config))
}

func (w winUserWalletCopy1Do) Clauses(conds ...clause.Expression) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserWalletCopy1Do) Returning(value interface{}, columns ...string) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserWalletCopy1Do) Not(conds ...gen.Condition) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserWalletCopy1Do) Or(conds ...gen.Condition) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserWalletCopy1Do) Select(conds ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserWalletCopy1Do) Where(conds ...gen.Condition) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserWalletCopy1Do) Order(conds ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserWalletCopy1Do) Distinct(cols ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserWalletCopy1Do) Omit(cols ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserWalletCopy1Do) Join(table schema.Tabler, on ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserWalletCopy1Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserWalletCopy1Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserWalletCopy1Do) Group(cols ...field.Expr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserWalletCopy1Do) Having(conds ...gen.Condition) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserWalletCopy1Do) Limit(limit int) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserWalletCopy1Do) Offset(offset int) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserWalletCopy1Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserWalletCopy1Do) Unscoped() IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserWalletCopy1Do) Create(values ...*model.WinUserWalletCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserWalletCopy1Do) CreateInBatches(values []*model.WinUserWalletCopy1, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserWalletCopy1Do) Save(values ...*model.WinUserWalletCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserWalletCopy1Do) First() (*model.WinUserWalletCopy1, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWalletCopy1), nil
	}
}

func (w winUserWalletCopy1Do) Take() (*model.WinUserWalletCopy1, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWalletCopy1), nil
	}
}

func (w winUserWalletCopy1Do) Last() (*model.WinUserWalletCopy1, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWalletCopy1), nil
	}
}

func (w winUserWalletCopy1Do) Find() ([]*model.WinUserWalletCopy1, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserWalletCopy1), err
}

func (w winUserWalletCopy1Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserWalletCopy1, err error) {
	buf := make([]*model.WinUserWalletCopy1, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserWalletCopy1Do) FindInBatches(result *[]*model.WinUserWalletCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserWalletCopy1Do) Attrs(attrs ...field.AssignExpr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserWalletCopy1Do) Assign(attrs ...field.AssignExpr) IWinUserWalletCopy1Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserWalletCopy1Do) Joins(fields ...field.RelationField) IWinUserWalletCopy1Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserWalletCopy1Do) Preload(fields ...field.RelationField) IWinUserWalletCopy1Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserWalletCopy1Do) FirstOrInit() (*model.WinUserWalletCopy1, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWalletCopy1), nil
	}
}

func (w winUserWalletCopy1Do) FirstOrCreate() (*model.WinUserWalletCopy1, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserWalletCopy1), nil
	}
}

func (w winUserWalletCopy1Do) FindByPage(offset int, limit int) (result []*model.WinUserWalletCopy1, count int64, err error) {
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

func (w winUserWalletCopy1Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserWalletCopy1Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserWalletCopy1Do) Delete(models ...*model.WinUserWalletCopy1) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserWalletCopy1Do) withDO(do gen.Dao) *winUserWalletCopy1Do {
	w.DO = *do.(*gen.DO)
	return w
}
