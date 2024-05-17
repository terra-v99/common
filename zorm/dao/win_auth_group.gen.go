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

func newWinAuthGroup(db *gorm.DB, opts ...gen.DOOption) winAuthGroup {
	_winAuthGroup := winAuthGroup{}

	_winAuthGroup.winAuthGroupDo.UseDB(db, opts...)
	_winAuthGroup.winAuthGroupDo.UseModel(&model.WinAuthGroup{})

	tableName := _winAuthGroup.winAuthGroupDo.TableName()
	_winAuthGroup.ALL = field.NewAsterisk(tableName)
	_winAuthGroup.ID = field.NewInt64(tableName, "id")
	_winAuthGroup.AdminGroupID = field.NewInt64(tableName, "admin_group_id")
	_winAuthGroup.Pid = field.NewInt64(tableName, "pid")
	_winAuthGroup.Parent = field.NewInt64(tableName, "parent")
	_winAuthGroup.Title = field.NewString(tableName, "title")
	_winAuthGroup.Status = field.NewInt64(tableName, "status")
	_winAuthGroup.Rules = field.NewString(tableName, "rules")
	_winAuthGroup.OperateID = field.NewInt64(tableName, "operate_id")
	_winAuthGroup.DataPermission = field.NewString(tableName, "data_permission")
	_winAuthGroup.CreatedAt = field.NewInt64(tableName, "created_at")
	_winAuthGroup.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winAuthGroup.fillFieldMap()

	return _winAuthGroup
}

type winAuthGroup struct {
	winAuthGroupDo

	ALL            field.Asterisk
	ID             field.Int64
	AdminGroupID   field.Int64  // 角色组ID
	Pid            field.Int64  // 父类ID
	Parent         field.Int64  // 创建人
	Title          field.String // 角色名
	Status         field.Int64  // 状态: 1-启用 0-冻结
	Rules          field.String // 菜单ID集合
	OperateID      field.Int64  // 用户组ID
	DataPermission field.String // 数据权限
	CreatedAt      field.Int64
	UpdatedAt      field.Int64

	fieldMap map[string]field.Expr
}

func (w winAuthGroup) Table(newTableName string) *winAuthGroup {
	w.winAuthGroupDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAuthGroup) As(alias string) *winAuthGroup {
	w.winAuthGroupDo.DO = *(w.winAuthGroupDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAuthGroup) updateTableName(table string) *winAuthGroup {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.AdminGroupID = field.NewInt64(table, "admin_group_id")
	w.Pid = field.NewInt64(table, "pid")
	w.Parent = field.NewInt64(table, "parent")
	w.Title = field.NewString(table, "title")
	w.Status = field.NewInt64(table, "status")
	w.Rules = field.NewString(table, "rules")
	w.OperateID = field.NewInt64(table, "operate_id")
	w.DataPermission = field.NewString(table, "data_permission")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winAuthGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAuthGroup) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["admin_group_id"] = w.AdminGroupID
	w.fieldMap["pid"] = w.Pid
	w.fieldMap["parent"] = w.Parent
	w.fieldMap["title"] = w.Title
	w.fieldMap["status"] = w.Status
	w.fieldMap["rules"] = w.Rules
	w.fieldMap["operate_id"] = w.OperateID
	w.fieldMap["data_permission"] = w.DataPermission
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winAuthGroup) clone(db *gorm.DB) winAuthGroup {
	w.winAuthGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAuthGroup) replaceDB(db *gorm.DB) winAuthGroup {
	w.winAuthGroupDo.ReplaceDB(db)
	return w
}

type winAuthGroupDo struct{ gen.DO }

type IWinAuthGroupDo interface {
	gen.SubQuery
	Debug() IWinAuthGroupDo
	WithContext(ctx context.Context) IWinAuthGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinAuthGroupDo
	WriteDB() IWinAuthGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinAuthGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinAuthGroupDo
	Not(conds ...gen.Condition) IWinAuthGroupDo
	Or(conds ...gen.Condition) IWinAuthGroupDo
	Select(conds ...field.Expr) IWinAuthGroupDo
	Where(conds ...gen.Condition) IWinAuthGroupDo
	Order(conds ...field.Expr) IWinAuthGroupDo
	Distinct(cols ...field.Expr) IWinAuthGroupDo
	Omit(cols ...field.Expr) IWinAuthGroupDo
	Join(table schema.Tabler, on ...field.Expr) IWinAuthGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinAuthGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinAuthGroupDo
	Group(cols ...field.Expr) IWinAuthGroupDo
	Having(conds ...gen.Condition) IWinAuthGroupDo
	Limit(limit int) IWinAuthGroupDo
	Offset(offset int) IWinAuthGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAuthGroupDo
	Unscoped() IWinAuthGroupDo
	Create(values ...*model.WinAuthGroup) error
	CreateInBatches(values []*model.WinAuthGroup, batchSize int) error
	Save(values ...*model.WinAuthGroup) error
	First() (*model.WinAuthGroup, error)
	Take() (*model.WinAuthGroup, error)
	Last() (*model.WinAuthGroup, error)
	Find() ([]*model.WinAuthGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAuthGroup, err error)
	FindInBatches(result *[]*model.WinAuthGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinAuthGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinAuthGroupDo
	Assign(attrs ...field.AssignExpr) IWinAuthGroupDo
	Joins(fields ...field.RelationField) IWinAuthGroupDo
	Preload(fields ...field.RelationField) IWinAuthGroupDo
	FirstOrInit() (*model.WinAuthGroup, error)
	FirstOrCreate() (*model.WinAuthGroup, error)
	FindByPage(offset int, limit int) (result []*model.WinAuthGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinAuthGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winAuthGroupDo) Debug() IWinAuthGroupDo {
	return w.withDO(w.DO.Debug())
}

func (w winAuthGroupDo) WithContext(ctx context.Context) IWinAuthGroupDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAuthGroupDo) ReadDB() IWinAuthGroupDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAuthGroupDo) WriteDB() IWinAuthGroupDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAuthGroupDo) Session(config *gorm.Session) IWinAuthGroupDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAuthGroupDo) Clauses(conds ...clause.Expression) IWinAuthGroupDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAuthGroupDo) Returning(value interface{}, columns ...string) IWinAuthGroupDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAuthGroupDo) Not(conds ...gen.Condition) IWinAuthGroupDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAuthGroupDo) Or(conds ...gen.Condition) IWinAuthGroupDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAuthGroupDo) Select(conds ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAuthGroupDo) Where(conds ...gen.Condition) IWinAuthGroupDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAuthGroupDo) Order(conds ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAuthGroupDo) Distinct(cols ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAuthGroupDo) Omit(cols ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAuthGroupDo) Join(table schema.Tabler, on ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAuthGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAuthGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAuthGroupDo) Group(cols ...field.Expr) IWinAuthGroupDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAuthGroupDo) Having(conds ...gen.Condition) IWinAuthGroupDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAuthGroupDo) Limit(limit int) IWinAuthGroupDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAuthGroupDo) Offset(offset int) IWinAuthGroupDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAuthGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAuthGroupDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAuthGroupDo) Unscoped() IWinAuthGroupDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAuthGroupDo) Create(values ...*model.WinAuthGroup) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAuthGroupDo) CreateInBatches(values []*model.WinAuthGroup, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAuthGroupDo) Save(values ...*model.WinAuthGroup) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAuthGroupDo) First() (*model.WinAuthGroup, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroup), nil
	}
}

func (w winAuthGroupDo) Take() (*model.WinAuthGroup, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroup), nil
	}
}

func (w winAuthGroupDo) Last() (*model.WinAuthGroup, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroup), nil
	}
}

func (w winAuthGroupDo) Find() ([]*model.WinAuthGroup, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAuthGroup), err
}

func (w winAuthGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAuthGroup, err error) {
	buf := make([]*model.WinAuthGroup, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAuthGroupDo) FindInBatches(result *[]*model.WinAuthGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAuthGroupDo) Attrs(attrs ...field.AssignExpr) IWinAuthGroupDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAuthGroupDo) Assign(attrs ...field.AssignExpr) IWinAuthGroupDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAuthGroupDo) Joins(fields ...field.RelationField) IWinAuthGroupDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAuthGroupDo) Preload(fields ...field.RelationField) IWinAuthGroupDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAuthGroupDo) FirstOrInit() (*model.WinAuthGroup, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroup), nil
	}
}

func (w winAuthGroupDo) FirstOrCreate() (*model.WinAuthGroup, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthGroup), nil
	}
}

func (w winAuthGroupDo) FindByPage(offset int, limit int) (result []*model.WinAuthGroup, count int64, err error) {
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

func (w winAuthGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAuthGroupDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAuthGroupDo) Delete(models ...*model.WinAuthGroup) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAuthGroupDo) withDO(do gen.Dao) *winAuthGroupDo {
	w.DO = *do.(*gen.DO)
	return w
}
