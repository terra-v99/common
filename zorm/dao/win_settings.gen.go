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

func newWinSetting(db *gorm.DB, opts ...gen.DOOption) winSetting {
	_winSetting := winSetting{}

	_winSetting.winSettingDo.UseDB(db, opts...)
	_winSetting.winSettingDo.UseModel(&model.WinSetting{})

	tableName := _winSetting.winSettingDo.TableName()
	_winSetting.ALL = field.NewAsterisk(tableName)
	_winSetting.Code = field.NewString(tableName, "code")
	_winSetting.Value = field.NewString(tableName, "value")
	_winSetting.CreatedAt = field.NewInt64(tableName, "created_at")
	_winSetting.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winSetting.fillFieldMap()

	return _winSetting
}

type winSetting struct {
	winSettingDo

	ALL       field.Asterisk
	Code      field.String // code
	Value     field.String // 字符串或json
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winSetting) Table(newTableName string) *winSetting {
	w.winSettingDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winSetting) As(alias string) *winSetting {
	w.winSettingDo.DO = *(w.winSettingDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winSetting) updateTableName(table string) *winSetting {
	w.ALL = field.NewAsterisk(table)
	w.Code = field.NewString(table, "code")
	w.Value = field.NewString(table, "value")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winSetting) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 4)
	w.fieldMap["code"] = w.Code
	w.fieldMap["value"] = w.Value
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winSetting) clone(db *gorm.DB) winSetting {
	w.winSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winSetting) replaceDB(db *gorm.DB) winSetting {
	w.winSettingDo.ReplaceDB(db)
	return w
}

type winSettingDo struct{ gen.DO }

type IWinSettingDo interface {
	gen.SubQuery
	Debug() IWinSettingDo
	WithContext(ctx context.Context) IWinSettingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinSettingDo
	WriteDB() IWinSettingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinSettingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinSettingDo
	Not(conds ...gen.Condition) IWinSettingDo
	Or(conds ...gen.Condition) IWinSettingDo
	Select(conds ...field.Expr) IWinSettingDo
	Where(conds ...gen.Condition) IWinSettingDo
	Order(conds ...field.Expr) IWinSettingDo
	Distinct(cols ...field.Expr) IWinSettingDo
	Omit(cols ...field.Expr) IWinSettingDo
	Join(table schema.Tabler, on ...field.Expr) IWinSettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinSettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinSettingDo
	Group(cols ...field.Expr) IWinSettingDo
	Having(conds ...gen.Condition) IWinSettingDo
	Limit(limit int) IWinSettingDo
	Offset(offset int) IWinSettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinSettingDo
	Unscoped() IWinSettingDo
	Create(values ...*model.WinSetting) error
	CreateInBatches(values []*model.WinSetting, batchSize int) error
	Save(values ...*model.WinSetting) error
	First() (*model.WinSetting, error)
	Take() (*model.WinSetting, error)
	Last() (*model.WinSetting, error)
	Find() ([]*model.WinSetting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinSetting, err error)
	FindInBatches(result *[]*model.WinSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinSetting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinSettingDo
	Assign(attrs ...field.AssignExpr) IWinSettingDo
	Joins(fields ...field.RelationField) IWinSettingDo
	Preload(fields ...field.RelationField) IWinSettingDo
	FirstOrInit() (*model.WinSetting, error)
	FirstOrCreate() (*model.WinSetting, error)
	FindByPage(offset int, limit int) (result []*model.WinSetting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinSettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winSettingDo) Debug() IWinSettingDo {
	return w.withDO(w.DO.Debug())
}

func (w winSettingDo) WithContext(ctx context.Context) IWinSettingDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winSettingDo) ReadDB() IWinSettingDo {
	return w.Clauses(dbresolver.Read)
}

func (w winSettingDo) WriteDB() IWinSettingDo {
	return w.Clauses(dbresolver.Write)
}

func (w winSettingDo) Session(config *gorm.Session) IWinSettingDo {
	return w.withDO(w.DO.Session(config))
}

func (w winSettingDo) Clauses(conds ...clause.Expression) IWinSettingDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winSettingDo) Returning(value interface{}, columns ...string) IWinSettingDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winSettingDo) Not(conds ...gen.Condition) IWinSettingDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winSettingDo) Or(conds ...gen.Condition) IWinSettingDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winSettingDo) Select(conds ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winSettingDo) Where(conds ...gen.Condition) IWinSettingDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winSettingDo) Order(conds ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winSettingDo) Distinct(cols ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winSettingDo) Omit(cols ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winSettingDo) Join(table schema.Tabler, on ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winSettingDo) Group(cols ...field.Expr) IWinSettingDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winSettingDo) Having(conds ...gen.Condition) IWinSettingDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winSettingDo) Limit(limit int) IWinSettingDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winSettingDo) Offset(offset int) IWinSettingDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinSettingDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winSettingDo) Unscoped() IWinSettingDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winSettingDo) Create(values ...*model.WinSetting) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winSettingDo) CreateInBatches(values []*model.WinSetting, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winSettingDo) Save(values ...*model.WinSetting) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winSettingDo) First() (*model.WinSetting, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSetting), nil
	}
}

func (w winSettingDo) Take() (*model.WinSetting, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSetting), nil
	}
}

func (w winSettingDo) Last() (*model.WinSetting, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSetting), nil
	}
}

func (w winSettingDo) Find() ([]*model.WinSetting, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinSetting), err
}

func (w winSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinSetting, err error) {
	buf := make([]*model.WinSetting, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winSettingDo) FindInBatches(result *[]*model.WinSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winSettingDo) Attrs(attrs ...field.AssignExpr) IWinSettingDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winSettingDo) Assign(attrs ...field.AssignExpr) IWinSettingDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winSettingDo) Joins(fields ...field.RelationField) IWinSettingDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winSettingDo) Preload(fields ...field.RelationField) IWinSettingDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winSettingDo) FirstOrInit() (*model.WinSetting, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSetting), nil
	}
}

func (w winSettingDo) FirstOrCreate() (*model.WinSetting, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinSetting), nil
	}
}

func (w winSettingDo) FindByPage(offset int, limit int) (result []*model.WinSetting, count int64, err error) {
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

func (w winSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winSettingDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winSettingDo) Delete(models ...*model.WinSetting) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winSettingDo) withDO(do gen.Dao) *winSettingDo {
	w.DO = *do.(*gen.DO)
	return w
}
