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

package v1beta1

import (
	v1beta1 "github.com/linkerd/linkerd2/controller/gen/apis/serverauthorization/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServerAuthorizationLister helps list ServerAuthorizations.
// All objects returned here must be treated as read-only.
type ServerAuthorizationLister interface {
	// List lists all ServerAuthorizations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ServerAuthorization, err error)
	// ServerAuthorizations returns an object that can list and get ServerAuthorizations.
	ServerAuthorizations(namespace string) ServerAuthorizationNamespaceLister
	ServerAuthorizationListerExpansion
}

// serverAuthorizationLister implements the ServerAuthorizationLister interface.
type serverAuthorizationLister struct {
	indexer cache.Indexer
}

// NewServerAuthorizationLister returns a new ServerAuthorizationLister.
func NewServerAuthorizationLister(indexer cache.Indexer) ServerAuthorizationLister {
	return &serverAuthorizationLister{indexer: indexer}
}

// List lists all ServerAuthorizations in the indexer.
func (s *serverAuthorizationLister) List(selector labels.Selector) (ret []*v1beta1.ServerAuthorization, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServerAuthorization))
	})
	return ret, err
}

// ServerAuthorizations returns an object that can list and get ServerAuthorizations.
func (s *serverAuthorizationLister) ServerAuthorizations(namespace string) ServerAuthorizationNamespaceLister {
	return serverAuthorizationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServerAuthorizationNamespaceLister helps list and get ServerAuthorizations.
// All objects returned here must be treated as read-only.
type ServerAuthorizationNamespaceLister interface {
	// List lists all ServerAuthorizations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ServerAuthorization, err error)
	// Get retrieves the ServerAuthorization from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ServerAuthorization, error)
	ServerAuthorizationNamespaceListerExpansion
}

// serverAuthorizationNamespaceLister implements the ServerAuthorizationNamespaceLister
// interface.
type serverAuthorizationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServerAuthorizations in the indexer for a given namespace.
func (s serverAuthorizationNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ServerAuthorization, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServerAuthorization))
	})
	return ret, err
}

// Get retrieves the ServerAuthorization from the indexer for a given namespace and name.
func (s serverAuthorizationNamespaceLister) Get(name string) (*v1beta1.ServerAuthorization, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("serverauthorization"), name)
	}
	return obj.(*v1beta1.ServerAuthorization), nil
}
