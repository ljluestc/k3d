package cluster

import (
	"testing"

	"github.com/k3d-io/k3d/v5/pkg/types"
)

func TestPrintClusters_YAML_NoPanic(t *testing.T) {
	// Create a dummy cluster with a loadbalancer
	cluster := &types.Cluster{
		Name: "test-cluster",
		Nodes: []*types.Node{
			{
				Name: "test-server-0",
				Role: types.ServerRole,
			},
		},
		ServerLoadBalancer: &types.Loadbalancer{
			Node: types.Node{
				Name: "test-lb",
				Role: types.LoadBalancerRole,
			},
			Config: &types.LoadbalancerConfig{
				Ports: map[string][]string{
					"80.tcp": {"test-server-0"},
				},
			},
		},
	}

	clusters := []*types.Cluster{cluster}
	flags := clusterFlags{
		output: "yaml",
	}

	// This function should not panic
	// We capture stdout to avoid cluttering test output, but mainly we rely on the test passing without panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PrintClusters panicked: %v", r)
			}
		}()
		PrintClusters(clusters, flags)
	}()
}

func TestPrintClusters_JSON_NoPanic(t *testing.T) {
	// Create a dummy cluster with a loadbalancer
	cluster := &types.Cluster{
		Name: "test-cluster",
		Nodes: []*types.Node{
			{
				Name: "test-server-0",
				Role: types.ServerRole,
			},
		},
		ServerLoadBalancer: &types.Loadbalancer{
			Node: types.Node{
				Name: "test-lb",
				Role: types.LoadBalancerRole,
			},
			Config: &types.LoadbalancerConfig{
				Ports: map[string][]string{
					"80.tcp": {"test-server-0"},
				},
			},
		},
	}

	clusters := []*types.Cluster{cluster}
	flags := clusterFlags{
		output: "json",
	}

	// This function should not panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PrintClusters panicked: %v", r)
			}
		}()
		PrintClusters(clusters, flags)
	}()
}
