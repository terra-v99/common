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

func newXxlJobLock(db *gorm.DB, opts ...gen.DOOption) xxlJobLock {
	_xxlJobLock := xxlJobLock{}

	_xxlJobLock.xxlJobLockDo.UseDB(db, opts...)
	_xxlJobLock.xxlJobLockDo.UseModel(&model.XxlJobLock{})

	tableName := _xxlJobLock.xxlJobLockDo.TableName()
	_xxlJobLock.ALL = field.NewAsterisk(tableName)
	_xxlJobLock.LockName = field.NewString(tableName, "lock_name")

	_xxlJobLock.fillFieldMap()

	return _xxlJobLock
}

type xxlJobLock struct {
	xxlJobLockDo

	ALL      field.Asterisk
	LockName field.String // 锁名称

	fieldMap map[string]field.Expr
}

func (x xxlJobLock) Table(newTableName string) *xxlJobLock {
	x.xxlJobLockDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xxlJobLock) As(alias string) *xxlJobLock {
	x.xxlJobLockDo.DO = *(x.xxlJobLockDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xxlJobLock) updateTableName(table string) *xxlJobLock {
	x.ALL = field.NewAsterisk(table)
	x.LockName = field.NewString(table, "lock_name")

	x.fillFieldMap()

	return x
}

func (x *xxlJobLock) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xxlJobLock) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 1)
	x.fieldMap["lock_name"] = x.LockName
}

func (x xxlJobLock) clone(db *gorm.DB) xxlJobLock {
	x.xxlJobLockDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xxlJobLock) replaceDB(db *gorm.DB) xxlJobLock {
	x.xxlJobLockDo.ReplaceDB(db)
	return x
}

type xxlJobLockDo struct{ gen.DO }

type IXxlJobLockDo interface {
	gen.SubQuery
	Debug() IXxlJobLockDo
	WithContext(ctx context.Context) IXxlJobLockDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXxlJobLockDo
	WriteDB() IXxlJobLockDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXxlJobLockDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXxlJobLockDo
	Not(conds ...gen.Condition) IXxlJobLockDo
	Or(conds ...gen.Condition) IXxlJobLockDo
	Select(conds ...field.Expr) IXxlJobLockDo
	Where(conds ...gen.Condition) IXxlJobLockDo
	Order(conds ...field.Expr) IXxlJobLockDo
	Distinct(cols ...field.Expr) IXxlJobLockDo
	Omit(cols ...field.Expr) IXxlJobLockDo
	Join(table schema.Tabler, on ...field.Expr) IXxlJobLockDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobLockDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobLockDo
	Group(cols ...field.Expr) IXxlJobLockDo
	Having(conds ...gen.Condition) IXxlJobLockDo
	Limit(limit int) IXxlJobLockDo
	Offset(offset int) IXxlJobLockDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobLockDo
	Unscoped() IXxlJobLockDo
	Create(values ...*model.XxlJobLock) error
	CreateInBatches(values []*model.XxlJobLock, batchSize int) error
	Save(values ...*model.XxlJobLock) error
	First() (*model.XxlJobLock, error)
	Take() (*model.XxlJobLock, error)
	Last() (*model.XxlJobLock, error)
	Find() ([]*model.XxlJobLock, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobLock, err error)
	FindInBatches(result *[]*model.XxlJobLock, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XxlJobLock) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXxlJobLockDo
	Assign(attrs ...field.AssignExpr) IXxlJobLockDo
	Joins(fields ...field.RelationField) IXxlJobLockDo
	Preload(fields ...field.RelationField) IXxlJobLockDo
	FirstOrInit() (*model.XxlJobLock, error)
	FirstOrCreate() (*model.XxlJobLock, error)
	FindByPage(offset int, limit int) (result []*model.XxlJobLock, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXxlJobLockDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xxlJobLockDo) Debug() IXxlJobLockDo {
	return x.withDO(x.DO.Debug())
}

func (x xxlJobLockDo) WithContext(ctx context.Context) IXxlJobLockDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xxlJobLockDo) ReadDB() IXxlJobLockDo {
	return x.Clauses(dbresolver.Read)
}

func (x xxlJobLockDo) WriteDB() IXxlJobLockDo {
	return x.Clauses(dbresolver.Write)
}

func (x xxlJobLockDo) Session(config *gorm.Session) IXxlJobLockDo {
	return x.withDO(x.DO.Session(config))
}

func (x xxlJobLockDo) Clauses(conds ...clause.Expression) IXxlJobLockDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xxlJobLockDo) Returning(value interface{}, columns ...string) IXxlJobLockDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xxlJobLockDo) Not(conds ...gen.Condition) IXxlJobLockDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xxlJobLockDo) Or(conds ...gen.Condition) IXxlJobLockDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xxlJobLockDo) Select(conds ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xxlJobLockDo) Where(conds ...gen.Condition) IXxlJobLockDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xxlJobLockDo) Order(conds ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xxlJobLockDo) Distinct(cols ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xxlJobLockDo) Omit(cols ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xxlJobLockDo) Join(table schema.Tabler, on ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xxlJobLockDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xxlJobLockDo) RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xxlJobLockDo) Group(cols ...field.Expr) IXxlJobLockDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xxlJobLockDo) Having(conds ...gen.Condition) IXxlJobLockDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xxlJobLockDo) Limit(limit int) IXxlJobLockDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xxlJobLockDo) Offset(offset int) IXxlJobLockDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xxlJobLockDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobLockDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xxlJobLockDo) Unscoped() IXxlJobLockDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xxlJobLockDo) Create(values ...*model.XxlJobLock) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xxlJobLockDo) CreateInBatches(values []*model.XxlJobLock, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xxlJobLockDo) Save(values ...*model.XxlJobLock) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xxlJobLockDo) First() (*model.XxlJobLock, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLock), nil
	}
}

func (x xxlJobLockDo) Take() (*model.XxlJobLock, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLock), nil
	}
}

func (x xxlJobLockDo) Last() (*model.XxlJobLock, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLock), nil
	}
}

func (x xxlJobLockDo) Find() ([]*model.XxlJobLock, error) {
	result, err := x.DO.Find()
	return result.([]*model.XxlJobLock), err
}

func (x xxlJobLockDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobLock, err error) {
	buf := make([]*model.XxlJobLock, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xxlJobLockDo) FindInBatches(result *[]*model.XxlJobLock, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xxlJobLockDo) Attrs(attrs ...field.AssignExpr) IXxlJobLockDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xxlJobLockDo) Assign(attrs ...field.AssignExpr) IXxlJobLockDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xxlJobLockDo) Joins(fields ...field.RelationField) IXxlJobLockDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xxlJobLockDo) Preload(fields ...field.RelationField) IXxlJobLockDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xxlJobLockDo) FirstOrInit() (*model.XxlJobLock, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLock), nil
	}
}

func (x xxlJobLockDo) FirstOrCreate() (*model.XxlJobLock, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLock), nil
	}
}

func (x xxlJobLockDo) FindByPage(offset int, limit int) (result []*model.XxlJobLock, count int64, err error) {
	result, err = x.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = x.Offset(-1).Limit(-1).Count()
	return
}

func (x xxlJobLockDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xxlJobLockDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xxlJobLockDo) Delete(models ...*model.XxlJobLock) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xxlJobLockDo) withDO(do gen.Dao) *xxlJobLockDo {
	x.DO = *do.(*gen.DO)
	return x
}
