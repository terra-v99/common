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

func newWinUserRecentGameLog(db *gorm.DB, opts ...gen.DOOption) winUserRecentGameLog {
	_winUserRecentGameLog := winUserRecentGameLog{}

	_winUserRecentGameLog.winUserRecentGameLogDo.UseDB(db, opts...)
	_winUserRecentGameLog.winUserRecentGameLogDo.UseModel(&model.WinUserRecentGameLog{})

	tableName := _winUserRecentGameLog.winUserRecentGameLogDo.TableName()
	_winUserRecentGameLog.ALL = field.NewAsterisk(tableName)
	_winUserRecentGameLog.ID = field.NewInt64(tableName, "id")
	_winUserRecentGameLog.UID = field.NewInt64(tableName, "uid")
	_winUserRecentGameLog.GameName = field.NewString(tableName, "game_name")
	_winUserRecentGameLog.GameType = field.NewInt64(tableName, "game_type")
	_winUserRecentGameLog.KeyID = field.NewString(tableName, "key_id")
	_winUserRecentGameLog.SlotID = field.NewString(tableName, "slot_id")
	_winUserRecentGameLog.GameListID = field.NewInt64(tableName, "game_list_id")
	_winUserRecentGameLog.GamePlatID = field.NewInt64(tableName, "game_plat_id")
	_winUserRecentGameLog.CreatedAt = field.NewInt64(tableName, "created_at")

	_winUserRecentGameLog.fillFieldMap()

	return _winUserRecentGameLog
}

type winUserRecentGameLog struct {
	winUserRecentGameLogDo

	ALL        field.Asterisk
	ID         field.Int64
	UID        field.Int64  // UID
	GameName   field.String // 游戏名称
	GameType   field.Int64  // 游戏分类
	KeyID      field.String // 游戏唯一
	SlotID     field.String // slot_id
	GameListID field.Int64  // 游戏分类id
	GamePlatID field.Int64  // 平台id
	CreatedAt  field.Int64  // 创建时间

	fieldMap map[string]field.Expr
}

func (w winUserRecentGameLog) Table(newTableName string) *winUserRecentGameLog {
	w.winUserRecentGameLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserRecentGameLog) As(alias string) *winUserRecentGameLog {
	w.winUserRecentGameLogDo.DO = *(w.winUserRecentGameLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserRecentGameLog) updateTableName(table string) *winUserRecentGameLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.GameName = field.NewString(table, "game_name")
	w.GameType = field.NewInt64(table, "game_type")
	w.KeyID = field.NewString(table, "key_id")
	w.SlotID = field.NewString(table, "slot_id")
	w.GameListID = field.NewInt64(table, "game_list_id")
	w.GamePlatID = field.NewInt64(table, "game_plat_id")
	w.CreatedAt = field.NewInt64(table, "created_at")

	w.fillFieldMap()

	return w
}

func (w *winUserRecentGameLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserRecentGameLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["game_name"] = w.GameName
	w.fieldMap["game_type"] = w.GameType
	w.fieldMap["key_id"] = w.KeyID
	w.fieldMap["slot_id"] = w.SlotID
	w.fieldMap["game_list_id"] = w.GameListID
	w.fieldMap["game_plat_id"] = w.GamePlatID
	w.fieldMap["created_at"] = w.CreatedAt
}

func (w winUserRecentGameLog) clone(db *gorm.DB) winUserRecentGameLog {
	w.winUserRecentGameLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserRecentGameLog) replaceDB(db *gorm.DB) winUserRecentGameLog {
	w.winUserRecentGameLogDo.ReplaceDB(db)
	return w
}

type winUserRecentGameLogDo struct{ gen.DO }

type IWinUserRecentGameLogDo interface {
	gen.SubQuery
	Debug() IWinUserRecentGameLogDo
	WithContext(ctx context.Context) IWinUserRecentGameLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserRecentGameLogDo
	WriteDB() IWinUserRecentGameLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserRecentGameLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserRecentGameLogDo
	Not(conds ...gen.Condition) IWinUserRecentGameLogDo
	Or(conds ...gen.Condition) IWinUserRecentGameLogDo
	Select(conds ...field.Expr) IWinUserRecentGameLogDo
	Where(conds ...gen.Condition) IWinUserRecentGameLogDo
	Order(conds ...field.Expr) IWinUserRecentGameLogDo
	Distinct(cols ...field.Expr) IWinUserRecentGameLogDo
	Omit(cols ...field.Expr) IWinUserRecentGameLogDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserRecentGameLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserRecentGameLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserRecentGameLogDo
	Group(cols ...field.Expr) IWinUserRecentGameLogDo
	Having(conds ...gen.Condition) IWinUserRecentGameLogDo
	Limit(limit int) IWinUserRecentGameLogDo
	Offset(offset int) IWinUserRecentGameLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserRecentGameLogDo
	Unscoped() IWinUserRecentGameLogDo
	Create(values ...*model.WinUserRecentGameLog) error
	CreateInBatches(values []*model.WinUserRecentGameLog, batchSize int) error
	Save(values ...*model.WinUserRecentGameLog) error
	First() (*model.WinUserRecentGameLog, error)
	Take() (*model.WinUserRecentGameLog, error)
	Last() (*model.WinUserRecentGameLog, error)
	Find() ([]*model.WinUserRecentGameLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserRecentGameLog, err error)
	FindInBatches(result *[]*model.WinUserRecentGameLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserRecentGameLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserRecentGameLogDo
	Assign(attrs ...field.AssignExpr) IWinUserRecentGameLogDo
	Joins(fields ...field.RelationField) IWinUserRecentGameLogDo
	Preload(fields ...field.RelationField) IWinUserRecentGameLogDo
	FirstOrInit() (*model.WinUserRecentGameLog, error)
	FirstOrCreate() (*model.WinUserRecentGameLog, error)
	FindByPage(offset int, limit int) (result []*model.WinUserRecentGameLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserRecentGameLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserRecentGameLogDo) Debug() IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserRecentGameLogDo) WithContext(ctx context.Context) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserRecentGameLogDo) ReadDB() IWinUserRecentGameLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserRecentGameLogDo) WriteDB() IWinUserRecentGameLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserRecentGameLogDo) Session(config *gorm.Session) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserRecentGameLogDo) Clauses(conds ...clause.Expression) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserRecentGameLogDo) Returning(value interface{}, columns ...string) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserRecentGameLogDo) Not(conds ...gen.Condition) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserRecentGameLogDo) Or(conds ...gen.Condition) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserRecentGameLogDo) Select(conds ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserRecentGameLogDo) Where(conds ...gen.Condition) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserRecentGameLogDo) Order(conds ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserRecentGameLogDo) Distinct(cols ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserRecentGameLogDo) Omit(cols ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserRecentGameLogDo) Join(table schema.Tabler, on ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserRecentGameLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserRecentGameLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserRecentGameLogDo) Group(cols ...field.Expr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserRecentGameLogDo) Having(conds ...gen.Condition) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserRecentGameLogDo) Limit(limit int) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserRecentGameLogDo) Offset(offset int) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserRecentGameLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserRecentGameLogDo) Unscoped() IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserRecentGameLogDo) Create(values ...*model.WinUserRecentGameLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserRecentGameLogDo) CreateInBatches(values []*model.WinUserRecentGameLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserRecentGameLogDo) Save(values ...*model.WinUserRecentGameLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserRecentGameLogDo) First() (*model.WinUserRecentGameLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserRecentGameLog), nil
	}
}

func (w winUserRecentGameLogDo) Take() (*model.WinUserRecentGameLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserRecentGameLog), nil
	}
}

func (w winUserRecentGameLogDo) Last() (*model.WinUserRecentGameLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserRecentGameLog), nil
	}
}

func (w winUserRecentGameLogDo) Find() ([]*model.WinUserRecentGameLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserRecentGameLog), err
}

func (w winUserRecentGameLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserRecentGameLog, err error) {
	buf := make([]*model.WinUserRecentGameLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserRecentGameLogDo) FindInBatches(result *[]*model.WinUserRecentGameLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserRecentGameLogDo) Attrs(attrs ...field.AssignExpr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserRecentGameLogDo) Assign(attrs ...field.AssignExpr) IWinUserRecentGameLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserRecentGameLogDo) Joins(fields ...field.RelationField) IWinUserRecentGameLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserRecentGameLogDo) Preload(fields ...field.RelationField) IWinUserRecentGameLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserRecentGameLogDo) FirstOrInit() (*model.WinUserRecentGameLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserRecentGameLog), nil
	}
}

func (w winUserRecentGameLogDo) FirstOrCreate() (*model.WinUserRecentGameLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserRecentGameLog), nil
	}
}

func (w winUserRecentGameLogDo) FindByPage(offset int, limit int) (result []*model.WinUserRecentGameLog, count int64, err error) {
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

func (w winUserRecentGameLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserRecentGameLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserRecentGameLogDo) Delete(models ...*model.WinUserRecentGameLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserRecentGameLogDo) withDO(do gen.Dao) *winUserRecentGameLogDo {
	w.DO = *do.(*gen.DO)
	return w
}
