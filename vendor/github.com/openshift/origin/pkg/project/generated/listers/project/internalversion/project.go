// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	project "github.com/openshift/origin/pkg/project/apis/project"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProjectLister helps list Projects.
type ProjectLister interface {
	// List lists all Projects in the indexer.
	List(selector labels.Selector) (ret []*project.Project, err error)
	// Get retrieves the Project from the index for a given name.
	Get(name string) (*project.Project, error)
	ProjectListerExpansion
}

// projectLister implements the ProjectLister interface.
type projectLister struct {
	indexer cache.Indexer
}

// NewProjectLister returns a new ProjectLister.
func NewProjectLister(indexer cache.Indexer) ProjectLister {
	return &projectLister{indexer: indexer}
}

// List lists all Projects in the indexer.
func (s *projectLister) List(selector labels.Selector) (ret []*project.Project, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*project.Project))
	})
	return ret, err
}

// Get retrieves the Project from the index for a given name.
func (s *projectLister) Get(name string) (*project.Project, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(project.Resource("project"), name)
	}
	return obj.(*project.Project), nil
}