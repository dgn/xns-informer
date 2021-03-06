// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	informers "k8s.io/client-go/informers/extensions/v1beta1"
	listers "k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type replicaSetInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ReplicaSetInformer = &replicaSetInformer{}

func NewReplicaSetInformer(f xnsinformers.SharedInformerFactory) informers.ReplicaSetInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("replicasets")
	informer := f.NamespacedResource(resource).Informer()

	return &replicaSetInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.ReplicaSet{}),
	}
}

func (i *replicaSetInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *replicaSetInformer) Lister() listers.ReplicaSetLister {
	return listers.NewReplicaSetLister(i.informer.GetIndexer())
}
