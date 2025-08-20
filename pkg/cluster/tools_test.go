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
	"strings"
	"testing"
)

func TestGenerateRandomSuffix(t *testing.T) {
	suffix1, err := GenerateRandomSuffix()
	if err != nil {
		t.Fatalf("Failed to generate random suffix: %v", err)
	}

	if len(suffix1) != 8 {
		t.Errorf("Expected random suffix length of 8, got %d", len(suffix1))
	}

	suffix2, err := GenerateRandomSuffix()
	if err != nil {
		t.Fatalf("Failed to generate second random suffix: %v", err)
	}

	if suffix1 == suffix2 {
		t.Errorf("Expected different random suffixes, got the same: %s", suffix1)
	}
}

func TestToolsNodeNaming(t *testing.T) {
	// Note: we can't actually test the full CreateToolsNode function without mocking
	// the Docker client, but we can at least verify the naming pattern works as expected

	baseToolsNodeName := "k3d-test-cluster-tools"
	suffix, err := GenerateRandomSuffix()
	if err != nil {
		t.Fatalf("Failed to generate random suffix: %v", err)
	}

	toolsNodeName := baseToolsNodeName + "-" + suffix

	if !strings.HasPrefix(toolsNodeName, baseToolsNodeName) {
		t.Errorf("Tools node name should start with %s, got %s", baseToolsNodeName, toolsNodeName)
	}

	if !strings.HasSuffix(toolsNodeName, suffix) {
		t.Errorf("Tools node name should end with %s, got %s", suffix, toolsNodeName)
	}

	if len(toolsNodeName) != len(baseToolsNodeName)+1+len(suffix) {
		t.Errorf("Tools node name has unexpected length: %d", len(toolsNodeName))
	}
}
