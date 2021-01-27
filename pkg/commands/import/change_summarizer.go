package _import

import (
	"fmt"
	"strings"

	kpack "github.com/pivotal/kpack/pkg/client/clientset/versioned"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"

	importpkg "github.com/pivotal/build-service-cli/pkg/import"
	buildk8s "github.com/pivotal/build-service-cli/pkg/k8s"
	"github.com/pivotal/build-service-cli/pkg/lifecycle"
)

type changeSummarizer struct {
	descriptor importpkg.DependencyDescriptor
	differ     *importpkg.ImportDiffer
	k8sClient  k8s.Interface
	kpClient   kpack.Interface

	hasChanges bool
	changes    strings.Builder
	diffs      strings.Builder
}

func summarizeChange(
	descriptor importpkg.DependencyDescriptor,
	differ *importpkg.ImportDiffer,
	clientSet buildk8s.ClientSet) (bool, string, error) {

	s := &changeSummarizer{
		descriptor: descriptor,
		differ:     differ,
		k8sClient:  clientSet.K8sClient,
		kpClient:   clientSet.KpackClient,
	}
	return s.summary()
}

func (cp *changeSummarizer) summary() (bool, string, error) {
	cp.changes.WriteString("Changes\n\n")

	if err := cp.writeLifecycleChange(); err != nil {
		return false, "", err
	}
	if err := cp.writeClusterStoresChange(); err != nil {
		return false, "", err
	}
	if err := cp.writeClusterStacksChange(); err != nil {
		return false, "", err
	}
	if err := cp.writeClusterBuildersChange(); err != nil {
		return false, "", err
	}

	return cp.hasChanges, cp.changes.String(), nil
}

func (cp *changeSummarizer) writeLifecycleChange() error {
	if cp.descriptor.HasLifecycleImage() {
		oldImg, err := lifecycle.GetImage(cp.k8sClient)
		if err != nil {
			return err
		}

		diff, err := cp.differ.DiffLifecycle(oldImg, cp.descriptor.Lifecycle.Image)
		if err != nil {
			return err
		}

		if err = cp.writeDiff(diff); err != nil {
			return err
		}
	}

	cp.writeChange("Lifecycle")
	return nil
}

func (cp *changeSummarizer) writeClusterStoresChange() error {
	for _, cs := range cp.descriptor.ClusterStores {
		curStore, err := cp.kpClient.KpackV1alpha1().ClusterStores().Get(cs.Name, metav1.GetOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			return err
		}
		if k8serrors.IsNotFound(err) {
			curStore = nil
		}

		diff, err := cp.differ.DiffClusterStore(curStore, cs)
		if err != nil {
			return err
		}
		if err = cp.writeDiff(diff); err != nil {
			return err
		}
	}

	cp.writeChange("ClusterStores")
	return nil
}

func (cp *changeSummarizer) writeClusterStacksChange() error {
	for _, cs := range cp.descriptor.GetClusterStacks() {
		curStack, err := cp.kpClient.KpackV1alpha1().ClusterStacks().Get(cs.Name, metav1.GetOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			return err
		}
		if k8serrors.IsNotFound(err) {
			curStack = nil
		}

		diff, err := cp.differ.DiffClusterStack(curStack, cs)
		if err != nil {
			return err
		}
		if err = cp.writeDiff(diff); err != nil {
			return err
		}
	}

	cp.writeChange("ClusterStacks")
	return nil
}

func (cp *changeSummarizer) writeClusterBuildersChange() error {
	for _, cb := range cp.descriptor.GetClusterBuilders() {
		curBuilder, err := cp.kpClient.KpackV1alpha1().ClusterBuilders().Get(cb.Name, metav1.GetOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			return err
		}
		if k8serrors.IsNotFound(err) {
			curBuilder = nil
		}

		diff, err := cp.differ.DiffClusterBuilder(curBuilder, cb)
		if err != nil {
			return err
		}
		if err = cp.writeDiff(diff); err != nil {
			return err
		}
	}

	cp.writeChange("ClusterBuilders")
	return nil
}

func (cp *changeSummarizer) writeChange(header string) {
	cp.changes.WriteString(fmt.Sprintf("%s\n\n", header))

	change := cp.diffs.String()
	if change == "" {
		cp.changes.WriteString("No Changes\n\n")
	} else {
		cp.changes.WriteString(change)
		cp.hasChanges = true
	}

	cp.diffs.Reset()
}

func (cp *changeSummarizer) writeDiff(diff string) error {
	var err error
	if diff != "" {
		_, err = cp.diffs.WriteString(fmt.Sprintf("%s\n\n", diff))
	}
	return err
}
