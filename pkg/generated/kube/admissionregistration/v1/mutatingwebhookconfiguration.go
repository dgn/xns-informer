// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/admissionregistration/v1"
	informers "k8s.io/client-go/informers/admissionregistration/v1"
	listers "k8s.io/client-go/listers/admissionregistration/v1"
	"k8s.io/client-go/tools/cache"
)

type mutatingWebhookConfigurationInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.MutatingWebhookConfigurationInformer = &mutatingWebhookConfigurationInformer{}

func NewMutatingWebhookConfigurationInformer(f xnsinformers.SharedInformerFactory) informers.MutatingWebhookConfigurationInformer {
	resource := v1.SchemeGroupVersion.WithResource("mutatingwebhookconfigurations")
	informer := f.ClusterResource(resource).Informer()

	return &mutatingWebhookConfigurationInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1.MutatingWebhookConfiguration{}),
	}
}

func (i *mutatingWebhookConfigurationInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *mutatingWebhookConfigurationInformer) Lister() listers.MutatingWebhookConfigurationLister {
	return listers.NewMutatingWebhookConfigurationLister(i.informer.GetIndexer())
}
