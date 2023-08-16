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
		field.String("uuid").NotEmpty().Unique(),
		field.String("state").NotEmpty().Default("RUNNING"),
		field.String("ipv6AddressPrivate"),
		field.String("ipv6AddressPublic"),
		field.String("ipv4AddressPrivate"),
		field.String("ipv4AddressPublic"),
		field.Int("memory"),
		field.Int("cpu"),
		field.Time("createdTime"),
		field.Time("expiredTime"),
		field.String("instanceName"),
		field.String("imageName"),
		field.String("osType"),
		field.String("manufacturer"),
		field.String("zone"),
		field.String("securityGroups"),
		field.String("billType"),
		field.String("chargeType"),
		field.Bool("isActive").Default(true),
		field.String("instanceType"),
	}
}

// Edges of the CloudHost.
func (CloudHost) Edges() []ent.Edge {
	return nil
}
