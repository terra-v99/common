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

	"gitlab.skig.tech/zero-core/common/zorm/model"
)

func newWinDict(db *gorm.DB, opts ...gen.DOOption) winDict {
	_winDict := winDict{}

	_winDict.winDictDo.UseDB(db, opts...)
	_winDict.winDictDo.UseModel(&model.WinDict{})

	tableName := _winDict.winDictDo.TableName()
	_winDict.ALL = field.NewAsterisk(tableName)
	_winDict.ID = field.NewInt64(tableName, "id")
	_winDict.Title = field.NewString(tableName, "title")
	_winDict.Category = field.NewString(tableName, "category")
	_winDict.Status = field.NewInt64(tableName, "status")
	_winDict.CreatedAt = field.NewInt64(tableName, "created_at")
	_winDict.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winDict.fillFieldMap()

	return _winDict
}

type winDict struct {
	winDictDo

	ALL       field.Asterisk
	ID        field.Int64
	Title     field.String // 名称
	Category  field.String // 种类
	Status    field.Int64  // 状态:1-启用 0-禁用
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winDict) Table(newTableName string) *winDict {
	w.winDictDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winDict) As(alias string) *winDict {
	w.winDictDo.DO = *(w.winDictDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winDict) updateTableName(table string) *winDict {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Title = field.NewString(table, "title")
	w.Category = field.NewString(table, "category")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winDict) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winDict) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["title"] = w.Title
	w.fieldMap["category"] = w.Category
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winDict) clone(db *gorm.DB) winDict {
	w.winDictDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winDict) replaceDB(db *gorm.DB) winDict {
	w.winDictDo.ReplaceDB(db)
	return w
}

type winDictDo struct{ gen.DO }

type IWinDictDo interface {
	gen.SubQuery
	Debug() IWinDictDo
	WithContext(ctx context.Context) IWinDictDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinDictDo
	WriteDB() IWinDictDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinDictDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinDictDo
	Not(conds ...gen.Condition) IWinDictDo
	Or(conds ...gen.Condition) IWinDictDo
	Select(conds ...field.Expr) IWinDictDo
	Where(conds ...gen.Condition) IWinDictDo
	Order(conds ...field.Expr) IWinDictDo
	Distinct(cols ...field.Expr) IWinDictDo
	Omit(cols ...field.Expr) IWinDictDo
	Join(table schema.Tabler, on ...field.Expr) IWinDictDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinDictDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinDictDo
	Group(cols ...field.Expr) IWinDictDo
	Having(conds ...gen.Condition) IWinDictDo
	Limit(limit int) IWinDictDo
	Offset(offset int) IWinDictDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinDictDo
	Unscoped() IWinDictDo
	Create(values ...*model.WinDict) error
	CreateInBatches(values []*model.WinDict, batchSize int) error
	Save(values ...*model.WinDict) error
	First() (*model.WinDict, error)
	Take() (*model.WinDict, error)
	Last() (*model.WinDict, error)
	Find() ([]*model.WinDict, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinDict, err error)
	FindInBatches(result *[]*model.WinDict, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinDict) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinDictDo
	Assign(attrs ...field.AssignExpr) IWinDictDo
	Joins(fields ...field.RelationField) IWinDictDo
	Preload(fields ...field.RelationField) IWinDictDo
	FirstOrInit() (*model.WinDict, error)
	FirstOrCreate() (*model.WinDict, error)
	FindByPage(offset int, limit int) (result []*model.WinDict, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinDictDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winDictDo) Debug() IWinDictDo {
	return w.withDO(w.DO.Debug())
}

func (w winDictDo) WithContext(ctx context.Context) IWinDictDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winDictDo) ReadDB() IWinDictDo {
	return w.Clauses(dbresolver.Read)
}

func (w winDictDo) WriteDB() IWinDictDo {
	return w.Clauses(dbresolver.Write)
}

func (w winDictDo) Session(config *gorm.Session) IWinDictDo {
	return w.withDO(w.DO.Session(config))
}

func (w winDictDo) Clauses(conds ...clause.Expression) IWinDictDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winDictDo) Returning(value interface{}, columns ...string) IWinDictDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winDictDo) Not(conds ...gen.Condition) IWinDictDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winDictDo) Or(conds ...gen.Condition) IWinDictDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winDictDo) Select(conds ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winDictDo) Where(conds ...gen.Condition) IWinDictDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winDictDo) Order(conds ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winDictDo) Distinct(cols ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winDictDo) Omit(cols ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winDictDo) Join(table schema.Tabler, on ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winDictDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winDictDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winDictDo) Group(cols ...field.Expr) IWinDictDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winDictDo) Having(conds ...gen.Condition) IWinDictDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winDictDo) Limit(limit int) IWinDictDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winDictDo) Offset(offset int) IWinDictDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winDictDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinDictDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winDictDo) Unscoped() IWinDictDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winDictDo) Create(values ...*model.WinDict) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winDictDo) CreateInBatches(values []*model.WinDict, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winDictDo) Save(values ...*model.WinDict) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winDictDo) First() (*model.WinDict, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDict), nil
	}
}

func (w winDictDo) Take() (*model.WinDict, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDict), nil
	}
}

func (w winDictDo) Last() (*model.WinDict, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDict), nil
	}
}

func (w winDictDo) Find() ([]*model.WinDict, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinDict), err
}

func (w winDictDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinDict, err error) {
	buf := make([]*model.WinDict, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winDictDo) FindInBatches(result *[]*model.WinDict, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winDictDo) Attrs(attrs ...field.AssignExpr) IWinDictDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winDictDo) Assign(attrs ...field.AssignExpr) IWinDictDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winDictDo) Joins(fields ...field.RelationField) IWinDictDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winDictDo) Preload(fields ...field.RelationField) IWinDictDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winDictDo) FirstOrInit() (*model.WinDict, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDict), nil
	}
}

func (w winDictDo) FirstOrCreate() (*model.WinDict, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDict), nil
	}
}

func (w winDictDo) FindByPage(offset int, limit int) (result []*model.WinDict, count int64, err error) {
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

func (w winDictDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winDictDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winDictDo) Delete(models ...*model.WinDict) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winDictDo) withDO(do gen.Dao) *winDictDo {
	w.DO = *do.(*gen.DO)
	return w
}
