// Code generated by ent, DO NOT EDIT.

package cloudhost

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cloudhost type in the database.
	Label = "cloud_host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInstanceId holds the string denoting the instanceid field in the database.
	FieldInstanceId = "instance_id"
	// FieldVpcId holds the string denoting the vpcid field in the database.
	FieldVpcId = "vpc_id"
	// FieldSubnetId holds the string denoting the subnetid field in the database.
	FieldSubnetId = "subnet_id"
	// FieldInstanceName holds the string denoting the instancename field in the database.
	FieldInstanceName = "instance_name"
	// FieldInstanceState holds the string denoting the instancestate field in the database.
	FieldInstanceState = "instance_state"
	// FieldCPU holds the string denoting the cpu field in the database.
	FieldCPU = "cpu"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldCreatedTime holds the string denoting the createdtime field in the database.
	FieldCreatedTime = "created_time"
	// FieldInstanceType holds the string denoting the instancetype field in the database.
	FieldInstanceType = "instance_type"
	// FieldEniLimit holds the string denoting the enilimit field in the database.
	FieldEniLimit = "eni_limit"
	// FieldEnilpLimit holds the string denoting the enilplimit field in the database.
	FieldEnilpLimit = "enilp_limit"
	// FieldInstanceEniCount holds the string denoting the instanceenicount field in the database.
	FieldInstanceEniCount = "instance_eni_count"
	// Table holds the table name of the cloudhost in the database.
	Table = "cloud_hosts"
)

// Columns holds all SQL columns for cloudhost fields.
var Columns = []string{
	FieldID,
	FieldInstanceId,
	FieldVpcId,
	FieldSubnetId,
	FieldInstanceName,
	FieldInstanceState,
	FieldCPU,
	FieldMemory,
	FieldCreatedTime,
	FieldInstanceType,
	FieldEniLimit,
	FieldEnilpLimit,
	FieldInstanceEniCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// InstanceIdValidator is a validator for the "instanceId" field. It is called by the builders before save.
	InstanceIdValidator func(string) error
	// VpcIdValidator is a validator for the "vpcId" field. It is called by the builders before save.
	VpcIdValidator func(string) error
)

// OrderOption defines the ordering options for the CloudHost queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByInstanceId orders the results by the instanceId field.
func ByInstanceId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceId, opts...).ToFunc()
}

// ByVpcId orders the results by the vpcId field.
func ByVpcId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVpcId, opts...).ToFunc()
}

// BySubnetId orders the results by the subnetId field.
func BySubnetId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubnetId, opts...).ToFunc()
}

// ByInstanceName orders the results by the instanceName field.
func ByInstanceName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceName, opts...).ToFunc()
}

// ByInstanceState orders the results by the instanceState field.
func ByInstanceState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceState, opts...).ToFunc()
}

// ByCPU orders the results by the cpu field.
func ByCPU(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCPU, opts...).ToFunc()
}

// ByMemory orders the results by the memory field.
func ByMemory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemory, opts...).ToFunc()
}

// ByCreatedTime orders the results by the createdTime field.
func ByCreatedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedTime, opts...).ToFunc()
}

// ByInstanceType orders the results by the instanceType field.
func ByInstanceType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceType, opts...).ToFunc()
}

// ByEniLimit orders the results by the eniLimit field.
func ByEniLimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEniLimit, opts...).ToFunc()
}

// ByEnilpLimit orders the results by the enilpLimit field.
func ByEnilpLimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnilpLimit, opts...).ToFunc()
}

// ByInstanceEniCount orders the results by the instanceEniCount field.
func ByInstanceEniCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceEniCount, opts...).ToFunc()
}
