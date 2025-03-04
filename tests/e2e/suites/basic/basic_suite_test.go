package basic

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/giantswarm/clustertest/pkg/logger"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/apptest-framework/pkg/config"
	"github.com/giantswarm/apptest-framework/pkg/state"
	"github.com/giantswarm/apptest-framework/pkg/suite"
)

const (
	isUpgrade = false
	namespace = "kube-system"
)

func TestBasic(t *testing.T) {
	suite.New(config.MustLoad("../../config.yaml")).
		WithInstallNamespace(namespace).
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		AfterClusterReady(func() {
			// Do any pre-install checks here (ensure the cluster has needed pre-reqs)
		}).
		BeforeUpgrade(func() {
			// Perform any checks between installing the latest released version
			// and upgrading it to the version to test
			// E.g. ensure that the initial install has completed and has settled before upgrading
		}).
		Tests(func() {
			It("should run Nvidia validator", func() {
				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).Should(Succeed())

				Eventually(func() (bool, error) {
					return validatorDaemonSetReady(wcClient, "nvidia-operator-validator")
				}).
					WithTimeout(15 * time.Minute).
					WithPolling(15 * time.Second).
					Should(Succeed())
			})
		}).
		Run(t, "Basic Test")
}

func validatorDaemonSetReady(wcClient client.Client, daemonSetName string) (bool, error) {
	logger.Log("Checking if nvidia-operator-validator DaemonSet is in running state")
	daemonSet := appsv1.DaemonSet{}
	err := wcClient.Get(state.GetContext(), types.NamespacedName{Name: daemonSetName, Namespace: namespace}, &daemonSet)
	if err != nil {
		logger.Log("Failed to get DaemonSet : %v", err)
		return false, err
	}

	if daemonSet.Status.DesiredNumberScheduled == daemonSet.Status.NumberReady {
		return true, nil
	}
	return false, nil
}
