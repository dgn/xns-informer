// Code generated by xns-informer-gen. DO NOT EDIT.
package v1

import (
	"fmt"

	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type podInformer struct {
	factory xnsinformers.InformerFactory
}

var _ informers.PodInformer = &podInformer{}

func (f *podInformer) resource() schema.GroupVersionResource {
	return v1.SchemeGroupVersion.WithResource("pods")
}

func (f *podInformer) Informer() cache.SharedIndexInformer {
	return f.factory.ForResource(f.resource()).Informer()
}

func (f *podInformer) Lister() listers.PodLister {
	return &podLister{lister: f.factory.ForResource(f.resource()).Lister()}
}

type podLister struct {
	lister cache.GenericLister
}

var _ listers.PodLister = &podLister{}

func (l *podLister) List(selector labels.Selector) (res []*v1.Pod, err error) {
	return listPod(l.lister, selector)
}

func (l *podLister) Pods(namespace string) listers.PodNamespaceLister {
	return &podNamespaceLister{lister: l.lister.ByNamespace(namespace)}
}

type podNamespaceLister struct {
	lister cache.GenericNamespaceLister
}

var _ listers.PodNamespaceLister = &podNamespaceLister{}

func (l *podNamespaceLister) List(selector labels.Selector) (res []*v1.Pod, err error) {
	return listPod(l.lister, selector)
}

func (l *podNamespaceLister) Get(name string) (*v1.Pod, error) {
	obj, err := l.lister.Get(name)
	if err != nil {
		return nil, err
	}

	return convertPod(obj)
}

func listPod(l xnsinformers.SimpleLister, s labels.Selector) (res []*v1.Pod, err error) {
	objects, err := l.List(s)
	if err != nil {
		return nil, err
	}

	for _, obj := range objects {
		o, err := convertPod(obj)
		if err != nil {
			return nil, err
		}

		res = append(res, o)
	}

	return res, nil
}

func convertPod(obj runtime.Object) (*v1.Pod, error) {
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("unstructured conversion failed")
	}

	out := &v1.Pod{}
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}