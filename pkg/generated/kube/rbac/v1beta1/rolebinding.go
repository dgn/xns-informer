// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/rbac/v1beta1"
	informers "k8s.io/client-go/informers/rbac/v1beta1"
	listers "k8s.io/client-go/listers/rbac/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type roleBindingInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.RoleBindingInformer = &roleBindingInformer{}

func NewRoleBindingInformer(f xnsinformers.SharedInformerFactory) informers.RoleBindingInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("rolebindings")
	informer := f.NamespacedResource(resource).Informer()

	return &roleBindingInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.RoleBinding{}),
	}
}

func (i *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *roleBindingInformer) Lister() listers.RoleBindingLister {
	return listers.NewRoleBindingLister(i.informer.GetIndexer())
}
