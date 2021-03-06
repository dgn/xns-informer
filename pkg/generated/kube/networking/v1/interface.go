// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/networking/v1"
)

type Interface interface {
	Ingresses() informers.IngressInformer
	IngressClasses() informers.IngressClassInformer
	NetworkPolicies() informers.NetworkPolicyInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) Ingresses() informers.IngressInformer {
	return NewIngressInformer(v.factory)
}
func (v *version) IngressClasses() informers.IngressClassInformer {
	return NewIngressClassInformer(v.factory)
}
func (v *version) NetworkPolicies() informers.NetworkPolicyInformer {
	return NewNetworkPolicyInformer(v.factory)
}
