// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/certificates/v1"
)

type Interface interface {
	CertificateSigningRequests() informers.CertificateSigningRequestInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) CertificateSigningRequests() informers.CertificateSigningRequestInformer {
	return NewCertificateSigningRequestInformer(v.factory)
}
