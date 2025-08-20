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
package client

import (
	"context"

	"github.com/k3d-io/k3d/v5/pkg/k3d"
	"github.com/k3d-io/k3d/v5/pkg/logger"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
	runtimeTypes "github.com/k3d-io/k3d/v5/pkg/runtimes/types"
)

var log = logger.NewLogger().WithFields(map[string]interface{}{"module": "client"})

// VolumeExists checks if a volume with the given name exists
func VolumeExists(ctx context.Context, runtime runtimes.Runtime, name string) (bool, error) {
	log.Tracef("Checking if volume '%s' exists", name)

	volumes, err := runtime.GetVolumes(ctx, runtimeTypes.WithVolumeName(name))
	if err != nil {
		return false, err
	}
	return len(volumes) > 0, nil
}

// GetClusterVolumes returns all volumes associated with a specific cluster
func GetClusterVolumes(ctx context.Context, runtime runtimes.Runtime, clusterName string) ([]string, error) {
	log.Tracef("Getting volumes for cluster '%s'", clusterName)

	return runtime.GetVolumes(ctx, runtimeTypes.WithVolumeLabel(k3d.LabelClusterName, clusterName))
}
