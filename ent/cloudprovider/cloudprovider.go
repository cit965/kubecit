// Code generated by ent, DO NOT EDIT.

package cloudprovider

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cloudprovider type in the database.
	Label = "cloud_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the cloudprovider in the database.
	Table = "cloud_providers"
)

// Columns holds all SQL columns for cloudprovider fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldSecret,
	FieldType,
	FieldName,
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
	// DefaultKey holds the default value on creation for the "key" field.
	DefaultKey string
)

// OrderOption defines the ordering options for the CloudProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// BySecret orders the results by the secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}
