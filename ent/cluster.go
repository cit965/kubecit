// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kubecit/ent/cluster"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Cluster is the model entity for the Cluster schema.
type Cluster struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Kubeconfig holds the value of the "kubeconfig" field.
	Kubeconfig string `json:"kubeconfig,omitempty"`
	// Alias holds the value of the "alias" field.
	Alias        string `json:"alias,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cluster) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cluster.FieldID:
			values[i] = new(sql.NullInt64)
		case cluster.FieldKubeconfig, cluster.FieldAlias:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cluster fields.
func (c *Cluster) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cluster.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cluster.FieldKubeconfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kubeconfig", values[i])
			} else if value.Valid {
				c.Kubeconfig = value.String
			}
		case cluster.FieldAlias:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alias", values[i])
			} else if value.Valid {
				c.Alias = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Cluster.
// This includes values selected through modifiers, order, etc.
func (c *Cluster) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Cluster.
// Note that you need to call Cluster.Unwrap() before calling this method if this Cluster
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cluster) Update() *ClusterUpdateOne {
	return NewClusterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Cluster entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cluster) Unwrap() *Cluster {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cluster is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cluster) String() string {
	var builder strings.Builder
	builder.WriteString("Cluster(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("kubeconfig=")
	builder.WriteString(c.Kubeconfig)
	builder.WriteString(", ")
	builder.WriteString("alias=")
	builder.WriteString(c.Alias)
	builder.WriteByte(')')
	return builder.String()
}

// Clusters is a parsable slice of Cluster.
type Clusters []*Cluster
