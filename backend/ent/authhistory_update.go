// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Minto312/passkey-practice/backend/ent/authhistory"
	"github.com/Minto312/passkey-practice/backend/ent/predicate"
	"github.com/Minto312/passkey-practice/backend/ent/user"
	"github.com/google/uuid"
)

// AuthHistoryUpdate is the builder for updating AuthHistory entities.
type AuthHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *AuthHistoryMutation
}

// Where appends a list predicates to the AuthHistoryUpdate builder.
func (ahu *AuthHistoryUpdate) Where(ps ...predicate.AuthHistory) *AuthHistoryUpdate {
	ahu.mutation.Where(ps...)
	return ahu
}

// SetMethod sets the "method" field.
func (ahu *AuthHistoryUpdate) SetMethod(s string) *AuthHistoryUpdate {
	ahu.mutation.SetMethod(s)
	return ahu
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ahu *AuthHistoryUpdate) SetNillableMethod(s *string) *AuthHistoryUpdate {
	if s != nil {
		ahu.SetMethod(*s)
	}
	return ahu
}

// SetIPAddress sets the "ip_address" field.
func (ahu *AuthHistoryUpdate) SetIPAddress(s string) *AuthHistoryUpdate {
	ahu.mutation.SetIPAddress(s)
	return ahu
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (ahu *AuthHistoryUpdate) SetNillableIPAddress(s *string) *AuthHistoryUpdate {
	if s != nil {
		ahu.SetIPAddress(*s)
	}
	return ahu
}

// SetUserAgent sets the "user_agent" field.
func (ahu *AuthHistoryUpdate) SetUserAgent(s string) *AuthHistoryUpdate {
	ahu.mutation.SetUserAgent(s)
	return ahu
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (ahu *AuthHistoryUpdate) SetNillableUserAgent(s *string) *AuthHistoryUpdate {
	if s != nil {
		ahu.SetUserAgent(*s)
	}
	return ahu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ahu *AuthHistoryUpdate) SetUserID(id uuid.UUID) *AuthHistoryUpdate {
	ahu.mutation.SetUserID(id)
	return ahu
}

// SetUser sets the "user" edge to the User entity.
func (ahu *AuthHistoryUpdate) SetUser(u *User) *AuthHistoryUpdate {
	return ahu.SetUserID(u.ID)
}

// Mutation returns the AuthHistoryMutation object of the builder.
func (ahu *AuthHistoryUpdate) Mutation() *AuthHistoryMutation {
	return ahu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ahu *AuthHistoryUpdate) ClearUser() *AuthHistoryUpdate {
	ahu.mutation.ClearUser()
	return ahu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ahu *AuthHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ahu.sqlSave, ahu.mutation, ahu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ahu *AuthHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ahu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ahu *AuthHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ahu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahu *AuthHistoryUpdate) ExecX(ctx context.Context) {
	if err := ahu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ahu *AuthHistoryUpdate) check() error {
	if ahu.mutation.UserCleared() && len(ahu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "AuthHistory.user"`)
	}
	return nil
}

func (ahu *AuthHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ahu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(authhistory.Table, authhistory.Columns, sqlgraph.NewFieldSpec(authhistory.FieldID, field.TypeUUID))
	if ps := ahu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ahu.mutation.Method(); ok {
		_spec.SetField(authhistory.FieldMethod, field.TypeString, value)
	}
	if value, ok := ahu.mutation.IPAddress(); ok {
		_spec.SetField(authhistory.FieldIPAddress, field.TypeString, value)
	}
	if value, ok := ahu.mutation.UserAgent(); ok {
		_spec.SetField(authhistory.FieldUserAgent, field.TypeString, value)
	}
	if ahu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authhistory.UserTable,
			Columns: []string{authhistory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ahu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authhistory.UserTable,
			Columns: []string{authhistory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ahu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ahu.mutation.done = true
	return n, nil
}

// AuthHistoryUpdateOne is the builder for updating a single AuthHistory entity.
type AuthHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthHistoryMutation
}

// SetMethod sets the "method" field.
func (ahuo *AuthHistoryUpdateOne) SetMethod(s string) *AuthHistoryUpdateOne {
	ahuo.mutation.SetMethod(s)
	return ahuo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ahuo *AuthHistoryUpdateOne) SetNillableMethod(s *string) *AuthHistoryUpdateOne {
	if s != nil {
		ahuo.SetMethod(*s)
	}
	return ahuo
}

// SetIPAddress sets the "ip_address" field.
func (ahuo *AuthHistoryUpdateOne) SetIPAddress(s string) *AuthHistoryUpdateOne {
	ahuo.mutation.SetIPAddress(s)
	return ahuo
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (ahuo *AuthHistoryUpdateOne) SetNillableIPAddress(s *string) *AuthHistoryUpdateOne {
	if s != nil {
		ahuo.SetIPAddress(*s)
	}
	return ahuo
}

// SetUserAgent sets the "user_agent" field.
func (ahuo *AuthHistoryUpdateOne) SetUserAgent(s string) *AuthHistoryUpdateOne {
	ahuo.mutation.SetUserAgent(s)
	return ahuo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (ahuo *AuthHistoryUpdateOne) SetNillableUserAgent(s *string) *AuthHistoryUpdateOne {
	if s != nil {
		ahuo.SetUserAgent(*s)
	}
	return ahuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ahuo *AuthHistoryUpdateOne) SetUserID(id uuid.UUID) *AuthHistoryUpdateOne {
	ahuo.mutation.SetUserID(id)
	return ahuo
}

// SetUser sets the "user" edge to the User entity.
func (ahuo *AuthHistoryUpdateOne) SetUser(u *User) *AuthHistoryUpdateOne {
	return ahuo.SetUserID(u.ID)
}

// Mutation returns the AuthHistoryMutation object of the builder.
func (ahuo *AuthHistoryUpdateOne) Mutation() *AuthHistoryMutation {
	return ahuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ahuo *AuthHistoryUpdateOne) ClearUser() *AuthHistoryUpdateOne {
	ahuo.mutation.ClearUser()
	return ahuo
}

// Where appends a list predicates to the AuthHistoryUpdate builder.
func (ahuo *AuthHistoryUpdateOne) Where(ps ...predicate.AuthHistory) *AuthHistoryUpdateOne {
	ahuo.mutation.Where(ps...)
	return ahuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ahuo *AuthHistoryUpdateOne) Select(field string, fields ...string) *AuthHistoryUpdateOne {
	ahuo.fields = append([]string{field}, fields...)
	return ahuo
}

// Save executes the query and returns the updated AuthHistory entity.
func (ahuo *AuthHistoryUpdateOne) Save(ctx context.Context) (*AuthHistory, error) {
	return withHooks(ctx, ahuo.sqlSave, ahuo.mutation, ahuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ahuo *AuthHistoryUpdateOne) SaveX(ctx context.Context) *AuthHistory {
	node, err := ahuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ahuo *AuthHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ahuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahuo *AuthHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ahuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ahuo *AuthHistoryUpdateOne) check() error {
	if ahuo.mutation.UserCleared() && len(ahuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "AuthHistory.user"`)
	}
	return nil
}

func (ahuo *AuthHistoryUpdateOne) sqlSave(ctx context.Context) (_node *AuthHistory, err error) {
	if err := ahuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(authhistory.Table, authhistory.Columns, sqlgraph.NewFieldSpec(authhistory.FieldID, field.TypeUUID))
	id, ok := ahuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ahuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authhistory.FieldID)
		for _, f := range fields {
			if !authhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ahuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ahuo.mutation.Method(); ok {
		_spec.SetField(authhistory.FieldMethod, field.TypeString, value)
	}
	if value, ok := ahuo.mutation.IPAddress(); ok {
		_spec.SetField(authhistory.FieldIPAddress, field.TypeString, value)
	}
	if value, ok := ahuo.mutation.UserAgent(); ok {
		_spec.SetField(authhistory.FieldUserAgent, field.TypeString, value)
	}
	if ahuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authhistory.UserTable,
			Columns: []string{authhistory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ahuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authhistory.UserTable,
			Columns: []string{authhistory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AuthHistory{config: ahuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ahuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ahuo.mutation.done = true
	return _node, nil
}
