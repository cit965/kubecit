// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit/ent/cloudhost"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CloudHostCreate is the builder for creating a CloudHost entity.
type CloudHostCreate struct {
	config
	mutation *CloudHostMutation
	hooks    []Hook
}

// SetInstanceId sets the "instanceId" field.
func (chc *CloudHostCreate) SetInstanceId(s string) *CloudHostCreate {
	chc.mutation.SetInstanceId(s)
	return chc
}

// SetVpcId sets the "vpcId" field.
func (chc *CloudHostCreate) SetVpcId(s string) *CloudHostCreate {
	chc.mutation.SetVpcId(s)
	return chc
}

// SetSubnetId sets the "subnetId" field.
func (chc *CloudHostCreate) SetSubnetId(s string) *CloudHostCreate {
	chc.mutation.SetSubnetId(s)
	return chc
}

// SetInstanceName sets the "instanceName" field.
func (chc *CloudHostCreate) SetInstanceName(s string) *CloudHostCreate {
	chc.mutation.SetInstanceName(s)
	return chc
}

// SetInstanceState sets the "instanceState" field.
func (chc *CloudHostCreate) SetInstanceState(s string) *CloudHostCreate {
	chc.mutation.SetInstanceState(s)
	return chc
}

// SetCPU sets the "cpu" field.
func (chc *CloudHostCreate) SetCPU(i int64) *CloudHostCreate {
	chc.mutation.SetCPU(i)
	return chc
}

// SetMemory sets the "memory" field.
func (chc *CloudHostCreate) SetMemory(i int64) *CloudHostCreate {
	chc.mutation.SetMemory(i)
	return chc
}

// SetCreatedTime sets the "createdTime" field.
func (chc *CloudHostCreate) SetCreatedTime(s string) *CloudHostCreate {
	chc.mutation.SetCreatedTime(s)
	return chc
}

// SetInstanceType sets the "instanceType" field.
func (chc *CloudHostCreate) SetInstanceType(s string) *CloudHostCreate {
	chc.mutation.SetInstanceType(s)
	return chc
}

// SetEniLimit sets the "eniLimit" field.
func (chc *CloudHostCreate) SetEniLimit(i int64) *CloudHostCreate {
	chc.mutation.SetEniLimit(i)
	return chc
}

// SetEnilpLimit sets the "enilpLimit" field.
func (chc *CloudHostCreate) SetEnilpLimit(i int64) *CloudHostCreate {
	chc.mutation.SetEnilpLimit(i)
	return chc
}

// SetInstanceEniCount sets the "instanceEniCount" field.
func (chc *CloudHostCreate) SetInstanceEniCount(i int64) *CloudHostCreate {
	chc.mutation.SetInstanceEniCount(i)
	return chc
}

// Mutation returns the CloudHostMutation object of the builder.
func (chc *CloudHostCreate) Mutation() *CloudHostMutation {
	return chc.mutation
}

// Save creates the CloudHost in the database.
func (chc *CloudHostCreate) Save(ctx context.Context) (*CloudHost, error) {
	return withHooks(ctx, chc.sqlSave, chc.mutation, chc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (chc *CloudHostCreate) SaveX(ctx context.Context) *CloudHost {
	v, err := chc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chc *CloudHostCreate) Exec(ctx context.Context) error {
	_, err := chc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chc *CloudHostCreate) ExecX(ctx context.Context) {
	if err := chc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chc *CloudHostCreate) check() error {
	if _, ok := chc.mutation.InstanceId(); !ok {
		return &ValidationError{Name: "instanceId", err: errors.New(`ent: missing required field "CloudHost.instanceId"`)}
	}
	if v, ok := chc.mutation.InstanceId(); ok {
		if err := cloudhost.InstanceIdValidator(v); err != nil {
			return &ValidationError{Name: "instanceId", err: fmt.Errorf(`ent: validator failed for field "CloudHost.instanceId": %w`, err)}
		}
	}
	if _, ok := chc.mutation.VpcId(); !ok {
		return &ValidationError{Name: "vpcId", err: errors.New(`ent: missing required field "CloudHost.vpcId"`)}
	}
	if v, ok := chc.mutation.VpcId(); ok {
		if err := cloudhost.VpcIdValidator(v); err != nil {
			return &ValidationError{Name: "vpcId", err: fmt.Errorf(`ent: validator failed for field "CloudHost.vpcId": %w`, err)}
		}
	}
	if _, ok := chc.mutation.SubnetId(); !ok {
		return &ValidationError{Name: "subnetId", err: errors.New(`ent: missing required field "CloudHost.subnetId"`)}
	}
	if _, ok := chc.mutation.InstanceName(); !ok {
		return &ValidationError{Name: "instanceName", err: errors.New(`ent: missing required field "CloudHost.instanceName"`)}
	}
	if _, ok := chc.mutation.InstanceState(); !ok {
		return &ValidationError{Name: "instanceState", err: errors.New(`ent: missing required field "CloudHost.instanceState"`)}
	}
	if _, ok := chc.mutation.CPU(); !ok {
		return &ValidationError{Name: "cpu", err: errors.New(`ent: missing required field "CloudHost.cpu"`)}
	}
	if _, ok := chc.mutation.Memory(); !ok {
		return &ValidationError{Name: "memory", err: errors.New(`ent: missing required field "CloudHost.memory"`)}
	}
	if _, ok := chc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "createdTime", err: errors.New(`ent: missing required field "CloudHost.createdTime"`)}
	}
	if _, ok := chc.mutation.InstanceType(); !ok {
		return &ValidationError{Name: "instanceType", err: errors.New(`ent: missing required field "CloudHost.instanceType"`)}
	}
	if _, ok := chc.mutation.EniLimit(); !ok {
		return &ValidationError{Name: "eniLimit", err: errors.New(`ent: missing required field "CloudHost.eniLimit"`)}
	}
	if _, ok := chc.mutation.EnilpLimit(); !ok {
		return &ValidationError{Name: "enilpLimit", err: errors.New(`ent: missing required field "CloudHost.enilpLimit"`)}
	}
	if _, ok := chc.mutation.InstanceEniCount(); !ok {
		return &ValidationError{Name: "instanceEniCount", err: errors.New(`ent: missing required field "CloudHost.instanceEniCount"`)}
	}
	return nil
}

func (chc *CloudHostCreate) sqlSave(ctx context.Context) (*CloudHost, error) {
	if err := chc.check(); err != nil {
		return nil, err
	}
	_node, _spec := chc.createSpec()
	if err := sqlgraph.CreateNode(ctx, chc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	chc.mutation.id = &_node.ID
	chc.mutation.done = true
	return _node, nil
}

func (chc *CloudHostCreate) createSpec() (*CloudHost, *sqlgraph.CreateSpec) {
	var (
		_node = &CloudHost{config: chc.config}
		_spec = sqlgraph.NewCreateSpec(cloudhost.Table, sqlgraph.NewFieldSpec(cloudhost.FieldID, field.TypeInt))
	)
	if value, ok := chc.mutation.InstanceId(); ok {
		_spec.SetField(cloudhost.FieldInstanceId, field.TypeString, value)
		_node.InstanceId = value
	}
	if value, ok := chc.mutation.VpcId(); ok {
		_spec.SetField(cloudhost.FieldVpcId, field.TypeString, value)
		_node.VpcId = value
	}
	if value, ok := chc.mutation.SubnetId(); ok {
		_spec.SetField(cloudhost.FieldSubnetId, field.TypeString, value)
		_node.SubnetId = value
	}
	if value, ok := chc.mutation.InstanceName(); ok {
		_spec.SetField(cloudhost.FieldInstanceName, field.TypeString, value)
		_node.InstanceName = value
	}
	if value, ok := chc.mutation.InstanceState(); ok {
		_spec.SetField(cloudhost.FieldInstanceState, field.TypeString, value)
		_node.InstanceState = value
	}
	if value, ok := chc.mutation.CPU(); ok {
		_spec.SetField(cloudhost.FieldCPU, field.TypeInt64, value)
		_node.CPU = value
	}
	if value, ok := chc.mutation.Memory(); ok {
		_spec.SetField(cloudhost.FieldMemory, field.TypeInt64, value)
		_node.Memory = value
	}
	if value, ok := chc.mutation.CreatedTime(); ok {
		_spec.SetField(cloudhost.FieldCreatedTime, field.TypeString, value)
		_node.CreatedTime = value
	}
	if value, ok := chc.mutation.InstanceType(); ok {
		_spec.SetField(cloudhost.FieldInstanceType, field.TypeString, value)
		_node.InstanceType = value
	}
	if value, ok := chc.mutation.EniLimit(); ok {
		_spec.SetField(cloudhost.FieldEniLimit, field.TypeInt64, value)
		_node.EniLimit = value
	}
	if value, ok := chc.mutation.EnilpLimit(); ok {
		_spec.SetField(cloudhost.FieldEnilpLimit, field.TypeInt64, value)
		_node.EnilpLimit = value
	}
	if value, ok := chc.mutation.InstanceEniCount(); ok {
		_spec.SetField(cloudhost.FieldInstanceEniCount, field.TypeInt64, value)
		_node.InstanceEniCount = value
	}
	return _node, _spec
}

// CloudHostCreateBulk is the builder for creating many CloudHost entities in bulk.
type CloudHostCreateBulk struct {
	config
	builders []*CloudHostCreate
}

// Save creates the CloudHost entities in the database.
func (chcb *CloudHostCreateBulk) Save(ctx context.Context) ([]*CloudHost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(chcb.builders))
	nodes := make([]*CloudHost, len(chcb.builders))
	mutators := make([]Mutator, len(chcb.builders))
	for i := range chcb.builders {
		func(i int, root context.Context) {
			builder := chcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CloudHostMutation)
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
					_, err = mutators[i+1].Mutate(root, chcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, chcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, chcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (chcb *CloudHostCreateBulk) SaveX(ctx context.Context) []*CloudHost {
	v, err := chcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chcb *CloudHostCreateBulk) Exec(ctx context.Context) error {
	_, err := chcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chcb *CloudHostCreateBulk) ExecX(ctx context.Context) {
	if err := chcb.Exec(ctx); err != nil {
		panic(err)
	}
}
