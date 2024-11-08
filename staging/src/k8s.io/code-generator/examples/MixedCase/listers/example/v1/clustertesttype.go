/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	examplev1 "k8s.io/code-generator/examples/MixedCase/apis/example/v1"
)

// ClusterTestTypeLister helps list ClusterTestTypes.
// All objects returned here must be treated as read-only.
type ClusterTestTypeLister interface {
	// List lists all ClusterTestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*examplev1.ClusterTestType, err error)
	// Get retrieves the ClusterTestType from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*examplev1.ClusterTestType, error)
	ClusterTestTypeListerExpansion
}

// clusterTestTypeLister implements the ClusterTestTypeLister interface.
type clusterTestTypeLister struct {
	listers.ResourceIndexer[*examplev1.ClusterTestType]
}

// NewClusterTestTypeLister returns a new ClusterTestTypeLister.
func NewClusterTestTypeLister(indexer cache.Indexer) ClusterTestTypeLister {
	return &clusterTestTypeLister{listers.New[*examplev1.ClusterTestType](indexer, examplev1.Resource("clustertesttype"))}
}
