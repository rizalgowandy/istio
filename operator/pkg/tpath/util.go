// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
util.go contains utility function for dealing with trees.
*/

package tpath

import (
	"sigs.k8s.io/yaml"

	"istio.io/istio/operator/pkg/util"
)

// AddSpecRoot adds a root node called "spec" to the given tree and returns the resulting tree.
func AddSpecRoot(tree string) (string, error) {
	t, nt := make(map[string]any), make(map[string]any)
	if err := yaml.Unmarshal([]byte(tree), &t); err != nil {
		return "", err
	}
	nt["spec"] = t
	out, err := yaml.Marshal(nt)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// GetConfigSubtree returns the subtree at the given path.
func GetConfigSubtree(manifest, path string) (string, error) {
	root := make(map[string]any)
	if err := yaml.Unmarshal([]byte(manifest), &root); err != nil {
		return "", err
	}

	nc, _, err := GetPathContext(root, util.PathFromString(path), false)
	if err != nil {
		return "", err
	}
	out, err := yaml.Marshal(nc.Node)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
