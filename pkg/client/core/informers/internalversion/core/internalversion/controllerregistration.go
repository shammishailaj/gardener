// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	core "github.com/gardener/gardener/pkg/apis/core"
	clientsetinternalversion "github.com/gardener/gardener/pkg/client/core/clientset/internalversion"
	internalinterfaces "github.com/gardener/gardener/pkg/client/core/informers/internalversion/internalinterfaces"
	internalversion "github.com/gardener/gardener/pkg/client/core/listers/core/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ControllerRegistrationInformer provides access to a shared informer and lister for
// ControllerRegistrations.
type ControllerRegistrationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ControllerRegistrationLister
}

type controllerRegistrationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewControllerRegistrationInformer constructs a new informer for ControllerRegistration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewControllerRegistrationInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredControllerRegistrationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredControllerRegistrationInformer constructs a new informer for ControllerRegistration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredControllerRegistrationInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Core().ControllerRegistrations().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Core().ControllerRegistrations().Watch(options)
			},
		},
		&core.ControllerRegistration{},
		resyncPeriod,
		indexers,
	)
}

func (f *controllerRegistrationInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredControllerRegistrationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *controllerRegistrationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&core.ControllerRegistration{}, f.defaultInformer)
}

func (f *controllerRegistrationInformer) Lister() internalversion.ControllerRegistrationLister {
	return internalversion.NewControllerRegistrationLister(f.Informer().GetIndexer())
}
