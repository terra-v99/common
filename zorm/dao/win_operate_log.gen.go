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

func newWinOperateLog(db *gorm.DB, opts ...gen.DOOption) winOperateLog {
	_winOperateLog := winOperateLog{}

	_winOperateLog.winOperateLogDo.UseDB(db, opts...)
	_winOperateLog.winOperateLogDo.UseModel(&model.WinOperateLog{})

	tableName := _winOperateLog.winOperateLogDo.TableName()
	_winOperateLog.ALL = field.NewAsterisk(tableName)
	_winOperateLog.ID = field.NewInt64(tableName, "id")
	_winOperateLog.UID = field.NewInt64(tableName, "uid")
	_winOperateLog.Username = field.NewString(tableName, "username")
	_winOperateLog.Title = field.NewString(tableName, "title")
	_winOperateLog.URL = field.NewString(tableName, "url")
	_winOperateLog.ReqParams = field.NewString(tableName, "req_params")
	_winOperateLog.ResParams = field.NewString(tableName, "res_params")
	_winOperateLog.IP = field.NewString(tableName, "ip")
	_winOperateLog.CreatedAt = field.NewInt64(tableName, "created_at")
	_winOperateLog.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winOperateLog.fillFieldMap()

	return _winOperateLog
}

type winOperateLog struct {
	winOperateLogDo

	ALL       field.Asterisk
	ID        field.Int64
	UID       field.Int64  // UID
	Username  field.String // 用户名
	Title     field.String // 标题
	URL       field.String // 请求url
	ReqParams field.String // 请求参数
	ResParams field.String // 响应参数
	IP        field.String // IP地址
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winOperateLog) Table(newTableName string) *winOperateLog {
	w.winOperateLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winOperateLog) As(alias string) *winOperateLog {
	w.winOperateLogDo.DO = *(w.winOperateLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winOperateLog) updateTableName(table string) *winOperateLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Title = field.NewString(table, "title")
	w.URL = field.NewString(table, "url")
	w.ReqParams = field.NewString(table, "req_params")
	w.ResParams = field.NewString(table, "res_params")
	w.IP = field.NewString(table, "ip")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winOperateLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winOperateLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 10)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["title"] = w.Title
	w.fieldMap["url"] = w.URL
	w.fieldMap["req_params"] = w.ReqParams
	w.fieldMap["res_params"] = w.ResParams
	w.fieldMap["ip"] = w.IP
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winOperateLog) clone(db *gorm.DB) winOperateLog {
	w.winOperateLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winOperateLog) replaceDB(db *gorm.DB) winOperateLog {
	w.winOperateLogDo.ReplaceDB(db)
	return w
}

type winOperateLogDo struct{ gen.DO }

type IWinOperateLogDo interface {
	gen.SubQuery
	Debug() IWinOperateLogDo
	WithContext(ctx context.Context) IWinOperateLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinOperateLogDo
	WriteDB() IWinOperateLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinOperateLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinOperateLogDo
	Not(conds ...gen.Condition) IWinOperateLogDo
	Or(conds ...gen.Condition) IWinOperateLogDo
	Select(conds ...field.Expr) IWinOperateLogDo
	Where(conds ...gen.Condition) IWinOperateLogDo
	Order(conds ...field.Expr) IWinOperateLogDo
	Distinct(cols ...field.Expr) IWinOperateLogDo
	Omit(cols ...field.Expr) IWinOperateLogDo
	Join(table schema.Tabler, on ...field.Expr) IWinOperateLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinOperateLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinOperateLogDo
	Group(cols ...field.Expr) IWinOperateLogDo
	Having(conds ...gen.Condition) IWinOperateLogDo
	Limit(limit int) IWinOperateLogDo
	Offset(offset int) IWinOperateLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinOperateLogDo
	Unscoped() IWinOperateLogDo
	Create(values ...*model.WinOperateLog) error
	CreateInBatches(values []*model.WinOperateLog, batchSize int) error
	Save(values ...*model.WinOperateLog) error
	First() (*model.WinOperateLog, error)
	Take() (*model.WinOperateLog, error)
	Last() (*model.WinOperateLog, error)
	Find() ([]*model.WinOperateLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinOperateLog, err error)
	FindInBatches(result *[]*model.WinOperateLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinOperateLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinOperateLogDo
	Assign(attrs ...field.AssignExpr) IWinOperateLogDo
	Joins(fields ...field.RelationField) IWinOperateLogDo
	Preload(fields ...field.RelationField) IWinOperateLogDo
	FirstOrInit() (*model.WinOperateLog, error)
	FirstOrCreate() (*model.WinOperateLog, error)
	FindByPage(offset int, limit int) (result []*model.WinOperateLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinOperateLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winOperateLogDo) Debug() IWinOperateLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winOperateLogDo) WithContext(ctx context.Context) IWinOperateLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winOperateLogDo) ReadDB() IWinOperateLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winOperateLogDo) WriteDB() IWinOperateLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winOperateLogDo) Session(config *gorm.Session) IWinOperateLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winOperateLogDo) Clauses(conds ...clause.Expression) IWinOperateLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winOperateLogDo) Returning(value interface{}, columns ...string) IWinOperateLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winOperateLogDo) Not(conds ...gen.Condition) IWinOperateLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winOperateLogDo) Or(conds ...gen.Condition) IWinOperateLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winOperateLogDo) Select(conds ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winOperateLogDo) Where(conds ...gen.Condition) IWinOperateLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winOperateLogDo) Order(conds ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winOperateLogDo) Distinct(cols ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winOperateLogDo) Omit(cols ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winOperateLogDo) Join(table schema.Tabler, on ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winOperateLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winOperateLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winOperateLogDo) Group(cols ...field.Expr) IWinOperateLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winOperateLogDo) Having(conds ...gen.Condition) IWinOperateLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winOperateLogDo) Limit(limit int) IWinOperateLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winOperateLogDo) Offset(offset int) IWinOperateLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winOperateLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinOperateLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winOperateLogDo) Unscoped() IWinOperateLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winOperateLogDo) Create(values ...*model.WinOperateLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winOperateLogDo) CreateInBatches(values []*model.WinOperateLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winOperateLogDo) Save(values ...*model.WinOperateLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winOperateLogDo) First() (*model.WinOperateLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperateLog), nil
	}
}

func (w winOperateLogDo) Take() (*model.WinOperateLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperateLog), nil
	}
}

func (w winOperateLogDo) Last() (*model.WinOperateLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperateLog), nil
	}
}

func (w winOperateLogDo) Find() ([]*model.WinOperateLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinOperateLog), err
}

func (w winOperateLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinOperateLog, err error) {
	buf := make([]*model.WinOperateLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winOperateLogDo) FindInBatches(result *[]*model.WinOperateLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winOperateLogDo) Attrs(attrs ...field.AssignExpr) IWinOperateLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winOperateLogDo) Assign(attrs ...field.AssignExpr) IWinOperateLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winOperateLogDo) Joins(fields ...field.RelationField) IWinOperateLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winOperateLogDo) Preload(fields ...field.RelationField) IWinOperateLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winOperateLogDo) FirstOrInit() (*model.WinOperateLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperateLog), nil
	}
}

func (w winOperateLogDo) FirstOrCreate() (*model.WinOperateLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperateLog), nil
	}
}

func (w winOperateLogDo) FindByPage(offset int, limit int) (result []*model.WinOperateLog, count int64, err error) {
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

func (w winOperateLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winOperateLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winOperateLogDo) Delete(models ...*model.WinOperateLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winOperateLogDo) withDO(do gen.Dao) *winOperateLogDo {
	w.DO = *do.(*gen.DO)
	return w
}
