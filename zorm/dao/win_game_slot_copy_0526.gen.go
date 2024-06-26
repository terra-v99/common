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

func newWinGameSlotCopy0526(db *gorm.DB, opts ...gen.DOOption) winGameSlotCopy0526 {
	_winGameSlotCopy0526 := winGameSlotCopy0526{}

	_winGameSlotCopy0526.winGameSlotCopy0526Do.UseDB(db, opts...)
	_winGameSlotCopy0526.winGameSlotCopy0526Do.UseModel(&model.WinGameSlotCopy0526{})

	tableName := _winGameSlotCopy0526.winGameSlotCopy0526Do.TableName()
	_winGameSlotCopy0526.ALL = field.NewAsterisk(tableName)
	_winGameSlotCopy0526.ID = field.NewString(tableName, "id")
	_winGameSlotCopy0526.GameID = field.NewInt64(tableName, "game_id")
	_winGameSlotCopy0526.GameGroupID = field.NewInt64(tableName, "game_group_id")
	_winGameSlotCopy0526.PlatID = field.NewInt64(tableName, "plat_id")
	_winGameSlotCopy0526.Provider = field.NewString(tableName, "provider")
	_winGameSlotCopy0526.Name = field.NewString(tableName, "name")
	_winGameSlotCopy0526.NameZh = field.NewString(tableName, "name_zh")
	_winGameSlotCopy0526.Img = field.NewString(tableName, "img")
	_winGameSlotCopy0526.ImgNew = field.NewString(tableName, "img_new")
	_winGameSlotCopy0526.IsNew = field.NewInt64(tableName, "is_new")
	_winGameSlotCopy0526.IsCasino = field.NewInt64(tableName, "is_casino")
	_winGameSlotCopy0526.GameTypeID = field.NewString(tableName, "game_type_id")
	_winGameSlotCopy0526.GameTypeName = field.NewString(tableName, "game_type_name")
	_winGameSlotCopy0526.FavoriteStar = field.NewInt64(tableName, "favorite_star")
	_winGameSlotCopy0526.HotStar = field.NewInt64(tableName, "hot_star")
	_winGameSlotCopy0526.Sort = field.NewInt64(tableName, "sort")
	_winGameSlotCopy0526.Status = field.NewInt64(tableName, "status")
	_winGameSlotCopy0526.Device = field.NewInt64(tableName, "device")
	_winGameSlotCopy0526.CreatedAt = field.NewInt64(tableName, "created_at")
	_winGameSlotCopy0526.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winGameSlotCopy0526.UpdatedUser = field.NewString(tableName, "updated_user")
	_winGameSlotCopy0526.Maintenance = field.NewString(tableName, "maintenance")
	_winGameSlotCopy0526.OperatorName = field.NewString(tableName, "operator_name")
	_winGameSlotCopy0526.TestI = field.NewInt64(tableName, "test_i")

	_winGameSlotCopy0526.fillFieldMap()

	return _winGameSlotCopy0526
}

type winGameSlotCopy0526 struct {
	winGameSlotCopy0526Do

	ALL          field.Asterisk
	ID           field.String // ID(关联BrandGameId)
	GameID       field.Int64  // 游戏ID(关联game_list)
	GameGroupID  field.Int64  // 游戏大类类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能
	PlatID       field.Int64  // 游戏平台id
	Provider     field.String // 游戏提供者
	Name         field.String // 简体名称
	NameZh       field.String // 游戏名字(中文)
	Img          field.String // 英文图片
	ImgNew       field.String // 新版游戏图片
	IsNew        field.Int64  // 是否新游戏:1-是 0-否
	IsCasino     field.Int64  // 是否推荐主页 0否 1是
	GameTypeID   field.String // 游戏类型ID(0r code)
	GameTypeName field.String // 游戏类型名称
	FavoriteStar field.Int64  // 收藏值
	HotStar      field.Int64  // 热度
	Sort         field.Int64  // 排序
	Status       field.Int64  // 状态:1-启用 0-停用
	Device       field.Int64  // 设备:0-all 1-pc 2-h5
	CreatedAt    field.Int64
	UpdatedAt    field.Int64
	UpdatedUser  field.String // 最后更新人
	Maintenance  field.String // 维护信息
	OperatorName field.String // 操作人姓名
	TestI        field.Int64

	fieldMap map[string]field.Expr
}

func (w winGameSlotCopy0526) Table(newTableName string) *winGameSlotCopy0526 {
	w.winGameSlotCopy0526Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winGameSlotCopy0526) As(alias string) *winGameSlotCopy0526 {
	w.winGameSlotCopy0526Do.DO = *(w.winGameSlotCopy0526Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winGameSlotCopy0526) updateTableName(table string) *winGameSlotCopy0526 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewString(table, "id")
	w.GameID = field.NewInt64(table, "game_id")
	w.GameGroupID = field.NewInt64(table, "game_group_id")
	w.PlatID = field.NewInt64(table, "plat_id")
	w.Provider = field.NewString(table, "provider")
	w.Name = field.NewString(table, "name")
	w.NameZh = field.NewString(table, "name_zh")
	w.Img = field.NewString(table, "img")
	w.ImgNew = field.NewString(table, "img_new")
	w.IsNew = field.NewInt64(table, "is_new")
	w.IsCasino = field.NewInt64(table, "is_casino")
	w.GameTypeID = field.NewString(table, "game_type_id")
	w.GameTypeName = field.NewString(table, "game_type_name")
	w.FavoriteStar = field.NewInt64(table, "favorite_star")
	w.HotStar = field.NewInt64(table, "hot_star")
	w.Sort = field.NewInt64(table, "sort")
	w.Status = field.NewInt64(table, "status")
	w.Device = field.NewInt64(table, "device")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.UpdatedUser = field.NewString(table, "updated_user")
	w.Maintenance = field.NewString(table, "maintenance")
	w.OperatorName = field.NewString(table, "operator_name")
	w.TestI = field.NewInt64(table, "test_i")

	w.fillFieldMap()

	return w
}

func (w *winGameSlotCopy0526) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winGameSlotCopy0526) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 24)
	w.fieldMap["id"] = w.ID
	w.fieldMap["game_id"] = w.GameID
	w.fieldMap["game_group_id"] = w.GameGroupID
	w.fieldMap["plat_id"] = w.PlatID
	w.fieldMap["provider"] = w.Provider
	w.fieldMap["name"] = w.Name
	w.fieldMap["name_zh"] = w.NameZh
	w.fieldMap["img"] = w.Img
	w.fieldMap["img_new"] = w.ImgNew
	w.fieldMap["is_new"] = w.IsNew
	w.fieldMap["is_casino"] = w.IsCasino
	w.fieldMap["game_type_id"] = w.GameTypeID
	w.fieldMap["game_type_name"] = w.GameTypeName
	w.fieldMap["favorite_star"] = w.FavoriteStar
	w.fieldMap["hot_star"] = w.HotStar
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["status"] = w.Status
	w.fieldMap["device"] = w.Device
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["updated_user"] = w.UpdatedUser
	w.fieldMap["maintenance"] = w.Maintenance
	w.fieldMap["operator_name"] = w.OperatorName
	w.fieldMap["test_i"] = w.TestI
}

func (w winGameSlotCopy0526) clone(db *gorm.DB) winGameSlotCopy0526 {
	w.winGameSlotCopy0526Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winGameSlotCopy0526) replaceDB(db *gorm.DB) winGameSlotCopy0526 {
	w.winGameSlotCopy0526Do.ReplaceDB(db)
	return w
}

type winGameSlotCopy0526Do struct{ gen.DO }

type IWinGameSlotCopy0526Do interface {
	gen.SubQuery
	Debug() IWinGameSlotCopy0526Do
	WithContext(ctx context.Context) IWinGameSlotCopy0526Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinGameSlotCopy0526Do
	WriteDB() IWinGameSlotCopy0526Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinGameSlotCopy0526Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinGameSlotCopy0526Do
	Not(conds ...gen.Condition) IWinGameSlotCopy0526Do
	Or(conds ...gen.Condition) IWinGameSlotCopy0526Do
	Select(conds ...field.Expr) IWinGameSlotCopy0526Do
	Where(conds ...gen.Condition) IWinGameSlotCopy0526Do
	Order(conds ...field.Expr) IWinGameSlotCopy0526Do
	Distinct(cols ...field.Expr) IWinGameSlotCopy0526Do
	Omit(cols ...field.Expr) IWinGameSlotCopy0526Do
	Join(table schema.Tabler, on ...field.Expr) IWinGameSlotCopy0526Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinGameSlotCopy0526Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinGameSlotCopy0526Do
	Group(cols ...field.Expr) IWinGameSlotCopy0526Do
	Having(conds ...gen.Condition) IWinGameSlotCopy0526Do
	Limit(limit int) IWinGameSlotCopy0526Do
	Offset(offset int) IWinGameSlotCopy0526Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinGameSlotCopy0526Do
	Unscoped() IWinGameSlotCopy0526Do
	Create(values ...*model.WinGameSlotCopy0526) error
	CreateInBatches(values []*model.WinGameSlotCopy0526, batchSize int) error
	Save(values ...*model.WinGameSlotCopy0526) error
	First() (*model.WinGameSlotCopy0526, error)
	Take() (*model.WinGameSlotCopy0526, error)
	Last() (*model.WinGameSlotCopy0526, error)
	Find() ([]*model.WinGameSlotCopy0526, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinGameSlotCopy0526, err error)
	FindInBatches(result *[]*model.WinGameSlotCopy0526, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinGameSlotCopy0526) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinGameSlotCopy0526Do
	Assign(attrs ...field.AssignExpr) IWinGameSlotCopy0526Do
	Joins(fields ...field.RelationField) IWinGameSlotCopy0526Do
	Preload(fields ...field.RelationField) IWinGameSlotCopy0526Do
	FirstOrInit() (*model.WinGameSlotCopy0526, error)
	FirstOrCreate() (*model.WinGameSlotCopy0526, error)
	FindByPage(offset int, limit int) (result []*model.WinGameSlotCopy0526, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinGameSlotCopy0526Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winGameSlotCopy0526Do) Debug() IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Debug())
}

func (w winGameSlotCopy0526Do) WithContext(ctx context.Context) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winGameSlotCopy0526Do) ReadDB() IWinGameSlotCopy0526Do {
	return w.Clauses(dbresolver.Read)
}

func (w winGameSlotCopy0526Do) WriteDB() IWinGameSlotCopy0526Do {
	return w.Clauses(dbresolver.Write)
}

func (w winGameSlotCopy0526Do) Session(config *gorm.Session) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Session(config))
}

func (w winGameSlotCopy0526Do) Clauses(conds ...clause.Expression) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winGameSlotCopy0526Do) Returning(value interface{}, columns ...string) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winGameSlotCopy0526Do) Not(conds ...gen.Condition) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winGameSlotCopy0526Do) Or(conds ...gen.Condition) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winGameSlotCopy0526Do) Select(conds ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winGameSlotCopy0526Do) Where(conds ...gen.Condition) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winGameSlotCopy0526Do) Order(conds ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winGameSlotCopy0526Do) Distinct(cols ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winGameSlotCopy0526Do) Omit(cols ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winGameSlotCopy0526Do) Join(table schema.Tabler, on ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winGameSlotCopy0526Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winGameSlotCopy0526Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winGameSlotCopy0526Do) Group(cols ...field.Expr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winGameSlotCopy0526Do) Having(conds ...gen.Condition) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winGameSlotCopy0526Do) Limit(limit int) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winGameSlotCopy0526Do) Offset(offset int) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winGameSlotCopy0526Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winGameSlotCopy0526Do) Unscoped() IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winGameSlotCopy0526Do) Create(values ...*model.WinGameSlotCopy0526) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winGameSlotCopy0526Do) CreateInBatches(values []*model.WinGameSlotCopy0526, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winGameSlotCopy0526Do) Save(values ...*model.WinGameSlotCopy0526) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winGameSlotCopy0526Do) First() (*model.WinGameSlotCopy0526, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinGameSlotCopy0526), nil
	}
}

func (w winGameSlotCopy0526Do) Take() (*model.WinGameSlotCopy0526, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinGameSlotCopy0526), nil
	}
}

func (w winGameSlotCopy0526Do) Last() (*model.WinGameSlotCopy0526, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinGameSlotCopy0526), nil
	}
}

func (w winGameSlotCopy0526Do) Find() ([]*model.WinGameSlotCopy0526, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinGameSlotCopy0526), err
}

func (w winGameSlotCopy0526Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinGameSlotCopy0526, err error) {
	buf := make([]*model.WinGameSlotCopy0526, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winGameSlotCopy0526Do) FindInBatches(result *[]*model.WinGameSlotCopy0526, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winGameSlotCopy0526Do) Attrs(attrs ...field.AssignExpr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winGameSlotCopy0526Do) Assign(attrs ...field.AssignExpr) IWinGameSlotCopy0526Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winGameSlotCopy0526Do) Joins(fields ...field.RelationField) IWinGameSlotCopy0526Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winGameSlotCopy0526Do) Preload(fields ...field.RelationField) IWinGameSlotCopy0526Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winGameSlotCopy0526Do) FirstOrInit() (*model.WinGameSlotCopy0526, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinGameSlotCopy0526), nil
	}
}

func (w winGameSlotCopy0526Do) FirstOrCreate() (*model.WinGameSlotCopy0526, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinGameSlotCopy0526), nil
	}
}

func (w winGameSlotCopy0526Do) FindByPage(offset int, limit int) (result []*model.WinGameSlotCopy0526, count int64, err error) {
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

func (w winGameSlotCopy0526Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winGameSlotCopy0526Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winGameSlotCopy0526Do) Delete(models ...*model.WinGameSlotCopy0526) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winGameSlotCopy0526Do) withDO(do gen.Dao) *winGameSlotCopy0526Do {
	w.DO = *do.(*gen.DO)
	return w
}
