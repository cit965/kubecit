package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CloudHost holds the schema definition for the CloudHost entity.
type CloudHost struct {
	ent.Schema
}

// Fields of the CloudHost.
func (CloudHost) Fields() []ent.Field {
	return []ent.Field{
		field.String("instanceId").NotEmpty().Unique(),
		field.String("vpcId").NotEmpty(),
		field.String("subnetId"),
		field.String("instanceName"),
		field.String("instanceState"),
		field.Int64("cpu"),
		field.Int64("memory"),
		field.String("createdTime"),
		field.String("instanceType"),
		field.Int64("eniLimit"),
		field.Int64("enilpLimit"),
		field.Int64("instanceEniCount"),
	}
}

// Edges of the CloudHost.
func (CloudHost) Edges() []ent.Edge {
	return nil
}
