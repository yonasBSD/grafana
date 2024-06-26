// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by lister-gen. DO NOT EDIT.

package v0alpha1

import (
	v0alpha1 "github.com/grafana/grafana/pkg/apis/alerting_notifications/v0alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TimeIntervalLister helps list TimeIntervals.
// All objects returned here must be treated as read-only.
type TimeIntervalLister interface {
	// List lists all TimeIntervals in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v0alpha1.TimeInterval, err error)
	// TimeIntervals returns an object that can list and get TimeIntervals.
	TimeIntervals(namespace string) TimeIntervalNamespaceLister
	TimeIntervalListerExpansion
}

// timeIntervalLister implements the TimeIntervalLister interface.
type timeIntervalLister struct {
	indexer cache.Indexer
}

// NewTimeIntervalLister returns a new TimeIntervalLister.
func NewTimeIntervalLister(indexer cache.Indexer) TimeIntervalLister {
	return &timeIntervalLister{indexer: indexer}
}

// List lists all TimeIntervals in the indexer.
func (s *timeIntervalLister) List(selector labels.Selector) (ret []*v0alpha1.TimeInterval, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v0alpha1.TimeInterval))
	})
	return ret, err
}

// TimeIntervals returns an object that can list and get TimeIntervals.
func (s *timeIntervalLister) TimeIntervals(namespace string) TimeIntervalNamespaceLister {
	return timeIntervalNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TimeIntervalNamespaceLister helps list and get TimeIntervals.
// All objects returned here must be treated as read-only.
type TimeIntervalNamespaceLister interface {
	// List lists all TimeIntervals in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v0alpha1.TimeInterval, err error)
	// Get retrieves the TimeInterval from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v0alpha1.TimeInterval, error)
	TimeIntervalNamespaceListerExpansion
}

// timeIntervalNamespaceLister implements the TimeIntervalNamespaceLister
// interface.
type timeIntervalNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TimeIntervals in the indexer for a given namespace.
func (s timeIntervalNamespaceLister) List(selector labels.Selector) (ret []*v0alpha1.TimeInterval, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v0alpha1.TimeInterval))
	})
	return ret, err
}

// Get retrieves the TimeInterval from the indexer for a given namespace and name.
func (s timeIntervalNamespaceLister) Get(name string) (*v0alpha1.TimeInterval, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v0alpha1.Resource("timeinterval"), name)
	}
	return obj.(*v0alpha1.TimeInterval), nil
}
