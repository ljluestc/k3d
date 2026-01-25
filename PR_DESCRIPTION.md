<!-- 
Hi there, have an early THANK YOU for your contribution!
k3d is a community-driven project, so we really highly appreciate any support.
Please make sure, you've read our Code of Conduct and the Contributing Guidelines :)
- Code of Conduct: https://github.com/k3d-io/k3d/blob/main/CODE_OF_CONDUCT.md
- Contributing Guidelines: https://github.com/k3d-io/k3d/blob/main/CONTRIBUTING.md
-->

# What

### Fixes
1.  **Panic in `k3d cluster list -o yaml`**:
    -   Refactored `Loadbalancer` struct in `pkg/types/loadbalancer.go` to embed `Node` by value instead of by pointer.
    -   Updated `cmd/cluster/clusterList.go` to use `*k3d.Cluster` in `jsonOutput` to avoid value copying issues.
    -   Updated usages of `ServerLoadBalancer.Node` in `pkg/client/cluster.go`, `pkg/client/loadbalancer.go`, `pkg/client/node.go`, `pkg/client/ports.go`, and `pkg/config/transform.go` to reflect the struct change (taking address of `Node` where a pointer is expected).

2.  **Race Condition in Cluster Creation**:
    -   Modified `pkg/client/cluster.go` to run `EnsureToolsNode` sequentially after `ClusterCreate` instead of in parallel. This avoids race conditions observed with rootless Podman where the tools node creation could interfere with cluster node creation or network operations.

### Features
3.  **Internal API Flag for Kubeconfig**:
    -   Added `--internal` flag to `k3d kubeconfig get` and `k3d kubeconfig merge` commands.
    -   Updated `pkg/client/kubeconfig.go` to support `UseInternalAPI` option, allowing retrieval of kubeconfig with the internal container IP/port instead of the host-mapped address.

# Why

-   **Panic Fix**: Fixes issue #1098. The panic was caused by the `Loadbalancer` struct embedding `*Node`. When `mapstructure` or YAML marshaling attempted to process the embedded pointer field, it led to issues, particularly when the node might not have been fully initialized or during the flattening process for output. Embedding `Node` by value resolves this marshaling issue.
-   **Race Condition**: Improves stability of cluster creation, especially in rootless Podman environments.
-   **Internal Flag**: Useful for workflows where the kubeconfig is needed for usage inside the same docker network or environments where the internal IP is preferred.

# Implications

-   **Breaking Change**: 
    -   `k3d.Loadbalancer` struct definition changed (internal Go API).
    -   No breaking changes for CLI users, only fixes and additive features.
-   **Affected Areas**:
    -   Cluster listing (serialization)
    -   Cluster creation (loadbalancer preparation, tools node creation)
    -   Loadbalancer configuration updates
    -   Node editing
    -   Kubeconfig generation

# How to Test

### 1. Test Panic Fix
1.  Create a cluster with a loadbalancer (default):
    ```bash
    k3d cluster create test-cluster
    ```
2.  List clusters with YAML output:
    ```bash
    k3d cluster list -o yaml
    ```
    **Expected Result**: The command should output the cluster details in YAML format without panicking.
3.  List clusters with JSON output:
    ```bash
    k3d cluster list -o json
    ```
    **Expected Result**: The command should output the cluster details in JSON format without panicking.

### 2. Test Internal Kubeconfig
1.  Get kubeconfig with internal flag:
    ```bash
    k3d kubeconfig get test-cluster --internal
    ```
    **Expected Result**: The output kubeconfig should contain the internal IP address (e.g., `172.x.x.x` or similar depending on network) instead of `0.0.0.0` or `127.0.0.1`.

### 3. Cleanup
```bash
k3d cluster delete test-cluster
```
