package client

import (
	"net/netip"
	"testing"

	k3d "github.com/k3d-io/k3d/v5/pkg/types"
)

func Test_kubeconfigSelectAPIEndpoint_Default(t *testing.T) {
	server0 := &k3d.Node{
		Name: "k3d-test-server-0",
		RuntimeLabels: map[string]string{
			k3d.LabelServerAPIPort: "12345",
			k3d.LabelServerAPIHost: "127.0.0.1",
		},
	}
	server1 := &k3d.Node{
		Name: "k3d-test-server-1",
		RuntimeLabels: map[string]string{
			k3d.LabelServerAPIPort: "54321",
			k3d.LabelServerAPIHost: "localhost",
		},
	}

	chosen, host, port, err := kubeconfigSelectAPIEndpoint([]*k3d.Node{server1, server0}, &WriteKubeConfigOptions{UseInternalAPI: false})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chosen != server1 {
		t.Fatalf("expected server1 to be chosen")
	}
	if host != "localhost" {
		t.Fatalf("expected host=localhost, got %q", host)
	}
	if port != "54321" {
		t.Fatalf("expected port=54321, got %q", port)
	}
}

func Test_kubeconfigSelectAPIEndpoint_InternalPrefersLoadBalancerName(t *testing.T) {
	server := &k3d.Node{
		Name: "k3d-test-server-0",
		RuntimeLabels: map[string]string{
			k3d.LabelServerAPIPort:      "12345",
			k3d.LabelServerAPIHost:      "127.0.0.1",
			k3d.LabelServerLoadBalancer: "k3d-test-serverlb",
		},
	}

	chosen, host, port, err := kubeconfigSelectAPIEndpoint([]*k3d.Node{server}, &WriteKubeConfigOptions{UseInternalAPI: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chosen != server {
		t.Fatalf("expected server to be chosen")
	}
	if host != "k3d-test-serverlb" {
		t.Fatalf("expected host=k3d-test-serverlb, got %q", host)
	}
	if port != k3d.DefaultAPIPort {
		t.Fatalf("expected port=%s, got %q", k3d.DefaultAPIPort, port)
	}
}

func Test_kubeconfigSelectAPIEndpoint_InternalFallsBackToNodeName(t *testing.T) {
	server := &k3d.Node{
		Name: "k3d-test-server-0",
		RuntimeLabels: map[string]string{
			k3d.LabelServerAPIPort: "12345",
			k3d.LabelServerAPIHost: "127.0.0.1",
		},
	}

	chosen, host, port, err := kubeconfigSelectAPIEndpoint([]*k3d.Node{server}, &WriteKubeConfigOptions{UseInternalAPI: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chosen != server {
		t.Fatalf("expected server to be chosen")
	}
	if host != "k3d-test-server-0" {
		t.Fatalf("expected host=k3d-test-server-0, got %q", host)
	}
	if port != k3d.DefaultAPIPort {
		t.Fatalf("expected port=%s, got %q", k3d.DefaultAPIPort, port)
	}
}

func Test_kubeconfigSelectAPIEndpoint_InternalFallsBackToNodeIP(t *testing.T) {
	server := &k3d.Node{
		RuntimeLabels: map[string]string{
			k3d.LabelServerAPIPort: "12345",
			k3d.LabelServerAPIHost: "127.0.0.1",
		},
		IP: k3d.NodeIP{IP: netip.MustParseAddr("10.0.0.10")},
	}

	chosen, host, port, err := kubeconfigSelectAPIEndpoint([]*k3d.Node{server}, &WriteKubeConfigOptions{UseInternalAPI: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if chosen != server {
		t.Fatalf("expected server to be chosen")
	}
	if host != "10.0.0.10" {
		t.Fatalf("expected host=10.0.0.10, got %q", host)
	}
	if port != k3d.DefaultAPIPort {
		t.Fatalf("expected port=%s, got %q", k3d.DefaultAPIPort, port)
	}
}
