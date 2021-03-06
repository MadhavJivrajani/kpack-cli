package _import

import (
	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"knative.dev/pkg/apis/duck"
)

type builderWaitable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status v1alpha1.BuilderStatus `json:"status"`
}

func builderHasResolved(expectedStoreGen, expectedStackGen int64) func(event watch.Event) (bool, error) {
	return func(e watch.Event) (bool, error) {
		u := &unstructured.Unstructured{}
		var err error
		u.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(e.Object)
		if err != nil {
			return false, err
		}

		bw := &builderWaitable{}
		if err := duck.FromUnstructured(u, bw); err != nil {
			return false, err
		}

		if (bw.Status.ObservedStoreGeneration != 0 && bw.Status.ObservedStoreGeneration < expectedStoreGen) || // ObservedStoreGeneration is 0 when kpack does not support it
			(bw.Status.ObservedStackGeneration != 0 && bw.Status.ObservedStackGeneration < expectedStackGen) { // ObservedStackGeneration is 0 when kpack does not support it
			return false, nil // still waiting on update
		}

		return true, nil
	}
}
