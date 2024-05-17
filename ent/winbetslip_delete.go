// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/terra-v99/common/ent/predicate"
	"github.com/terra-v99/common/ent/winbetslip"
)

// WinBetslipDelete is the builder for deleting a WinBetslip entity.
type WinBetslipDelete struct {
	config
	hooks    []Hook
	mutation *WinBetslipMutation
}

// Where appends a list predicates to the WinBetslipDelete builder.
func (wbd *WinBetslipDelete) Where(ps ...predicate.WinBetslip) *WinBetslipDelete {
	wbd.mutation.Where(ps...)
	return wbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wbd *WinBetslipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wbd.sqlExec, wbd.mutation, wbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wbd *WinBetslipDelete) ExecX(ctx context.Context) int {
	n, err := wbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wbd *WinBetslipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(winbetslip.Table, sqlgraph.NewFieldSpec(winbetslip.FieldID, field.TypeInt))
	if ps := wbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wbd.mutation.done = true
	return affected, err
}

// WinBetslipDeleteOne is the builder for deleting a single WinBetslip entity.
type WinBetslipDeleteOne struct {
	wbd *WinBetslipDelete
}

// Where appends a list predicates to the WinBetslipDelete builder.
func (wbdo *WinBetslipDeleteOne) Where(ps ...predicate.WinBetslip) *WinBetslipDeleteOne {
	wbdo.wbd.mutation.Where(ps...)
	return wbdo
}

// Exec executes the deletion query.
func (wbdo *WinBetslipDeleteOne) Exec(ctx context.Context) error {
	n, err := wbdo.wbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{winbetslip.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wbdo *WinBetslipDeleteOne) ExecX(ctx context.Context) {
	if err := wbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
