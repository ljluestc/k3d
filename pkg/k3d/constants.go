// Package k3d provides core types and utilities for k3d.
package k3d

// DefaultObjectNamePrefix is the prefix used for all k3d objects
const DefaultObjectNamePrefix = "k3d"

// DefaultToolsImageRepo is the default image repository used for the tools container
const DefaultToolsImageRepo = "ghcr.io/k3d-io/k3d-tools"

// Labels used to identify k3d objects
const (
	// LabelClusterName is the label used to identify which cluster a resource belongs to
	LabelClusterName = "k3d.cluster.name"
	// LabelRole is the label used to identify the role of a node
	LabelRole = "k3d.role"
	// LabelClusterToken is the label used to identify the cluster token
	LabelClusterToken = "k3d.cluster.token"
)

// DefaultRuntimeLabels are labels that are applied to all k3d objects
var DefaultRuntimeLabels = map[string]string{
	"app": "k3d",
}

// Node roles
const (
	// NoRole specifies that a node has no specific role
	NoRole = "noRole"
)
