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

func newWinUserBankList(db *gorm.DB, opts ...gen.DOOption) winUserBankList {
	_winUserBankList := winUserBankList{}

	_winUserBankList.winUserBankListDo.UseDB(db, opts...)
	_winUserBankList.winUserBankListDo.UseModel(&model.WinUserBankList{})

	tableName := _winUserBankList.winUserBankListDo.TableName()
	_winUserBankList.ALL = field.NewAsterisk(tableName)
	_winUserBankList.ID = field.NewInt64(tableName, "id")
	_winUserBankList.UID = field.NewInt64(tableName, "uid")
	_winUserBankList.Username = field.NewString(tableName, "username")
	_winUserBankList.CategoryCurrency = field.NewInt64(tableName, "category_currency")
	_winUserBankList.CategoryTransfer = field.NewInt64(tableName, "category_transfer")
	_winUserBankList.Address = field.NewString(tableName, "address")
	_winUserBankList.Status = field.NewInt64(tableName, "status")
	_winUserBankList.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserBankList.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winUserBankList.CreatedBy = field.NewString(tableName, "created_by")
	_winUserBankList.UpdatedBy = field.NewString(tableName, "updated_by")
	_winUserBankList.OperatorName = field.NewString(tableName, "operator_name")

	_winUserBankList.fillFieldMap()

	return _winUserBankList
}

type winUserBankList struct {
	winUserBankListDo

	ALL              field.Asterisk
	ID               field.Int64
	UID              field.Int64  // UID
	Username         field.String // 用户名
	CategoryCurrency field.Int64  // 货币类型:0-数字货币 1-法币
	CategoryTransfer field.Int64  // 转账类型：1-TRC,2-ERC,3-BANK,4-PIX,5-GCASH
	Address          field.String // 提款地址
	Status           field.Int64  // 状态:1-默认地址(启用) 2-正常启用 3-删除
	CreatedAt        field.Int64
	UpdatedAt        field.Int64
	CreatedBy        field.String // 后台创建人
	UpdatedBy        field.String // 后台修改人
	OperatorName     field.String // 操作人姓名

	fieldMap map[string]field.Expr
}

func (w winUserBankList) Table(newTableName string) *winUserBankList {
	w.winUserBankListDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserBankList) As(alias string) *winUserBankList {
	w.winUserBankListDo.DO = *(w.winUserBankListDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserBankList) updateTableName(table string) *winUserBankList {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.CategoryCurrency = field.NewInt64(table, "category_currency")
	w.CategoryTransfer = field.NewInt64(table, "category_transfer")
	w.Address = field.NewString(table, "address")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.CreatedBy = field.NewString(table, "created_by")
	w.UpdatedBy = field.NewString(table, "updated_by")
	w.OperatorName = field.NewString(table, "operator_name")

	w.fillFieldMap()

	return w
}

func (w *winUserBankList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserBankList) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 12)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["category_currency"] = w.CategoryCurrency
	w.fieldMap["category_transfer"] = w.CategoryTransfer
	w.fieldMap["address"] = w.Address
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["created_by"] = w.CreatedBy
	w.fieldMap["updated_by"] = w.UpdatedBy
	w.fieldMap["operator_name"] = w.OperatorName
}

func (w winUserBankList) clone(db *gorm.DB) winUserBankList {
	w.winUserBankListDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserBankList) replaceDB(db *gorm.DB) winUserBankList {
	w.winUserBankListDo.ReplaceDB(db)
	return w
}

type winUserBankListDo struct{ gen.DO }

type IWinUserBankListDo interface {
	gen.SubQuery
	Debug() IWinUserBankListDo
	WithContext(ctx context.Context) IWinUserBankListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserBankListDo
	WriteDB() IWinUserBankListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserBankListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserBankListDo
	Not(conds ...gen.Condition) IWinUserBankListDo
	Or(conds ...gen.Condition) IWinUserBankListDo
	Select(conds ...field.Expr) IWinUserBankListDo
	Where(conds ...gen.Condition) IWinUserBankListDo
	Order(conds ...field.Expr) IWinUserBankListDo
	Distinct(cols ...field.Expr) IWinUserBankListDo
	Omit(cols ...field.Expr) IWinUserBankListDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserBankListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserBankListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserBankListDo
	Group(cols ...field.Expr) IWinUserBankListDo
	Having(conds ...gen.Condition) IWinUserBankListDo
	Limit(limit int) IWinUserBankListDo
	Offset(offset int) IWinUserBankListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserBankListDo
	Unscoped() IWinUserBankListDo
	Create(values ...*model.WinUserBankList) error
	CreateInBatches(values []*model.WinUserBankList, batchSize int) error
	Save(values ...*model.WinUserBankList) error
	First() (*model.WinUserBankList, error)
	Take() (*model.WinUserBankList, error)
	Last() (*model.WinUserBankList, error)
	Find() ([]*model.WinUserBankList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserBankList, err error)
	FindInBatches(result *[]*model.WinUserBankList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserBankList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserBankListDo
	Assign(attrs ...field.AssignExpr) IWinUserBankListDo
	Joins(fields ...field.RelationField) IWinUserBankListDo
	Preload(fields ...field.RelationField) IWinUserBankListDo
	FirstOrInit() (*model.WinUserBankList, error)
	FirstOrCreate() (*model.WinUserBankList, error)
	FindByPage(offset int, limit int) (result []*model.WinUserBankList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserBankListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserBankListDo) Debug() IWinUserBankListDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserBankListDo) WithContext(ctx context.Context) IWinUserBankListDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserBankListDo) ReadDB() IWinUserBankListDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserBankListDo) WriteDB() IWinUserBankListDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserBankListDo) Session(config *gorm.Session) IWinUserBankListDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserBankListDo) Clauses(conds ...clause.Expression) IWinUserBankListDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserBankListDo) Returning(value interface{}, columns ...string) IWinUserBankListDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserBankListDo) Not(conds ...gen.Condition) IWinUserBankListDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserBankListDo) Or(conds ...gen.Condition) IWinUserBankListDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserBankListDo) Select(conds ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserBankListDo) Where(conds ...gen.Condition) IWinUserBankListDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserBankListDo) Order(conds ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserBankListDo) Distinct(cols ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserBankListDo) Omit(cols ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserBankListDo) Join(table schema.Tabler, on ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserBankListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserBankListDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserBankListDo) Group(cols ...field.Expr) IWinUserBankListDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserBankListDo) Having(conds ...gen.Condition) IWinUserBankListDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserBankListDo) Limit(limit int) IWinUserBankListDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserBankListDo) Offset(offset int) IWinUserBankListDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserBankListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserBankListDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserBankListDo) Unscoped() IWinUserBankListDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserBankListDo) Create(values ...*model.WinUserBankList) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserBankListDo) CreateInBatches(values []*model.WinUserBankList, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserBankListDo) Save(values ...*model.WinUserBankList) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserBankListDo) First() (*model.WinUserBankList, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserBankList), nil
	}
}

func (w winUserBankListDo) Take() (*model.WinUserBankList, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserBankList), nil
	}
}

func (w winUserBankListDo) Last() (*model.WinUserBankList, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserBankList), nil
	}
}

func (w winUserBankListDo) Find() ([]*model.WinUserBankList, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserBankList), err
}

func (w winUserBankListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserBankList, err error) {
	buf := make([]*model.WinUserBankList, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserBankListDo) FindInBatches(result *[]*model.WinUserBankList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserBankListDo) Attrs(attrs ...field.AssignExpr) IWinUserBankListDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserBankListDo) Assign(attrs ...field.AssignExpr) IWinUserBankListDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserBankListDo) Joins(fields ...field.RelationField) IWinUserBankListDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserBankListDo) Preload(fields ...field.RelationField) IWinUserBankListDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserBankListDo) FirstOrInit() (*model.WinUserBankList, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserBankList), nil
	}
}

func (w winUserBankListDo) FirstOrCreate() (*model.WinUserBankList, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserBankList), nil
	}
}

func (w winUserBankListDo) FindByPage(offset int, limit int) (result []*model.WinUserBankList, count int64, err error) {
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

func (w winUserBankListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserBankListDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserBankListDo) Delete(models ...*model.WinUserBankList) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserBankListDo) withDO(do gen.Dao) *winUserBankListDo {
	w.DO = *do.(*gen.DO)
	return w
}
