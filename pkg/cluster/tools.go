/*
Copyright © 2020-2023 The k3d Author(s)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cluster

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/k3d-io/k3d/v5/pkg/k3d"
	l "github.com/k3d-io/k3d/v5/pkg/logger"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
	"github.com/k3d-io/k3d/v5/pkg/types"
)

// GenerateRandomSuffix generates a random suffix for container names
func GenerateRandomSuffix() (string, error) {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate random suffix: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// CreateToolsNode creates the tools node for a cluster with conflict handling
func CreateToolsNode(ctx context.Context, runtime runtimes.Runtime, cluster *types.Cluster) (*types.Node, error) {
	// Check for existing tools containers with the base name pattern
	baseToolsNodeName := fmt.Sprintf("%s-%s-tools", k3d.DefaultObjectNamePrefix, cluster.Name)

	// Find existing nodes with the base name pattern
	existingNodes, err := runtime.GetNodesByLabel(ctx, map[string]string{
		"name": baseToolsNodeName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list existing tools containers: %w", err)
	}

	// Remove existing tools containers to avoid conflicts
	for _, node := range existingNodes {
		l.Log().Infof("Removing existing tools container %s", node.Name)
		if err := runtime.DeleteNode(ctx, node); err != nil {
			return nil, fmt.Errorf("failed to remove existing tools container %s: %w", node.Name, err)
		}
	}

	// Generate a unique suffix for the tools node name
	suffix, err := GenerateRandomSuffix()
	if err != nil {
		return nil, fmt.Errorf("failed to generate unique suffix for tools node: %w", err)
	}
	toolsNodeName := fmt.Sprintf("%s-%s", baseToolsNodeName, suffix)

	// Create the tools node with the unique name
	node := &types.Node{
		Name:  toolsNodeName,
		Role:  "noRole",                                                  // Use string directly instead of constant
		Image: fmt.Sprintf("%s:%s", k3d.DefaultToolsImageRepo, "latest"), // Use literal string
		Volumes: []string{
			fmt.Sprintf("%s-%s-images:/k3d/images", k3d.DefaultObjectNamePrefix, cluster.Name),
			"/var/run/docker.sock:/var/run/docker.sock",
		},
		Env:     []string{},
		Restart: false,
		Labels: map[string]string{
			k3d.LabelClusterName:  cluster.Name,
			k3d.LabelRole:         k3d.NoRole,
			k3d.LabelClusterToken: cluster.Token,
		},
	}

	// Add default labels
	for k, v := range k3d.DefaultRuntimeLabels {
		node.Labels[k] = v
	}

	// Create the tools node
	if err := runtime.CreateNode(ctx, node); err != nil {
		return nil, fmt.Errorf("failed to create tools node '%s': %w", node.Name, err)
	}

	l.Log().Infof("Created tools node '%s'", node.Name)
	return node, nil
}
