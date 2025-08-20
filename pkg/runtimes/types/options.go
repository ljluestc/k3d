package types

// NodeListOptions defines options for listing nodes
type NodeListOptions struct {
	FilterLabels map[string]string
	FilterName   string
}

// NodeListOpt is a function that modifies NodeListOptions
type NodeListOpt func(*NodeListOptions)

// WithNodeLabel returns a NodeListOpt that adds a label filter
func WithNodeLabel(key, value string) NodeListOpt {
	return func(o *NodeListOptions) {
		if o.FilterLabels == nil {
			o.FilterLabels = make(map[string]string)
		}
		o.FilterLabels[key] = value
	}
}

// WithNodeName returns a NodeListOpt that filters by node name
func WithNodeName(name string) NodeListOpt {
	return func(o *NodeListOptions) {
		o.FilterName = name
	}
}

// VolumeListOptions defines options for listing volumes
type VolumeListOptions struct {
	FilterLabels map[string]string
	FilterName   string
}

// VolumeListOpt is a function that modifies VolumeListOptions
type VolumeListOpt func(*VolumeListOptions)

// WithVolumeLabel returns a VolumeListOpt that adds a label filter
func WithVolumeLabel(key, value string) VolumeListOpt {
	return func(o *VolumeListOptions) {
		if o.FilterLabels == nil {
			o.FilterLabels = make(map[string]string)
		}
		o.FilterLabels[key] = value
	}
}

// WithVolumeName returns a VolumeListOpt that filters by volume name
func WithVolumeName(name string) VolumeListOpt {
	return func(o *VolumeListOptions) {
		o.FilterName = name
	}
}

// NetworkListOptions defines options for listing networks
type NetworkListOptions struct {
	FilterLabels map[string]string
	FilterName   string
}

// NetworkListOpt is a function that modifies NetworkListOptions
type NetworkListOpt func(*NetworkListOptions)

// WithNetworkLabel returns a NetworkListOpt that adds a label filter
func WithNetworkLabel(key, value string) NetworkListOpt {
	return func(o *NetworkListOptions) {
		if o.FilterLabels == nil {
			o.FilterLabels = make(map[string]string)
		}
		o.FilterLabels[key] = value
	}
}

// WithNetworkName returns a NetworkListOpt that filters by network name
func WithNetworkName(name string) NetworkListOpt {
	return func(o *NetworkListOptions) {
		o.FilterName = name
	}
}

// ImageListOptions defines options for listing images
type ImageListOptions struct {
	FilterReference string
}

// ImageListOpt is a function that modifies ImageListOptions
type ImageListOpt func(*ImageListOptions)

// WithImageReference returns an ImageListOpt that filters by image reference
func WithImageReference(reference string) ImageListOpt {
	return func(o *ImageListOptions) {
		o.FilterReference = reference
	}
}
