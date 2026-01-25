# What
Update `jsonOutput` struct in `cmd/cluster/clusterList.go` to embed `*k3d.Cluster` instead of partial `k3d.Cluster`.

# Why
Fixes #1098. Resolves `panic: Option ,inline needs a struct value field` when running `k3d cluster ls -o yaml`, caused by `gopkg.in/yaml.v2` handling of inline embedded structs.
