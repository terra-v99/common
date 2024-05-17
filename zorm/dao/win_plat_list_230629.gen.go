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

func newWinPlatList230629(db *gorm.DB, opts ...gen.DOOption) winPlatList230629 {
	_winPlatList230629 := winPlatList230629{}

	_winPlatList230629.winPlatList230629Do.UseDB(db, opts...)
	_winPlatList230629.winPlatList230629Do.UseModel(&model.WinPlatList230629{})

	tableName := _winPlatList230629.winPlatList230629Do.TableName()
	_winPlatList230629.ALL = field.NewAsterisk(tableName)
	_winPlatList230629.ID = field.NewInt64(tableName, "id")
	_winPlatList230629.Code = field.NewString(tableName, "code")
	_winPlatList230629.Name = field.NewString(tableName, "name")
	_winPlatList230629.Config = field.NewString(tableName, "config")
	_winPlatList230629.Rate = field.NewString(tableName, "rate")
	_winPlatList230629.Sort = field.NewInt64(tableName, "sort")
	_winPlatList230629.Status = field.NewInt64(tableName, "status")
	_winPlatList230629.CreatedAt = field.NewInt64(tableName, "created_at")
	_winPlatList230629.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winPlatList230629.fillFieldMap()

	return _winPlatList230629
}

type winPlatList230629 struct {
	winPlatList230629Do

	ALL       field.Asterisk
	ID        field.Int64
	Code      field.String // 平台编码
	Name      field.String // 平台名称
	Config    field.String // 配置信息
	Rate      field.String // 费率
	Sort      field.Int64  // 排序(从高到底、ID降序)
	Status    field.Int64  // 状态: 1-启用 0-停用
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winPlatList230629) Table(newTableName string) *winPlatList230629 {
	w.winPlatList230629Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winPlatList230629) As(alias string) *winPlatList230629 {
	w.winPlatList230629Do.DO = *(w.winPlatList230629Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winPlatList230629) updateTableName(table string) *winPlatList230629 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Code = field.NewString(table, "code")
	w.Name = field.NewString(table, "name")
	w.Config = field.NewString(table, "config")
	w.Rate = field.NewString(table, "rate")
	w.Sort = field.NewInt64(table, "sort")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winPlatList230629) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winPlatList230629) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["code"] = w.Code
	w.fieldMap["name"] = w.Name
	w.fieldMap["config"] = w.Config
	w.fieldMap["rate"] = w.Rate
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winPlatList230629) clone(db *gorm.DB) winPlatList230629 {
	w.winPlatList230629Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winPlatList230629) replaceDB(db *gorm.DB) winPlatList230629 {
	w.winPlatList230629Do.ReplaceDB(db)
	return w
}

type winPlatList230629Do struct{ gen.DO }

type IWinPlatList230629Do interface {
	gen.SubQuery
	Debug() IWinPlatList230629Do
	WithContext(ctx context.Context) IWinPlatList230629Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinPlatList230629Do
	WriteDB() IWinPlatList230629Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinPlatList230629Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinPlatList230629Do
	Not(conds ...gen.Condition) IWinPlatList230629Do
	Or(conds ...gen.Condition) IWinPlatList230629Do
	Select(conds ...field.Expr) IWinPlatList230629Do
	Where(conds ...gen.Condition) IWinPlatList230629Do
	Order(conds ...field.Expr) IWinPlatList230629Do
	Distinct(cols ...field.Expr) IWinPlatList230629Do
	Omit(cols ...field.Expr) IWinPlatList230629Do
	Join(table schema.Tabler, on ...field.Expr) IWinPlatList230629Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinPlatList230629Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinPlatList230629Do
	Group(cols ...field.Expr) IWinPlatList230629Do
	Having(conds ...gen.Condition) IWinPlatList230629Do
	Limit(limit int) IWinPlatList230629Do
	Offset(offset int) IWinPlatList230629Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPlatList230629Do
	Unscoped() IWinPlatList230629Do
	Create(values ...*model.WinPlatList230629) error
	CreateInBatches(values []*model.WinPlatList230629, batchSize int) error
	Save(values ...*model.WinPlatList230629) error
	First() (*model.WinPlatList230629, error)
	Take() (*model.WinPlatList230629, error)
	Last() (*model.WinPlatList230629, error)
	Find() ([]*model.WinPlatList230629, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPlatList230629, err error)
	FindInBatches(result *[]*model.WinPlatList230629, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinPlatList230629) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinPlatList230629Do
	Assign(attrs ...field.AssignExpr) IWinPlatList230629Do
	Joins(fields ...field.RelationField) IWinPlatList230629Do
	Preload(fields ...field.RelationField) IWinPlatList230629Do
	FirstOrInit() (*model.WinPlatList230629, error)
	FirstOrCreate() (*model.WinPlatList230629, error)
	FindByPage(offset int, limit int) (result []*model.WinPlatList230629, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinPlatList230629Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winPlatList230629Do) Debug() IWinPlatList230629Do {
	return w.withDO(w.DO.Debug())
}

func (w winPlatList230629Do) WithContext(ctx context.Context) IWinPlatList230629Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winPlatList230629Do) ReadDB() IWinPlatList230629Do {
	return w.Clauses(dbresolver.Read)
}

func (w winPlatList230629Do) WriteDB() IWinPlatList230629Do {
	return w.Clauses(dbresolver.Write)
}

func (w winPlatList230629Do) Session(config *gorm.Session) IWinPlatList230629Do {
	return w.withDO(w.DO.Session(config))
}

func (w winPlatList230629Do) Clauses(conds ...clause.Expression) IWinPlatList230629Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winPlatList230629Do) Returning(value interface{}, columns ...string) IWinPlatList230629Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winPlatList230629Do) Not(conds ...gen.Condition) IWinPlatList230629Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winPlatList230629Do) Or(conds ...gen.Condition) IWinPlatList230629Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winPlatList230629Do) Select(conds ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winPlatList230629Do) Where(conds ...gen.Condition) IWinPlatList230629Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winPlatList230629Do) Order(conds ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winPlatList230629Do) Distinct(cols ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winPlatList230629Do) Omit(cols ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winPlatList230629Do) Join(table schema.Tabler, on ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winPlatList230629Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winPlatList230629Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winPlatList230629Do) Group(cols ...field.Expr) IWinPlatList230629Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winPlatList230629Do) Having(conds ...gen.Condition) IWinPlatList230629Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winPlatList230629Do) Limit(limit int) IWinPlatList230629Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winPlatList230629Do) Offset(offset int) IWinPlatList230629Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winPlatList230629Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPlatList230629Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winPlatList230629Do) Unscoped() IWinPlatList230629Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winPlatList230629Do) Create(values ...*model.WinPlatList230629) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winPlatList230629Do) CreateInBatches(values []*model.WinPlatList230629, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winPlatList230629Do) Save(values ...*model.WinPlatList230629) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winPlatList230629Do) First() (*model.WinPlatList230629, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPlatList230629), nil
	}
}

func (w winPlatList230629Do) Take() (*model.WinPlatList230629, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPlatList230629), nil
	}
}

func (w winPlatList230629Do) Last() (*model.WinPlatList230629, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPlatList230629), nil
	}
}

func (w winPlatList230629Do) Find() ([]*model.WinPlatList230629, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinPlatList230629), err
}

func (w winPlatList230629Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPlatList230629, err error) {
	buf := make([]*model.WinPlatList230629, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winPlatList230629Do) FindInBatches(result *[]*model.WinPlatList230629, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winPlatList230629Do) Attrs(attrs ...field.AssignExpr) IWinPlatList230629Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winPlatList230629Do) Assign(attrs ...field.AssignExpr) IWinPlatList230629Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winPlatList230629Do) Joins(fields ...field.RelationField) IWinPlatList230629Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winPlatList230629Do) Preload(fields ...field.RelationField) IWinPlatList230629Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winPlatList230629Do) FirstOrInit() (*model.WinPlatList230629, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPlatList230629), nil
	}
}

func (w winPlatList230629Do) FirstOrCreate() (*model.WinPlatList230629, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPlatList230629), nil
	}
}

func (w winPlatList230629Do) FindByPage(offset int, limit int) (result []*model.WinPlatList230629, count int64, err error) {
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

func (w winPlatList230629Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winPlatList230629Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winPlatList230629Do) Delete(models ...*model.WinPlatList230629) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winPlatList230629Do) withDO(do gen.Dao) *winPlatList230629Do {
	w.DO = *do.(*gen.DO)
	return w
}
