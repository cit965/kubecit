// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kubecit/ent/cloudhost"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CloudHost is the model entity for the CloudHost schema.
type CloudHost struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// InstanceId holds the value of the "instanceId" field.
	InstanceId string `json:"instanceId,omitempty"`
	// VpcId holds the value of the "vpcId" field.
	VpcId string `json:"vpcId,omitempty"`
	// SubnetId holds the value of the "subnetId" field.
	SubnetId string `json:"subnetId,omitempty"`
	// InstanceName holds the value of the "instanceName" field.
	InstanceName string `json:"instanceName,omitempty"`
	// InstanceState holds the value of the "instanceState" field.
	InstanceState string `json:"instanceState,omitempty"`
	// CPU holds the value of the "cpu" field.
	CPU int64 `json:"cpu,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory int64 `json:"memory,omitempty"`
	// CreatedTime holds the value of the "createdTime" field.
	CreatedTime string `json:"createdTime,omitempty"`
	// InstanceType holds the value of the "instanceType" field.
	InstanceType string `json:"instanceType,omitempty"`
	// EniLimit holds the value of the "eniLimit" field.
	EniLimit int64 `json:"eniLimit,omitempty"`
	// EnilpLimit holds the value of the "enilpLimit" field.
	EnilpLimit int64 `json:"enilpLimit,omitempty"`
	// InstanceEniCount holds the value of the "instanceEniCount" field.
	InstanceEniCount int64 `json:"instanceEniCount,omitempty"`
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CloudHost) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cloudhost.FieldID, cloudhost.FieldCPU, cloudhost.FieldMemory, cloudhost.FieldEniLimit, cloudhost.FieldEnilpLimit, cloudhost.FieldInstanceEniCount:
			values[i] = new(sql.NullInt64)
		case cloudhost.FieldInstanceId, cloudhost.FieldVpcId, cloudhost.FieldSubnetId, cloudhost.FieldInstanceName, cloudhost.FieldInstanceState, cloudhost.FieldCreatedTime, cloudhost.FieldInstanceType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CloudHost fields.
func (ch *CloudHost) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cloudhost.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ch.ID = int(value.Int64)
		case cloudhost.FieldInstanceId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instanceId", values[i])
			} else if value.Valid {
				ch.InstanceId = value.String
			}
		case cloudhost.FieldVpcId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vpcId", values[i])
			} else if value.Valid {
				ch.VpcId = value.String
			}
		case cloudhost.FieldSubnetId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subnetId", values[i])
			} else if value.Valid {
				ch.SubnetId = value.String
			}
		case cloudhost.FieldInstanceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instanceName", values[i])
			} else if value.Valid {
				ch.InstanceName = value.String
			}
		case cloudhost.FieldInstanceState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instanceState", values[i])
			} else if value.Valid {
				ch.InstanceState = value.String
			}
		case cloudhost.FieldCPU:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu", values[i])
			} else if value.Valid {
				ch.CPU = value.Int64
			}
		case cloudhost.FieldMemory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				ch.Memory = value.Int64
			}
		case cloudhost.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field createdTime", values[i])
			} else if value.Valid {
				ch.CreatedTime = value.String
			}
		case cloudhost.FieldInstanceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instanceType", values[i])
			} else if value.Valid {
				ch.InstanceType = value.String
			}
		case cloudhost.FieldEniLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field eniLimit", values[i])
			} else if value.Valid {
				ch.EniLimit = value.Int64
			}
		case cloudhost.FieldEnilpLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field enilpLimit", values[i])
			} else if value.Valid {
				ch.EnilpLimit = value.Int64
			}
		case cloudhost.FieldInstanceEniCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field instanceEniCount", values[i])
			} else if value.Valid {
				ch.InstanceEniCount = value.Int64
			}
		default:
			ch.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CloudHost.
// This includes values selected through modifiers, order, etc.
func (ch *CloudHost) Value(name string) (ent.Value, error) {
	return ch.selectValues.Get(name)
}

// Update returns a builder for updating this CloudHost.
// Note that you need to call CloudHost.Unwrap() before calling this method if this CloudHost
// was returned from a transaction, and the transaction was committed or rolled back.
func (ch *CloudHost) Update() *CloudHostUpdateOne {
	return NewCloudHostClient(ch.config).UpdateOne(ch)
}

// Unwrap unwraps the CloudHost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ch *CloudHost) Unwrap() *CloudHost {
	_tx, ok := ch.config.driver.(*txDriver)
	if !ok {
		panic("ent: CloudHost is not a transactional entity")
	}
	ch.config.driver = _tx.drv
	return ch
}

// String implements the fmt.Stringer.
func (ch *CloudHost) String() string {
	var builder strings.Builder
	builder.WriteString("CloudHost(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ch.ID))
	builder.WriteString("instanceId=")
	builder.WriteString(ch.InstanceId)
	builder.WriteString(", ")
	builder.WriteString("vpcId=")
	builder.WriteString(ch.VpcId)
	builder.WriteString(", ")
	builder.WriteString("subnetId=")
	builder.WriteString(ch.SubnetId)
	builder.WriteString(", ")
	builder.WriteString("instanceName=")
	builder.WriteString(ch.InstanceName)
	builder.WriteString(", ")
	builder.WriteString("instanceState=")
	builder.WriteString(ch.InstanceState)
	builder.WriteString(", ")
	builder.WriteString("cpu=")
	builder.WriteString(fmt.Sprintf("%v", ch.CPU))
	builder.WriteString(", ")
	builder.WriteString("memory=")
	builder.WriteString(fmt.Sprintf("%v", ch.Memory))
	builder.WriteString(", ")
	builder.WriteString("createdTime=")
	builder.WriteString(ch.CreatedTime)
	builder.WriteString(", ")
	builder.WriteString("instanceType=")
	builder.WriteString(ch.InstanceType)
	builder.WriteString(", ")
	builder.WriteString("eniLimit=")
	builder.WriteString(fmt.Sprintf("%v", ch.EniLimit))
	builder.WriteString(", ")
	builder.WriteString("enilpLimit=")
	builder.WriteString(fmt.Sprintf("%v", ch.EnilpLimit))
	builder.WriteString(", ")
	builder.WriteString("instanceEniCount=")
	builder.WriteString(fmt.Sprintf("%v", ch.InstanceEniCount))
	builder.WriteByte(')')
	return builder.String()
}

// CloudHosts is a parsable slice of CloudHost.
type CloudHosts []*CloudHost