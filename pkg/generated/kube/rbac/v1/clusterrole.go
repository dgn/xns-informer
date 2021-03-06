// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/rbac/v1"
	informers "k8s.io/client-go/informers/rbac/v1"
	listers "k8s.io/client-go/listers/rbac/v1"
	"k8s.io/client-go/tools/cache"
)

type clusterRoleInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ClusterRoleInformer = &clusterRoleInformer{}

func NewClusterRoleInformer(f xnsinformers.SharedInformerFactory) informers.ClusterRoleInformer {
	resource := v1.SchemeGroupVersion.WithResource("clusterroles")
	informer := f.ClusterResource(resource).Informer()

	return &clusterRoleInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1.ClusterRole{}),
	}
}

func (i *clusterRoleInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *clusterRoleInformer) Lister() listers.ClusterRoleLister {
	return listers.NewClusterRoleLister(i.informer.GetIndexer())
}
