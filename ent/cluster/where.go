// Code generated by ent, DO NOT EDIT.

package cluster

import (
	"kubecit/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Cluster {
	return predicate.Cluster(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Cluster {
	return predicate.Cluster(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Cluster {
	return predicate.Cluster(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Cluster {
	return predicate.Cluster(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Cluster {
	return predicate.Cluster(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Cluster {
	return predicate.Cluster(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Cluster {
	return predicate.Cluster(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Cluster {
	return predicate.Cluster(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Cluster {
	return predicate.Cluster(sql.FieldLTE(FieldID, id))
}

// Kubeconfig applies equality check predicate on the "kubeconfig" field. It's identical to KubeconfigEQ.
func Kubeconfig(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldEQ(FieldKubeconfig, v))
}

// Alias applies equality check predicate on the "alias" field. It's identical to AliasEQ.
func Alias(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldEQ(FieldAlias, v))
}

// KubeconfigEQ applies the EQ predicate on the "kubeconfig" field.
func KubeconfigEQ(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldEQ(FieldKubeconfig, v))
}

// KubeconfigNEQ applies the NEQ predicate on the "kubeconfig" field.
func KubeconfigNEQ(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldNEQ(FieldKubeconfig, v))
}

// KubeconfigIn applies the In predicate on the "kubeconfig" field.
func KubeconfigIn(vs ...string) predicate.Cluster {
	return predicate.Cluster(sql.FieldIn(FieldKubeconfig, vs...))
}

// KubeconfigNotIn applies the NotIn predicate on the "kubeconfig" field.
func KubeconfigNotIn(vs ...string) predicate.Cluster {
	return predicate.Cluster(sql.FieldNotIn(FieldKubeconfig, vs...))
}

// KubeconfigGT applies the GT predicate on the "kubeconfig" field.
func KubeconfigGT(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldGT(FieldKubeconfig, v))
}

// KubeconfigGTE applies the GTE predicate on the "kubeconfig" field.
func KubeconfigGTE(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldGTE(FieldKubeconfig, v))
}

// KubeconfigLT applies the LT predicate on the "kubeconfig" field.
func KubeconfigLT(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldLT(FieldKubeconfig, v))
}

// KubeconfigLTE applies the LTE predicate on the "kubeconfig" field.
func KubeconfigLTE(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldLTE(FieldKubeconfig, v))
}

// KubeconfigContains applies the Contains predicate on the "kubeconfig" field.
func KubeconfigContains(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldContains(FieldKubeconfig, v))
}

// KubeconfigHasPrefix applies the HasPrefix predicate on the "kubeconfig" field.
func KubeconfigHasPrefix(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldHasPrefix(FieldKubeconfig, v))
}

// KubeconfigHasSuffix applies the HasSuffix predicate on the "kubeconfig" field.
func KubeconfigHasSuffix(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldHasSuffix(FieldKubeconfig, v))
}

// KubeconfigEqualFold applies the EqualFold predicate on the "kubeconfig" field.
func KubeconfigEqualFold(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldEqualFold(FieldKubeconfig, v))
}

// KubeconfigContainsFold applies the ContainsFold predicate on the "kubeconfig" field.
func KubeconfigContainsFold(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldContainsFold(FieldKubeconfig, v))
}

// AliasEQ applies the EQ predicate on the "alias" field.
func AliasEQ(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldEQ(FieldAlias, v))
}

// AliasNEQ applies the NEQ predicate on the "alias" field.
func AliasNEQ(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldNEQ(FieldAlias, v))
}

// AliasIn applies the In predicate on the "alias" field.
func AliasIn(vs ...string) predicate.Cluster {
	return predicate.Cluster(sql.FieldIn(FieldAlias, vs...))
}

// AliasNotIn applies the NotIn predicate on the "alias" field.
func AliasNotIn(vs ...string) predicate.Cluster {
	return predicate.Cluster(sql.FieldNotIn(FieldAlias, vs...))
}

// AliasGT applies the GT predicate on the "alias" field.
func AliasGT(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldGT(FieldAlias, v))
}

// AliasGTE applies the GTE predicate on the "alias" field.
func AliasGTE(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldGTE(FieldAlias, v))
}

// AliasLT applies the LT predicate on the "alias" field.
func AliasLT(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldLT(FieldAlias, v))
}

// AliasLTE applies the LTE predicate on the "alias" field.
func AliasLTE(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldLTE(FieldAlias, v))
}

// AliasContains applies the Contains predicate on the "alias" field.
func AliasContains(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldContains(FieldAlias, v))
}

// AliasHasPrefix applies the HasPrefix predicate on the "alias" field.
func AliasHasPrefix(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldHasPrefix(FieldAlias, v))
}

// AliasHasSuffix applies the HasSuffix predicate on the "alias" field.
func AliasHasSuffix(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldHasSuffix(FieldAlias, v))
}

// AliasEqualFold applies the EqualFold predicate on the "alias" field.
func AliasEqualFold(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldEqualFold(FieldAlias, v))
}

// AliasContainsFold applies the ContainsFold predicate on the "alias" field.
func AliasContainsFold(v string) predicate.Cluster {
	return predicate.Cluster(sql.FieldContainsFold(FieldAlias, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Cluster) predicate.Cluster {
	return predicate.Cluster(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Cluster) predicate.Cluster {
	return predicate.Cluster(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Cluster) predicate.Cluster {
	return predicate.Cluster(func(s *sql.Selector) {
		p(s.Not())
	})
}
