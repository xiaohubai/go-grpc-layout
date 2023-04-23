// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

func newCasbinRule(db *gorm.DB, opts ...gen.DOOption) casbinRule {
	_casbinRule := casbinRule{}

	_casbinRule.casbinRuleDo.UseDB(db, opts...)
	_casbinRule.casbinRuleDo.UseModel(&model.CasbinRule{})

	tableName := _casbinRule.casbinRuleDo.TableName()
	_casbinRule.ALL = field.NewAsterisk(tableName)
	_casbinRule.ID = field.NewInt64(tableName, "id")
	_casbinRule.Ptype = field.NewString(tableName, "ptype")
	_casbinRule.V0 = field.NewString(tableName, "v0")
	_casbinRule.V1 = field.NewString(tableName, "v1")
	_casbinRule.V2 = field.NewString(tableName, "v2")
	_casbinRule.V3 = field.NewString(tableName, "v3")
	_casbinRule.V4 = field.NewString(tableName, "v4")
	_casbinRule.V5 = field.NewString(tableName, "v5")
	_casbinRule.Desc = field.NewString(tableName, "desc")
	_casbinRule.CreateAt = field.NewTime(tableName, "create_at")
	_casbinRule.UpdateAt = field.NewTime(tableName, "update_at")
	_casbinRule.DeletedAt = field.NewField(tableName, "deleted_at")
	_casbinRule.CreateUser = field.NewString(tableName, "create_user")
	_casbinRule.UpdateUser = field.NewString(tableName, "update_user")

	_casbinRule.fillFieldMap()

	return _casbinRule
}

type casbinRule struct {
	casbinRuleDo casbinRuleDo

	ALL        field.Asterisk
	ID         field.Int64
	Ptype      field.String
	V0         field.String
	V1         field.String
	V2         field.String
	V3         field.String
	V4         field.String
	V5         field.String
	Desc       field.String // 描述
	CreateAt   field.Time   // 记录创建时间
	UpdateAt   field.Time   // 记录修改时间
	DeletedAt  field.Field  // 删除时间
	CreateUser field.String // 创建人
	UpdateUser field.String // 修改人

	fieldMap map[string]field.Expr
}

func (c casbinRule) Table(newTableName string) *casbinRule {
	c.casbinRuleDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c casbinRule) As(alias string) *casbinRule {
	c.casbinRuleDo.DO = *(c.casbinRuleDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *casbinRule) updateTableName(table string) *casbinRule {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Ptype = field.NewString(table, "ptype")
	c.V0 = field.NewString(table, "v0")
	c.V1 = field.NewString(table, "v1")
	c.V2 = field.NewString(table, "v2")
	c.V3 = field.NewString(table, "v3")
	c.V4 = field.NewString(table, "v4")
	c.V5 = field.NewString(table, "v5")
	c.Desc = field.NewString(table, "desc")
	c.CreateAt = field.NewTime(table, "create_at")
	c.UpdateAt = field.NewTime(table, "update_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.CreateUser = field.NewString(table, "create_user")
	c.UpdateUser = field.NewString(table, "update_user")

	c.fillFieldMap()

	return c
}

func (c *casbinRule) WithContext(ctx context.Context) *casbinRuleDo {
	return c.casbinRuleDo.WithContext(ctx)
}

func (c casbinRule) TableName() string { return c.casbinRuleDo.TableName() }

func (c casbinRule) Alias() string { return c.casbinRuleDo.Alias() }

func (c *casbinRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *casbinRule) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 14)
	c.fieldMap["id"] = c.ID
	c.fieldMap["ptype"] = c.Ptype
	c.fieldMap["v0"] = c.V0
	c.fieldMap["v1"] = c.V1
	c.fieldMap["v2"] = c.V2
	c.fieldMap["v3"] = c.V3
	c.fieldMap["v4"] = c.V4
	c.fieldMap["v5"] = c.V5
	c.fieldMap["desc"] = c.Desc
	c.fieldMap["create_at"] = c.CreateAt
	c.fieldMap["update_at"] = c.UpdateAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["create_user"] = c.CreateUser
	c.fieldMap["update_user"] = c.UpdateUser
}

func (c casbinRule) clone(db *gorm.DB) casbinRule {
	c.casbinRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c casbinRule) replaceDB(db *gorm.DB) casbinRule {
	c.casbinRuleDo.ReplaceDB(db)
	return c
}

type casbinRuleDo struct{ gen.DO }

func (c casbinRuleDo) Debug() *casbinRuleDo {
	return c.withDO(c.DO.Debug())
}

func (c casbinRuleDo) WithContext(ctx context.Context) *casbinRuleDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c casbinRuleDo) ReadDB() *casbinRuleDo {
	return c.Clauses(dbresolver.Read)
}

func (c casbinRuleDo) WriteDB() *casbinRuleDo {
	return c.Clauses(dbresolver.Write)
}

func (c casbinRuleDo) Session(config *gorm.Session) *casbinRuleDo {
	return c.withDO(c.DO.Session(config))
}

func (c casbinRuleDo) Clauses(conds ...clause.Expression) *casbinRuleDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c casbinRuleDo) Returning(value interface{}, columns ...string) *casbinRuleDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c casbinRuleDo) Not(conds ...gen.Condition) *casbinRuleDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c casbinRuleDo) Or(conds ...gen.Condition) *casbinRuleDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c casbinRuleDo) Select(conds ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c casbinRuleDo) Where(conds ...gen.Condition) *casbinRuleDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c casbinRuleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *casbinRuleDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c casbinRuleDo) Order(conds ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c casbinRuleDo) Distinct(cols ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c casbinRuleDo) Omit(cols ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c casbinRuleDo) Join(table schema.Tabler, on ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c casbinRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c casbinRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c casbinRuleDo) Group(cols ...field.Expr) *casbinRuleDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c casbinRuleDo) Having(conds ...gen.Condition) *casbinRuleDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c casbinRuleDo) Limit(limit int) *casbinRuleDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c casbinRuleDo) Offset(offset int) *casbinRuleDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c casbinRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *casbinRuleDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c casbinRuleDo) Unscoped() *casbinRuleDo {
	return c.withDO(c.DO.Unscoped())
}

func (c casbinRuleDo) Create(values ...*model.CasbinRule) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c casbinRuleDo) CreateInBatches(values []*model.CasbinRule, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c casbinRuleDo) Save(values ...*model.CasbinRule) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c casbinRuleDo) First() (*model.CasbinRule, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CasbinRule), nil
	}
}

func (c casbinRuleDo) Take() (*model.CasbinRule, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CasbinRule), nil
	}
}

func (c casbinRuleDo) Last() (*model.CasbinRule, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CasbinRule), nil
	}
}

func (c casbinRuleDo) Find() ([]*model.CasbinRule, error) {
	result, err := c.DO.Find()
	return result.([]*model.CasbinRule), err
}

func (c casbinRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CasbinRule, err error) {
	buf := make([]*model.CasbinRule, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c casbinRuleDo) FindInBatches(result *[]*model.CasbinRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c casbinRuleDo) Attrs(attrs ...field.AssignExpr) *casbinRuleDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c casbinRuleDo) Assign(attrs ...field.AssignExpr) *casbinRuleDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c casbinRuleDo) Joins(fields ...field.RelationField) *casbinRuleDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c casbinRuleDo) Preload(fields ...field.RelationField) *casbinRuleDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c casbinRuleDo) FirstOrInit() (*model.CasbinRule, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CasbinRule), nil
	}
}

func (c casbinRuleDo) FirstOrCreate() (*model.CasbinRule, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CasbinRule), nil
	}
}

func (c casbinRuleDo) FindByPage(offset int, limit int) (result []*model.CasbinRule, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c casbinRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c casbinRuleDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c casbinRuleDo) Delete(models ...*model.CasbinRule) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *casbinRuleDo) withDO(do gen.Dao) *casbinRuleDo {
	c.DO = *do.(*gen.DO)
	return c
}
