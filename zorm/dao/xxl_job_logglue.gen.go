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

func newXxlJobLogglue(db *gorm.DB, opts ...gen.DOOption) xxlJobLogglue {
	_xxlJobLogglue := xxlJobLogglue{}

	_xxlJobLogglue.xxlJobLogglueDo.UseDB(db, opts...)
	_xxlJobLogglue.xxlJobLogglueDo.UseModel(&model.XxlJobLogglue{})

	tableName := _xxlJobLogglue.xxlJobLogglueDo.TableName()
	_xxlJobLogglue.ALL = field.NewAsterisk(tableName)
	_xxlJobLogglue.ID = field.NewInt64(tableName, "id")
	_xxlJobLogglue.JobID = field.NewInt64(tableName, "job_id")
	_xxlJobLogglue.GlueType = field.NewString(tableName, "glue_type")
	_xxlJobLogglue.GlueSource = field.NewString(tableName, "glue_source")
	_xxlJobLogglue.GlueRemark = field.NewString(tableName, "glue_remark")
	_xxlJobLogglue.AddTime = field.NewTime(tableName, "add_time")
	_xxlJobLogglue.UpdateTime = field.NewTime(tableName, "update_time")

	_xxlJobLogglue.fillFieldMap()

	return _xxlJobLogglue
}

type xxlJobLogglue struct {
	xxlJobLogglueDo

	ALL        field.Asterisk
	ID         field.Int64
	JobID      field.Int64  // 任务，主键ID
	GlueType   field.String // GLUE类型
	GlueSource field.String // GLUE源代码
	GlueRemark field.String // GLUE备注
	AddTime    field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (x xxlJobLogglue) Table(newTableName string) *xxlJobLogglue {
	x.xxlJobLogglueDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xxlJobLogglue) As(alias string) *xxlJobLogglue {
	x.xxlJobLogglueDo.DO = *(x.xxlJobLogglueDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xxlJobLogglue) updateTableName(table string) *xxlJobLogglue {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.JobID = field.NewInt64(table, "job_id")
	x.GlueType = field.NewString(table, "glue_type")
	x.GlueSource = field.NewString(table, "glue_source")
	x.GlueRemark = field.NewString(table, "glue_remark")
	x.AddTime = field.NewTime(table, "add_time")
	x.UpdateTime = field.NewTime(table, "update_time")

	x.fillFieldMap()

	return x
}

func (x *xxlJobLogglue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xxlJobLogglue) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 7)
	x.fieldMap["id"] = x.ID
	x.fieldMap["job_id"] = x.JobID
	x.fieldMap["glue_type"] = x.GlueType
	x.fieldMap["glue_source"] = x.GlueSource
	x.fieldMap["glue_remark"] = x.GlueRemark
	x.fieldMap["add_time"] = x.AddTime
	x.fieldMap["update_time"] = x.UpdateTime
}

func (x xxlJobLogglue) clone(db *gorm.DB) xxlJobLogglue {
	x.xxlJobLogglueDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xxlJobLogglue) replaceDB(db *gorm.DB) xxlJobLogglue {
	x.xxlJobLogglueDo.ReplaceDB(db)
	return x
}

type xxlJobLogglueDo struct{ gen.DO }

type IXxlJobLogglueDo interface {
	gen.SubQuery
	Debug() IXxlJobLogglueDo
	WithContext(ctx context.Context) IXxlJobLogglueDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXxlJobLogglueDo
	WriteDB() IXxlJobLogglueDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXxlJobLogglueDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXxlJobLogglueDo
	Not(conds ...gen.Condition) IXxlJobLogglueDo
	Or(conds ...gen.Condition) IXxlJobLogglueDo
	Select(conds ...field.Expr) IXxlJobLogglueDo
	Where(conds ...gen.Condition) IXxlJobLogglueDo
	Order(conds ...field.Expr) IXxlJobLogglueDo
	Distinct(cols ...field.Expr) IXxlJobLogglueDo
	Omit(cols ...field.Expr) IXxlJobLogglueDo
	Join(table schema.Tabler, on ...field.Expr) IXxlJobLogglueDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogglueDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogglueDo
	Group(cols ...field.Expr) IXxlJobLogglueDo
	Having(conds ...gen.Condition) IXxlJobLogglueDo
	Limit(limit int) IXxlJobLogglueDo
	Offset(offset int) IXxlJobLogglueDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobLogglueDo
	Unscoped() IXxlJobLogglueDo
	Create(values ...*model.XxlJobLogglue) error
	CreateInBatches(values []*model.XxlJobLogglue, batchSize int) error
	Save(values ...*model.XxlJobLogglue) error
	First() (*model.XxlJobLogglue, error)
	Take() (*model.XxlJobLogglue, error)
	Last() (*model.XxlJobLogglue, error)
	Find() ([]*model.XxlJobLogglue, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobLogglue, err error)
	FindInBatches(result *[]*model.XxlJobLogglue, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XxlJobLogglue) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXxlJobLogglueDo
	Assign(attrs ...field.AssignExpr) IXxlJobLogglueDo
	Joins(fields ...field.RelationField) IXxlJobLogglueDo
	Preload(fields ...field.RelationField) IXxlJobLogglueDo
	FirstOrInit() (*model.XxlJobLogglue, error)
	FirstOrCreate() (*model.XxlJobLogglue, error)
	FindByPage(offset int, limit int) (result []*model.XxlJobLogglue, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXxlJobLogglueDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xxlJobLogglueDo) Debug() IXxlJobLogglueDo {
	return x.withDO(x.DO.Debug())
}

func (x xxlJobLogglueDo) WithContext(ctx context.Context) IXxlJobLogglueDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xxlJobLogglueDo) ReadDB() IXxlJobLogglueDo {
	return x.Clauses(dbresolver.Read)
}

func (x xxlJobLogglueDo) WriteDB() IXxlJobLogglueDo {
	return x.Clauses(dbresolver.Write)
}

func (x xxlJobLogglueDo) Session(config *gorm.Session) IXxlJobLogglueDo {
	return x.withDO(x.DO.Session(config))
}

func (x xxlJobLogglueDo) Clauses(conds ...clause.Expression) IXxlJobLogglueDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xxlJobLogglueDo) Returning(value interface{}, columns ...string) IXxlJobLogglueDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xxlJobLogglueDo) Not(conds ...gen.Condition) IXxlJobLogglueDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xxlJobLogglueDo) Or(conds ...gen.Condition) IXxlJobLogglueDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xxlJobLogglueDo) Select(conds ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xxlJobLogglueDo) Where(conds ...gen.Condition) IXxlJobLogglueDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xxlJobLogglueDo) Order(conds ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xxlJobLogglueDo) Distinct(cols ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xxlJobLogglueDo) Omit(cols ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xxlJobLogglueDo) Join(table schema.Tabler, on ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xxlJobLogglueDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xxlJobLogglueDo) RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xxlJobLogglueDo) Group(cols ...field.Expr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xxlJobLogglueDo) Having(conds ...gen.Condition) IXxlJobLogglueDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xxlJobLogglueDo) Limit(limit int) IXxlJobLogglueDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xxlJobLogglueDo) Offset(offset int) IXxlJobLogglueDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xxlJobLogglueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobLogglueDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xxlJobLogglueDo) Unscoped() IXxlJobLogglueDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xxlJobLogglueDo) Create(values ...*model.XxlJobLogglue) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xxlJobLogglueDo) CreateInBatches(values []*model.XxlJobLogglue, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xxlJobLogglueDo) Save(values ...*model.XxlJobLogglue) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xxlJobLogglueDo) First() (*model.XxlJobLogglue, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogglue), nil
	}
}

func (x xxlJobLogglueDo) Take() (*model.XxlJobLogglue, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogglue), nil
	}
}

func (x xxlJobLogglueDo) Last() (*model.XxlJobLogglue, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogglue), nil
	}
}

func (x xxlJobLogglueDo) Find() ([]*model.XxlJobLogglue, error) {
	result, err := x.DO.Find()
	return result.([]*model.XxlJobLogglue), err
}

func (x xxlJobLogglueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobLogglue, err error) {
	buf := make([]*model.XxlJobLogglue, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xxlJobLogglueDo) FindInBatches(result *[]*model.XxlJobLogglue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xxlJobLogglueDo) Attrs(attrs ...field.AssignExpr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xxlJobLogglueDo) Assign(attrs ...field.AssignExpr) IXxlJobLogglueDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xxlJobLogglueDo) Joins(fields ...field.RelationField) IXxlJobLogglueDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xxlJobLogglueDo) Preload(fields ...field.RelationField) IXxlJobLogglueDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xxlJobLogglueDo) FirstOrInit() (*model.XxlJobLogglue, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogglue), nil
	}
}

func (x xxlJobLogglueDo) FirstOrCreate() (*model.XxlJobLogglue, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLogglue), nil
	}
}

func (x xxlJobLogglueDo) FindByPage(offset int, limit int) (result []*model.XxlJobLogglue, count int64, err error) {
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

func (x xxlJobLogglueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xxlJobLogglueDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xxlJobLogglueDo) Delete(models ...*model.XxlJobLogglue) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xxlJobLogglueDo) withDO(do gen.Dao) *xxlJobLogglueDo {
	x.DO = *do.(*gen.DO)
	return x
}
