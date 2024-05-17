// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/terra-v99/common/ent/windict"
)

// WinDictCreate is the builder for creating a WinDict entity.
type WinDictCreate struct {
	config
	mutation *WinDictMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (wdc *WinDictCreate) SetTitle(s string) *WinDictCreate {
	wdc.mutation.SetTitle(s)
	return wdc
}

// SetCategory sets the "category" field.
func (wdc *WinDictCreate) SetCategory(s string) *WinDictCreate {
	wdc.mutation.SetCategory(s)
	return wdc
}

// SetStatus sets the "status" field.
func (wdc *WinDictCreate) SetStatus(b bool) *WinDictCreate {
	wdc.mutation.SetStatus(b)
	return wdc
}

// SetCreatedAt sets the "created_at" field.
func (wdc *WinDictCreate) SetCreatedAt(i int32) *WinDictCreate {
	wdc.mutation.SetCreatedAt(i)
	return wdc
}

// SetUpdatedAt sets the "updated_at" field.
func (wdc *WinDictCreate) SetUpdatedAt(i int32) *WinDictCreate {
	wdc.mutation.SetUpdatedAt(i)
	return wdc
}

// SetID sets the "id" field.
func (wdc *WinDictCreate) SetID(i int32) *WinDictCreate {
	wdc.mutation.SetID(i)
	return wdc
}

// Mutation returns the WinDictMutation object of the builder.
func (wdc *WinDictCreate) Mutation() *WinDictMutation {
	return wdc.mutation
}

// Save creates the WinDict in the database.
func (wdc *WinDictCreate) Save(ctx context.Context) (*WinDict, error) {
	return withHooks(ctx, wdc.sqlSave, wdc.mutation, wdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wdc *WinDictCreate) SaveX(ctx context.Context) *WinDict {
	v, err := wdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdc *WinDictCreate) Exec(ctx context.Context) error {
	_, err := wdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdc *WinDictCreate) ExecX(ctx context.Context) {
	if err := wdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wdc *WinDictCreate) check() error {
	if _, ok := wdc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "WinDict.title"`)}
	}
	if _, ok := wdc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "WinDict.category"`)}
	}
	if _, ok := wdc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "WinDict.status"`)}
	}
	if _, ok := wdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WinDict.created_at"`)}
	}
	if _, ok := wdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WinDict.updated_at"`)}
	}
	return nil
}

func (wdc *WinDictCreate) sqlSave(ctx context.Context) (*WinDict, error) {
	if err := wdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	wdc.mutation.id = &_node.ID
	wdc.mutation.done = true
	return _node, nil
}

func (wdc *WinDictCreate) createSpec() (*WinDict, *sqlgraph.CreateSpec) {
	var (
		_node = &WinDict{config: wdc.config}
		_spec = sqlgraph.NewCreateSpec(windict.Table, sqlgraph.NewFieldSpec(windict.FieldID, field.TypeInt32))
	)
	if id, ok := wdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wdc.mutation.Title(); ok {
		_spec.SetField(windict.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := wdc.mutation.Category(); ok {
		_spec.SetField(windict.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := wdc.mutation.Status(); ok {
		_spec.SetField(windict.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if value, ok := wdc.mutation.CreatedAt(); ok {
		_spec.SetField(windict.FieldCreatedAt, field.TypeInt32, value)
		_node.CreatedAt = value
	}
	if value, ok := wdc.mutation.UpdatedAt(); ok {
		_spec.SetField(windict.FieldUpdatedAt, field.TypeInt32, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WinDictCreateBulk is the builder for creating many WinDict entities in bulk.
type WinDictCreateBulk struct {
	config
	builders []*WinDictCreate
}

// Save creates the WinDict entities in the database.
func (wdcb *WinDictCreateBulk) Save(ctx context.Context) ([]*WinDict, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wdcb.builders))
	nodes := make([]*WinDict, len(wdcb.builders))
	mutators := make([]Mutator, len(wdcb.builders))
	for i := range wdcb.builders {
		func(i int, root context.Context) {
			builder := wdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinDictMutation)
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
					_, err = mutators[i+1].Mutate(root, wdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wdcb *WinDictCreateBulk) SaveX(ctx context.Context) []*WinDict {
	v, err := wdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdcb *WinDictCreateBulk) Exec(ctx context.Context) error {
	_, err := wdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdcb *WinDictCreateBulk) ExecX(ctx context.Context) {
	if err := wdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
