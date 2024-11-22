// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/inchori/starknet-indexer/ent/predicate"
	"github.com/inchori/starknet-indexer/ent/starknetdeclaretx"
)

// StarknetDeclareTxDelete is the builder for deleting a StarknetDeclareTx entity.
type StarknetDeclareTxDelete struct {
	config
	hooks    []Hook
	mutation *StarknetDeclareTxMutation
}

// Where appends a list predicates to the StarknetDeclareTxDelete builder.
func (sdtd *StarknetDeclareTxDelete) Where(ps ...predicate.StarknetDeclareTx) *StarknetDeclareTxDelete {
	sdtd.mutation.Where(ps...)
	return sdtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdtd *StarknetDeclareTxDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sdtd.sqlExec, sdtd.mutation, sdtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtd *StarknetDeclareTxDelete) ExecX(ctx context.Context) int {
	n, err := sdtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdtd *StarknetDeclareTxDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(starknetdeclaretx.Table, sqlgraph.NewFieldSpec(starknetdeclaretx.FieldID, field.TypeInt))
	if ps := sdtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sdtd.mutation.done = true
	return affected, err
}

// StarknetDeclareTxDeleteOne is the builder for deleting a single StarknetDeclareTx entity.
type StarknetDeclareTxDeleteOne struct {
	sdtd *StarknetDeclareTxDelete
}

// Where appends a list predicates to the StarknetDeclareTxDelete builder.
func (sdtdo *StarknetDeclareTxDeleteOne) Where(ps ...predicate.StarknetDeclareTx) *StarknetDeclareTxDeleteOne {
	sdtdo.sdtd.mutation.Where(ps...)
	return sdtdo
}

// Exec executes the deletion query.
func (sdtdo *StarknetDeclareTxDeleteOne) Exec(ctx context.Context) error {
	n, err := sdtdo.sdtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{starknetdeclaretx.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtdo *StarknetDeclareTxDeleteOne) ExecX(ctx context.Context) {
	if err := sdtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
