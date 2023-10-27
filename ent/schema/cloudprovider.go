package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CloudProvider holds the schema definition for the CloudProvider entity.
type CloudProvider struct {
	ent.Schema
}

// Fields of the CloudProvider.
func (CloudProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").
			Default("unknown"),
		field.String("secret"),
		field.Int64("type").Comment("1是腾讯云2是华为云"),
		field.String("name"),
	}
}

// Edges of the CloudProvider.
func (CloudProvider) Edges() []ent.Edge {
	return nil
}
