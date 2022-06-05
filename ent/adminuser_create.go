// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/avored/go-ecommerce/ent/adminuser"
)

// AdminUserCreate is the builder for creating a AdminUser entity.
type AdminUserCreate struct {
	config
	mutation *AdminUserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (auc *AdminUserCreate) SetCreatedAt(t time.Time) *AdminUserCreate {
	auc.mutation.SetCreatedAt(t)
	return auc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableCreatedAt(t *time.Time) *AdminUserCreate {
	if t != nil {
		auc.SetCreatedAt(*t)
	}
	return auc
}

// SetUpdatedAt sets the "updated_at" field.
func (auc *AdminUserCreate) SetUpdatedAt(t time.Time) *AdminUserCreate {
	auc.mutation.SetUpdatedAt(t)
	return auc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableUpdatedAt(t *time.Time) *AdminUserCreate {
	if t != nil {
		auc.SetUpdatedAt(*t)
	}
	return auc
}

// SetFirstName sets the "first_name" field.
func (auc *AdminUserCreate) SetFirstName(s string) *AdminUserCreate {
	auc.mutation.SetFirstName(s)
	return auc
}

// SetLastName sets the "last_name" field.
func (auc *AdminUserCreate) SetLastName(s string) *AdminUserCreate {
	auc.mutation.SetLastName(s)
	return auc
}

// SetEmail sets the "email" field.
func (auc *AdminUserCreate) SetEmail(s string) *AdminUserCreate {
	auc.mutation.SetEmail(s)
	return auc
}

// SetPassword sets the "password" field.
func (auc *AdminUserCreate) SetPassword(s string) *AdminUserCreate {
	auc.mutation.SetPassword(s)
	return auc
}

// Mutation returns the AdminUserMutation object of the builder.
func (auc *AdminUserCreate) Mutation() *AdminUserMutation {
	return auc.mutation
}

// Save creates the AdminUser in the database.
func (auc *AdminUserCreate) Save(ctx context.Context) (*AdminUser, error) {
	var (
		err  error
		node *AdminUser
	)
	auc.defaults()
	if len(auc.hooks) == 0 {
		if err = auc.check(); err != nil {
			return nil, err
		}
		node, err = auc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auc.check(); err != nil {
				return nil, err
			}
			auc.mutation = mutation
			if node, err = auc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(auc.hooks) - 1; i >= 0; i-- {
			if auc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auc *AdminUserCreate) SaveX(ctx context.Context) *AdminUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auc *AdminUserCreate) Exec(ctx context.Context) error {
	_, err := auc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auc *AdminUserCreate) ExecX(ctx context.Context) {
	if err := auc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auc *AdminUserCreate) defaults() {
	if _, ok := auc.mutation.CreatedAt(); !ok {
		v := adminuser.DefaultCreatedAt()
		auc.mutation.SetCreatedAt(v)
	}
	if _, ok := auc.mutation.UpdatedAt(); !ok {
		v := adminuser.DefaultUpdatedAt()
		auc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auc *AdminUserCreate) check() error {
	if _, ok := auc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AdminUser.created_at"`)}
	}
	if _, ok := auc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AdminUser.updated_at"`)}
	}
	if _, ok := auc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "AdminUser.first_name"`)}
	}
	if _, ok := auc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "AdminUser.last_name"`)}
	}
	if _, ok := auc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "AdminUser.email"`)}
	}
	if _, ok := auc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "AdminUser.password"`)}
	}
	return nil
}

func (auc *AdminUserCreate) sqlSave(ctx context.Context) (*AdminUser, error) {
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (auc *AdminUserCreate) createSpec() (*AdminUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AdminUser{config: auc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: adminuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adminuser.FieldID,
			},
		}
	)
	if value, ok := auc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adminuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := auc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adminuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := auc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adminuser.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := auc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adminuser.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := auc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adminuser.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := auc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adminuser.FieldPassword,
		})
		_node.Password = value
	}
	return _node, _spec
}

// AdminUserCreateBulk is the builder for creating many AdminUser entities in bulk.
type AdminUserCreateBulk struct {
	config
	builders []*AdminUserCreate
}

// Save creates the AdminUser entities in the database.
func (aucb *AdminUserCreateBulk) Save(ctx context.Context) ([]*AdminUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*AdminUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aucb *AdminUserCreateBulk) SaveX(ctx context.Context) []*AdminUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aucb *AdminUserCreateBulk) Exec(ctx context.Context) error {
	_, err := aucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucb *AdminUserCreateBulk) ExecX(ctx context.Context) {
	if err := aucb.Exec(ctx); err != nil {
		panic(err)
	}
}
