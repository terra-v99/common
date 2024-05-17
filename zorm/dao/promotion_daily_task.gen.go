// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"gitlab.skig.tech/zero-core/common/zorm/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPromotionDailyTask(db *gorm.DB, opts ...gen.DOOption) promotionDailyTask {
	_promotionDailyTask := promotionDailyTask{}

	_promotionDailyTask.promotionDailyTaskDo.UseDB(db, opts...)
	_promotionDailyTask.promotionDailyTaskDo.UseModel(&model.PromotionDailyTask{})

	tableName := _promotionDailyTask.promotionDailyTaskDo.TableName()
	_promotionDailyTask.ALL = field.NewAsterisk(tableName)
	_promotionDailyTask.ID = field.NewInt64(tableName, "id")
	_promotionDailyTask.PromotionsID = field.NewInt64(tableName, "promotions_id")
	_promotionDailyTask.UserID = field.NewInt64(tableName, "user_id")
	_promotionDailyTask.TaskTitle = field.NewString(tableName, "task_title")
	_promotionDailyTask.BetAmount = field.NewInt64(tableName, "bet_amount")
	_promotionDailyTask.Bonus = field.NewInt64(tableName, "bonus")
	_promotionDailyTask.CodeMultiple = field.NewInt64(tableName, "code_multiple")
	_promotionDailyTask.RotatingNumber = field.NewInt64(tableName, "rotating_number")
	_promotionDailyTask.PerBetAmount = field.NewField(tableName, "per_bet_amount")
	_promotionDailyTask.ValidDays = field.NewInt64(tableName, "valid_days")
	_promotionDailyTask.SortID = field.NewInt64(tableName, "sort_id")
	_promotionDailyTask.Status = field.NewInt64(tableName, "status")
	_promotionDailyTask.TaskDate = field.NewInt64(tableName, "task_date")
	_promotionDailyTask.UpdateAt = field.NewInt64(tableName, "update_at")

	_promotionDailyTask.fillFieldMap()

	return _promotionDailyTask
}

type promotionDailyTask struct {
	promotionDailyTaskDo

	ALL            field.Asterisk
	ID             field.Int64  // 主键ID
	PromotionsID   field.Int64  // 优惠活动ID
	UserID         field.Int64  // 用户ID
	TaskTitle      field.String // 任务标题: task1, task2
	BetAmount      field.Int64  // 有效投注金额
	Bonus          field.Int64  // 彩金金额
	CodeMultiple   field.Int64  // 打码倍数
	RotatingNumber field.Int64  // 旋转数量
	PerBetAmount   field.Field  // 单次投注金额
	ValidDays      field.Int64  // 有效期天数
	SortID         field.Int64  // 任务序号ID: 1-6
	Status         field.Int64  // 任务状态: 0-未完成, 1-已完成
	TaskDate       field.Int64  // 任务日期
	UpdateAt       field.Int64  // 修改时间

	fieldMap map[string]field.Expr
}

func (p promotionDailyTask) Table(newTableName string) *promotionDailyTask {
	p.promotionDailyTaskDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promotionDailyTask) As(alias string) *promotionDailyTask {
	p.promotionDailyTaskDo.DO = *(p.promotionDailyTaskDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promotionDailyTask) updateTableName(table string) *promotionDailyTask {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.PromotionsID = field.NewInt64(table, "promotions_id")
	p.UserID = field.NewInt64(table, "user_id")
	p.TaskTitle = field.NewString(table, "task_title")
	p.BetAmount = field.NewInt64(table, "bet_amount")
	p.Bonus = field.NewInt64(table, "bonus")
	p.CodeMultiple = field.NewInt64(table, "code_multiple")
	p.RotatingNumber = field.NewInt64(table, "rotating_number")
	p.PerBetAmount = field.NewField(table, "per_bet_amount")
	p.ValidDays = field.NewInt64(table, "valid_days")
	p.SortID = field.NewInt64(table, "sort_id")
	p.Status = field.NewInt64(table, "status")
	p.TaskDate = field.NewInt64(table, "task_date")
	p.UpdateAt = field.NewInt64(table, "update_at")

	p.fillFieldMap()

	return p
}

func (p *promotionDailyTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promotionDailyTask) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 14)
	p.fieldMap["id"] = p.ID
	p.fieldMap["promotions_id"] = p.PromotionsID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["task_title"] = p.TaskTitle
	p.fieldMap["bet_amount"] = p.BetAmount
	p.fieldMap["bonus"] = p.Bonus
	p.fieldMap["code_multiple"] = p.CodeMultiple
	p.fieldMap["rotating_number"] = p.RotatingNumber
	p.fieldMap["per_bet_amount"] = p.PerBetAmount
	p.fieldMap["valid_days"] = p.ValidDays
	p.fieldMap["sort_id"] = p.SortID
	p.fieldMap["status"] = p.Status
	p.fieldMap["task_date"] = p.TaskDate
	p.fieldMap["update_at"] = p.UpdateAt
}

func (p promotionDailyTask) clone(db *gorm.DB) promotionDailyTask {
	p.promotionDailyTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promotionDailyTask) replaceDB(db *gorm.DB) promotionDailyTask {
	p.promotionDailyTaskDo.ReplaceDB(db)
	return p
}

type promotionDailyTaskDo struct{ gen.DO }

type IPromotionDailyTaskDo interface {
	gen.SubQuery
	Debug() IPromotionDailyTaskDo
	WithContext(ctx context.Context) IPromotionDailyTaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPromotionDailyTaskDo
	WriteDB() IPromotionDailyTaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPromotionDailyTaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPromotionDailyTaskDo
	Not(conds ...gen.Condition) IPromotionDailyTaskDo
	Or(conds ...gen.Condition) IPromotionDailyTaskDo
	Select(conds ...field.Expr) IPromotionDailyTaskDo
	Where(conds ...gen.Condition) IPromotionDailyTaskDo
	Order(conds ...field.Expr) IPromotionDailyTaskDo
	Distinct(cols ...field.Expr) IPromotionDailyTaskDo
	Omit(cols ...field.Expr) IPromotionDailyTaskDo
	Join(table schema.Tabler, on ...field.Expr) IPromotionDailyTaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPromotionDailyTaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPromotionDailyTaskDo
	Group(cols ...field.Expr) IPromotionDailyTaskDo
	Having(conds ...gen.Condition) IPromotionDailyTaskDo
	Limit(limit int) IPromotionDailyTaskDo
	Offset(offset int) IPromotionDailyTaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPromotionDailyTaskDo
	Unscoped() IPromotionDailyTaskDo
	Create(values ...*model.PromotionDailyTask) error
	CreateInBatches(values []*model.PromotionDailyTask, batchSize int) error
	Save(values ...*model.PromotionDailyTask) error
	First() (*model.PromotionDailyTask, error)
	Take() (*model.PromotionDailyTask, error)
	Last() (*model.PromotionDailyTask, error)
	Find() ([]*model.PromotionDailyTask, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromotionDailyTask, err error)
	FindInBatches(result *[]*model.PromotionDailyTask, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PromotionDailyTask) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPromotionDailyTaskDo
	Assign(attrs ...field.AssignExpr) IPromotionDailyTaskDo
	Joins(fields ...field.RelationField) IPromotionDailyTaskDo
	Preload(fields ...field.RelationField) IPromotionDailyTaskDo
	FirstOrInit() (*model.PromotionDailyTask, error)
	FirstOrCreate() (*model.PromotionDailyTask, error)
	FindByPage(offset int, limit int) (result []*model.PromotionDailyTask, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPromotionDailyTaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p promotionDailyTaskDo) Debug() IPromotionDailyTaskDo {
	return p.withDO(p.DO.Debug())
}

func (p promotionDailyTaskDo) WithContext(ctx context.Context) IPromotionDailyTaskDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promotionDailyTaskDo) ReadDB() IPromotionDailyTaskDo {
	return p.Clauses(dbresolver.Read)
}

func (p promotionDailyTaskDo) WriteDB() IPromotionDailyTaskDo {
	return p.Clauses(dbresolver.Write)
}

func (p promotionDailyTaskDo) Session(config *gorm.Session) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Session(config))
}

func (p promotionDailyTaskDo) Clauses(conds ...clause.Expression) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promotionDailyTaskDo) Returning(value interface{}, columns ...string) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promotionDailyTaskDo) Not(conds ...gen.Condition) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promotionDailyTaskDo) Or(conds ...gen.Condition) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promotionDailyTaskDo) Select(conds ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promotionDailyTaskDo) Where(conds ...gen.Condition) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promotionDailyTaskDo) Order(conds ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promotionDailyTaskDo) Distinct(cols ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promotionDailyTaskDo) Omit(cols ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promotionDailyTaskDo) Join(table schema.Tabler, on ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promotionDailyTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promotionDailyTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promotionDailyTaskDo) Group(cols ...field.Expr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promotionDailyTaskDo) Having(conds ...gen.Condition) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promotionDailyTaskDo) Limit(limit int) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promotionDailyTaskDo) Offset(offset int) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promotionDailyTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promotionDailyTaskDo) Unscoped() IPromotionDailyTaskDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promotionDailyTaskDo) Create(values ...*model.PromotionDailyTask) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promotionDailyTaskDo) CreateInBatches(values []*model.PromotionDailyTask, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promotionDailyTaskDo) Save(values ...*model.PromotionDailyTask) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promotionDailyTaskDo) First() (*model.PromotionDailyTask, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromotionDailyTask), nil
	}
}

func (p promotionDailyTaskDo) Take() (*model.PromotionDailyTask, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromotionDailyTask), nil
	}
}

func (p promotionDailyTaskDo) Last() (*model.PromotionDailyTask, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromotionDailyTask), nil
	}
}

func (p promotionDailyTaskDo) Find() ([]*model.PromotionDailyTask, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromotionDailyTask), err
}

func (p promotionDailyTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromotionDailyTask, err error) {
	buf := make([]*model.PromotionDailyTask, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promotionDailyTaskDo) FindInBatches(result *[]*model.PromotionDailyTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promotionDailyTaskDo) Attrs(attrs ...field.AssignExpr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promotionDailyTaskDo) Assign(attrs ...field.AssignExpr) IPromotionDailyTaskDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promotionDailyTaskDo) Joins(fields ...field.RelationField) IPromotionDailyTaskDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promotionDailyTaskDo) Preload(fields ...field.RelationField) IPromotionDailyTaskDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promotionDailyTaskDo) FirstOrInit() (*model.PromotionDailyTask, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromotionDailyTask), nil
	}
}

func (p promotionDailyTaskDo) FirstOrCreate() (*model.PromotionDailyTask, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromotionDailyTask), nil
	}
}

func (p promotionDailyTaskDo) FindByPage(offset int, limit int) (result []*model.PromotionDailyTask, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p promotionDailyTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promotionDailyTaskDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promotionDailyTaskDo) Delete(models ...*model.PromotionDailyTask) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promotionDailyTaskDo) withDO(do gen.Dao) *promotionDailyTaskDo {
	p.DO = *do.(*gen.DO)
	return p
}