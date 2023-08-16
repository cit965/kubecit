// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kubecit/ent/cluster"
	"kubecit/ent/schema"
	"kubecit/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	clusterFields := schema.Cluster{}.Fields()
	_ = clusterFields
	// clusterDescKubeconfig is the schema descriptor for kubeconfig field.
	clusterDescKubeconfig := clusterFields[0].Descriptor()
	// cluster.DefaultKubeconfig holds the default value on creation for the kubeconfig field.
	cluster.DefaultKubeconfig = clusterDescKubeconfig.Default.(string)
	// clusterDescAlias is the schema descriptor for alias field.
	clusterDescAlias := clusterFields[1].Descriptor()
	// cluster.DefaultAlias holds the default value on creation for the alias field.
	cluster.DefaultAlias = clusterDescAlias.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
}
