package runtimes

import (
	"context"
	
	"github.com/k3d-io/k3d/v5/pkg/types"
)

// Runtime defines the interface that a runtime (like Docker) has to implement
type Runtime interface {
	GetNodesByLabel(ctx context.Context, labels map[string]string) ([]*types.Node, error)
	DeleteNode(ctx context.Context, node *types.Node) error
	CreateNode(ctx context.Context, node *types.Node) error
	GetVolumes(ctx context.Context, filters ...string) ([]string, error)
}

// ContainerClient defines an interface for container operations
type ContainerClient interface {
	Close() error
}
