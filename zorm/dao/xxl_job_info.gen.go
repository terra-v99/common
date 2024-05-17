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

func newXxlJobInfo(db *gorm.DB, opts ...gen.DOOption) xxlJobInfo {
	_xxlJobInfo := xxlJobInfo{}

	_xxlJobInfo.xxlJobInfoDo.UseDB(db, opts...)
	_xxlJobInfo.xxlJobInfoDo.UseModel(&model.XxlJobInfo{})

	tableName := _xxlJobInfo.xxlJobInfoDo.TableName()
	_xxlJobInfo.ALL = field.NewAsterisk(tableName)
	_xxlJobInfo.ID = field.NewInt64(tableName, "id")
	_xxlJobInfo.JobGroup = field.NewInt64(tableName, "job_group")
	_xxlJobInfo.JobDesc = field.NewString(tableName, "job_desc")
	_xxlJobInfo.AddTime = field.NewTime(tableName, "add_time")
	_xxlJobInfo.UpdateTime = field.NewTime(tableName, "update_time")
	_xxlJobInfo.Author = field.NewString(tableName, "author")
	_xxlJobInfo.AlarmEmail = field.NewString(tableName, "alarm_email")
	_xxlJobInfo.ScheduleType = field.NewString(tableName, "schedule_type")
	_xxlJobInfo.ScheduleConf = field.NewString(tableName, "schedule_conf")
	_xxlJobInfo.MisfireStrategy = field.NewString(tableName, "misfire_strategy")
	_xxlJobInfo.ExecutorRouteStrategy = field.NewString(tableName, "executor_route_strategy")
	_xxlJobInfo.ExecutorHandler = field.NewString(tableName, "executor_handler")
	_xxlJobInfo.ExecutorParam = field.NewString(tableName, "executor_param")
	_xxlJobInfo.ExecutorBlockStrategy = field.NewString(tableName, "executor_block_strategy")
	_xxlJobInfo.ExecutorTimeout = field.NewInt64(tableName, "executor_timeout")
	_xxlJobInfo.ExecutorFailRetryCount = field.NewInt64(tableName, "executor_fail_retry_count")
	_xxlJobInfo.GlueType = field.NewString(tableName, "glue_type")
	_xxlJobInfo.GlueSource = field.NewString(tableName, "glue_source")
	_xxlJobInfo.GlueRemark = field.NewString(tableName, "glue_remark")
	_xxlJobInfo.GlueUpdatetime = field.NewTime(tableName, "glue_updatetime")
	_xxlJobInfo.ChildJobid = field.NewString(tableName, "child_jobid")
	_xxlJobInfo.TriggerStatus = field.NewInt64(tableName, "trigger_status")
	_xxlJobInfo.TriggerLastTime = field.NewInt64(tableName, "trigger_last_time")
	_xxlJobInfo.TriggerNextTime = field.NewInt64(tableName, "trigger_next_time")

	_xxlJobInfo.fillFieldMap()

	return _xxlJobInfo
}

type xxlJobInfo struct {
	xxlJobInfoDo

	ALL                    field.Asterisk
	ID                     field.Int64
	JobGroup               field.Int64 // 执行器主键ID
	JobDesc                field.String
	AddTime                field.Time
	UpdateTime             field.Time
	Author                 field.String // 作者
	AlarmEmail             field.String // 报警邮件
	ScheduleType           field.String // 调度类型
	ScheduleConf           field.String // 调度配置，值含义取决于调度类型
	MisfireStrategy        field.String // 调度过期策略
	ExecutorRouteStrategy  field.String // 执行器路由策略
	ExecutorHandler        field.String // 执行器任务handler
	ExecutorParam          field.String // 执行器任务参数
	ExecutorBlockStrategy  field.String // 阻塞处理策略
	ExecutorTimeout        field.Int64  // 任务执行超时时间，单位秒
	ExecutorFailRetryCount field.Int64  // 失败重试次数
	GlueType               field.String // GLUE类型
	GlueSource             field.String // GLUE源代码
	GlueRemark             field.String // GLUE备注
	GlueUpdatetime         field.Time   // GLUE更新时间
	ChildJobid             field.String // 子任务ID，多个逗号分隔
	TriggerStatus          field.Int64  // 调度状态：0-停止，1-运行
	TriggerLastTime        field.Int64  // 上次调度时间
	TriggerNextTime        field.Int64  // 下次调度时间

	fieldMap map[string]field.Expr
}

func (x xxlJobInfo) Table(newTableName string) *xxlJobInfo {
	x.xxlJobInfoDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xxlJobInfo) As(alias string) *xxlJobInfo {
	x.xxlJobInfoDo.DO = *(x.xxlJobInfoDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xxlJobInfo) updateTableName(table string) *xxlJobInfo {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.JobGroup = field.NewInt64(table, "job_group")
	x.JobDesc = field.NewString(table, "job_desc")
	x.AddTime = field.NewTime(table, "add_time")
	x.UpdateTime = field.NewTime(table, "update_time")
	x.Author = field.NewString(table, "author")
	x.AlarmEmail = field.NewString(table, "alarm_email")
	x.ScheduleType = field.NewString(table, "schedule_type")
	x.ScheduleConf = field.NewString(table, "schedule_conf")
	x.MisfireStrategy = field.NewString(table, "misfire_strategy")
	x.ExecutorRouteStrategy = field.NewString(table, "executor_route_strategy")
	x.ExecutorHandler = field.NewString(table, "executor_handler")
	x.ExecutorParam = field.NewString(table, "executor_param")
	x.ExecutorBlockStrategy = field.NewString(table, "executor_block_strategy")
	x.ExecutorTimeout = field.NewInt64(table, "executor_timeout")
	x.ExecutorFailRetryCount = field.NewInt64(table, "executor_fail_retry_count")
	x.GlueType = field.NewString(table, "glue_type")
	x.GlueSource = field.NewString(table, "glue_source")
	x.GlueRemark = field.NewString(table, "glue_remark")
	x.GlueUpdatetime = field.NewTime(table, "glue_updatetime")
	x.ChildJobid = field.NewString(table, "child_jobid")
	x.TriggerStatus = field.NewInt64(table, "trigger_status")
	x.TriggerLastTime = field.NewInt64(table, "trigger_last_time")
	x.TriggerNextTime = field.NewInt64(table, "trigger_next_time")

	x.fillFieldMap()

	return x
}

func (x *xxlJobInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xxlJobInfo) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 24)
	x.fieldMap["id"] = x.ID
	x.fieldMap["job_group"] = x.JobGroup
	x.fieldMap["job_desc"] = x.JobDesc
	x.fieldMap["add_time"] = x.AddTime
	x.fieldMap["update_time"] = x.UpdateTime
	x.fieldMap["author"] = x.Author
	x.fieldMap["alarm_email"] = x.AlarmEmail
	x.fieldMap["schedule_type"] = x.ScheduleType
	x.fieldMap["schedule_conf"] = x.ScheduleConf
	x.fieldMap["misfire_strategy"] = x.MisfireStrategy
	x.fieldMap["executor_route_strategy"] = x.ExecutorRouteStrategy
	x.fieldMap["executor_handler"] = x.ExecutorHandler
	x.fieldMap["executor_param"] = x.ExecutorParam
	x.fieldMap["executor_block_strategy"] = x.ExecutorBlockStrategy
	x.fieldMap["executor_timeout"] = x.ExecutorTimeout
	x.fieldMap["executor_fail_retry_count"] = x.ExecutorFailRetryCount
	x.fieldMap["glue_type"] = x.GlueType
	x.fieldMap["glue_source"] = x.GlueSource
	x.fieldMap["glue_remark"] = x.GlueRemark
	x.fieldMap["glue_updatetime"] = x.GlueUpdatetime
	x.fieldMap["child_jobid"] = x.ChildJobid
	x.fieldMap["trigger_status"] = x.TriggerStatus
	x.fieldMap["trigger_last_time"] = x.TriggerLastTime
	x.fieldMap["trigger_next_time"] = x.TriggerNextTime
}

func (x xxlJobInfo) clone(db *gorm.DB) xxlJobInfo {
	x.xxlJobInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xxlJobInfo) replaceDB(db *gorm.DB) xxlJobInfo {
	x.xxlJobInfoDo.ReplaceDB(db)
	return x
}

type xxlJobInfoDo struct{ gen.DO }

type IXxlJobInfoDo interface {
	gen.SubQuery
	Debug() IXxlJobInfoDo
	WithContext(ctx context.Context) IXxlJobInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXxlJobInfoDo
	WriteDB() IXxlJobInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXxlJobInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXxlJobInfoDo
	Not(conds ...gen.Condition) IXxlJobInfoDo
	Or(conds ...gen.Condition) IXxlJobInfoDo
	Select(conds ...field.Expr) IXxlJobInfoDo
	Where(conds ...gen.Condition) IXxlJobInfoDo
	Order(conds ...field.Expr) IXxlJobInfoDo
	Distinct(cols ...field.Expr) IXxlJobInfoDo
	Omit(cols ...field.Expr) IXxlJobInfoDo
	Join(table schema.Tabler, on ...field.Expr) IXxlJobInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobInfoDo
	Group(cols ...field.Expr) IXxlJobInfoDo
	Having(conds ...gen.Condition) IXxlJobInfoDo
	Limit(limit int) IXxlJobInfoDo
	Offset(offset int) IXxlJobInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobInfoDo
	Unscoped() IXxlJobInfoDo
	Create(values ...*model.XxlJobInfo) error
	CreateInBatches(values []*model.XxlJobInfo, batchSize int) error
	Save(values ...*model.XxlJobInfo) error
	First() (*model.XxlJobInfo, error)
	Take() (*model.XxlJobInfo, error)
	Last() (*model.XxlJobInfo, error)
	Find() ([]*model.XxlJobInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobInfo, err error)
	FindInBatches(result *[]*model.XxlJobInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XxlJobInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXxlJobInfoDo
	Assign(attrs ...field.AssignExpr) IXxlJobInfoDo
	Joins(fields ...field.RelationField) IXxlJobInfoDo
	Preload(fields ...field.RelationField) IXxlJobInfoDo
	FirstOrInit() (*model.XxlJobInfo, error)
	FirstOrCreate() (*model.XxlJobInfo, error)
	FindByPage(offset int, limit int) (result []*model.XxlJobInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXxlJobInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xxlJobInfoDo) Debug() IXxlJobInfoDo {
	return x.withDO(x.DO.Debug())
}

func (x xxlJobInfoDo) WithContext(ctx context.Context) IXxlJobInfoDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xxlJobInfoDo) ReadDB() IXxlJobInfoDo {
	return x.Clauses(dbresolver.Read)
}

func (x xxlJobInfoDo) WriteDB() IXxlJobInfoDo {
	return x.Clauses(dbresolver.Write)
}

func (x xxlJobInfoDo) Session(config *gorm.Session) IXxlJobInfoDo {
	return x.withDO(x.DO.Session(config))
}

func (x xxlJobInfoDo) Clauses(conds ...clause.Expression) IXxlJobInfoDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xxlJobInfoDo) Returning(value interface{}, columns ...string) IXxlJobInfoDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xxlJobInfoDo) Not(conds ...gen.Condition) IXxlJobInfoDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xxlJobInfoDo) Or(conds ...gen.Condition) IXxlJobInfoDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xxlJobInfoDo) Select(conds ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xxlJobInfoDo) Where(conds ...gen.Condition) IXxlJobInfoDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xxlJobInfoDo) Order(conds ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xxlJobInfoDo) Distinct(cols ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xxlJobInfoDo) Omit(cols ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xxlJobInfoDo) Join(table schema.Tabler, on ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xxlJobInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xxlJobInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xxlJobInfoDo) Group(cols ...field.Expr) IXxlJobInfoDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xxlJobInfoDo) Having(conds ...gen.Condition) IXxlJobInfoDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xxlJobInfoDo) Limit(limit int) IXxlJobInfoDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xxlJobInfoDo) Offset(offset int) IXxlJobInfoDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xxlJobInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobInfoDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xxlJobInfoDo) Unscoped() IXxlJobInfoDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xxlJobInfoDo) Create(values ...*model.XxlJobInfo) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xxlJobInfoDo) CreateInBatches(values []*model.XxlJobInfo, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xxlJobInfoDo) Save(values ...*model.XxlJobInfo) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xxlJobInfoDo) First() (*model.XxlJobInfo, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobInfo), nil
	}
}

func (x xxlJobInfoDo) Take() (*model.XxlJobInfo, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobInfo), nil
	}
}

func (x xxlJobInfoDo) Last() (*model.XxlJobInfo, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobInfo), nil
	}
}

func (x xxlJobInfoDo) Find() ([]*model.XxlJobInfo, error) {
	result, err := x.DO.Find()
	return result.([]*model.XxlJobInfo), err
}

func (x xxlJobInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobInfo, err error) {
	buf := make([]*model.XxlJobInfo, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xxlJobInfoDo) FindInBatches(result *[]*model.XxlJobInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xxlJobInfoDo) Attrs(attrs ...field.AssignExpr) IXxlJobInfoDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xxlJobInfoDo) Assign(attrs ...field.AssignExpr) IXxlJobInfoDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xxlJobInfoDo) Joins(fields ...field.RelationField) IXxlJobInfoDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xxlJobInfoDo) Preload(fields ...field.RelationField) IXxlJobInfoDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xxlJobInfoDo) FirstOrInit() (*model.XxlJobInfo, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobInfo), nil
	}
}

func (x xxlJobInfoDo) FirstOrCreate() (*model.XxlJobInfo, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobInfo), nil
	}
}

func (x xxlJobInfoDo) FindByPage(offset int, limit int) (result []*model.XxlJobInfo, count int64, err error) {
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

func (x xxlJobInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xxlJobInfoDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xxlJobInfoDo) Delete(models ...*model.XxlJobInfo) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xxlJobInfoDo) withDO(do gen.Dao) *xxlJobInfoDo {
	x.DO = *do.(*gen.DO)
	return x
}
