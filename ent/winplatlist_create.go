// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/terra-v99/common/ent/winplatlist"
)

// WinPlatListCreate is the builder for creating a WinPlatList entity.
type WinPlatListCreate struct {
	config
	mutation *WinPlatListMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (wplc *WinPlatListCreate) SetCode(s string) *WinPlatListCreate {
	wplc.mutation.SetCode(s)
	return wplc
}

// SetName sets the "name" field.
func (wplc *WinPlatListCreate) SetName(s string) *WinPlatListCreate {
	wplc.mutation.SetName(s)
	return wplc
}

// SetConfig sets the "config" field.
func (wplc *WinPlatListCreate) SetConfig(s string) *WinPlatListCreate {
	wplc.mutation.SetConfig(s)
	return wplc
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (wplc *WinPlatListCreate) SetNillableConfig(s *string) *WinPlatListCreate {
	if s != nil {
		wplc.SetConfig(*s)
	}
	return wplc
}

// SetRate sets the "rate" field.
func (wplc *WinPlatListCreate) SetRate(s string) *WinPlatListCreate {
	wplc.mutation.SetRate(s)
	return wplc
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (wplc *WinPlatListCreate) SetNillableRate(s *string) *WinPlatListCreate {
	if s != nil {
		wplc.SetRate(*s)
	}
	return wplc
}

// SetSort sets the "sort" field.
func (wplc *WinPlatListCreate) SetSort(i int8) *WinPlatListCreate {
	wplc.mutation.SetSort(i)
	return wplc
}

// SetStatus sets the "status" field.
func (wplc *WinPlatListCreate) SetStatus(i int8) *WinPlatListCreate {
	wplc.mutation.SetStatus(i)
	return wplc
}

// SetCreatedAt sets the "created_at" field.
func (wplc *WinPlatListCreate) SetCreatedAt(i int32) *WinPlatListCreate {
	wplc.mutation.SetCreatedAt(i)
	return wplc
}

// SetUpdatedAt sets the "updated_at" field.
func (wplc *WinPlatListCreate) SetUpdatedAt(i int32) *WinPlatListCreate {
	wplc.mutation.SetUpdatedAt(i)
	return wplc
}

// SetID sets the "id" field.
func (wplc *WinPlatListCreate) SetID(i int32) *WinPlatListCreate {
	wplc.mutation.SetID(i)
	return wplc
}

// Mutation returns the WinPlatListMutation object of the builder.
func (wplc *WinPlatListCreate) Mutation() *WinPlatListMutation {
	return wplc.mutation
}

// Save creates the WinPlatList in the database.
func (wplc *WinPlatListCreate) Save(ctx context.Context) (*WinPlatList, error) {
	return withHooks(ctx, wplc.sqlSave, wplc.mutation, wplc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wplc *WinPlatListCreate) SaveX(ctx context.Context) *WinPlatList {
	v, err := wplc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wplc *WinPlatListCreate) Exec(ctx context.Context) error {
	_, err := wplc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wplc *WinPlatListCreate) ExecX(ctx context.Context) {
	if err := wplc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wplc *WinPlatListCreate) check() error {
	if _, ok := wplc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "WinPlatList.code"`)}
	}
	if _, ok := wplc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WinPlatList.name"`)}
	}
	if _, ok := wplc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "WinPlatList.sort"`)}
	}
	if _, ok := wplc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "WinPlatList.status"`)}
	}
	if _, ok := wplc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WinPlatList.created_at"`)}
	}
	if _, ok := wplc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WinPlatList.updated_at"`)}
	}
	return nil
}

func (wplc *WinPlatListCreate) sqlSave(ctx context.Context) (*WinPlatList, error) {
	if err := wplc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wplc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wplc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	wplc.mutation.id = &_node.ID
	wplc.mutation.done = true
	return _node, nil
}

func (wplc *WinPlatListCreate) createSpec() (*WinPlatList, *sqlgraph.CreateSpec) {
	var (
		_node = &WinPlatList{config: wplc.config}
		_spec = sqlgraph.NewCreateSpec(winplatlist.Table, sqlgraph.NewFieldSpec(winplatlist.FieldID, field.TypeInt32))
	)
	if id, ok := wplc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wplc.mutation.Code(); ok {
		_spec.SetField(winplatlist.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := wplc.mutation.Name(); ok {
		_spec.SetField(winplatlist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wplc.mutation.Config(); ok {
		_spec.SetField(winplatlist.FieldConfig, field.TypeString, value)
		_node.Config = value
	}
	if value, ok := wplc.mutation.Rate(); ok {
		_spec.SetField(winplatlist.FieldRate, field.TypeString, value)
		_node.Rate = value
	}
	if value, ok := wplc.mutation.Sort(); ok {
		_spec.SetField(winplatlist.FieldSort, field.TypeInt8, value)
		_node.Sort = value
	}
	if value, ok := wplc.mutation.Status(); ok {
		_spec.SetField(winplatlist.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := wplc.mutation.CreatedAt(); ok {
		_spec.SetField(winplatlist.FieldCreatedAt, field.TypeInt32, value)
		_node.CreatedAt = value
	}
	if value, ok := wplc.mutation.UpdatedAt(); ok {
		_spec.SetField(winplatlist.FieldUpdatedAt, field.TypeInt32, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WinPlatListCreateBulk is the builder for creating many WinPlatList entities in bulk.
type WinPlatListCreateBulk struct {
	config
	builders []*WinPlatListCreate
}

// Save creates the WinPlatList entities in the database.
func (wplcb *WinPlatListCreateBulk) Save(ctx context.Context) ([]*WinPlatList, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wplcb.builders))
	nodes := make([]*WinPlatList, len(wplcb.builders))
	mutators := make([]Mutator, len(wplcb.builders))
	for i := range wplcb.builders {
		func(i int, root context.Context) {
			builder := wplcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinPlatListMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wplcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wplcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wplcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wplcb *WinPlatListCreateBulk) SaveX(ctx context.Context) []*WinPlatList {
	v, err := wplcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wplcb *WinPlatListCreateBulk) Exec(ctx context.Context) error {
	_, err := wplcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wplcb *WinPlatListCreateBulk) ExecX(ctx context.Context) {
	if err := wplcb.Exec(ctx); err != nil {
		panic(err)
	}
}
