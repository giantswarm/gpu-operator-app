package basic

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/apptest-framework/pkg/config"
	"github.com/giantswarm/apptest-framework/pkg/state"
	"github.com/giantswarm/apptest-framework/pkg/suite"
	"github.com/giantswarm/clustertest/pkg/logger"
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
			It("should have all GPU operator components running", func() {
				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).Should(Succeed())

				// Define components to check - name and type
				components := map[string]string{
					// DaemonSets
					"nvidia-operator-validator":                  "daemonset",
					"gpu-feature-discovery":                      "daemonset",
					"nvidia-dcgm-exporter":                       "daemonset",
					"nvidia-device-plugin-daemonset":             "daemonset",
					"gpu-operator-node-feature-discovery-worker": "daemonset",

					// Deployments
					"gpu-operator": "deployment",
					"gpu-operator-node-feature-discovery-master": "deployment",
				}

				// Check each component
				for name, kind := range components {
					By(fmt.Sprintf("Checking if %s is running", name))
					var isReady bool

					Eventually(func() bool {
						if kind == "daemonset" {
							isReady = validatorDaemonSetReady(wcClient, name)
						} else if kind == "deployment" {
							isReady = deploymentReady(wcClient, name)
						}
						return isReady
					}).
						WithTimeout(7*time.Minute).
						WithPolling(15*time.Second).
						Should(BeTrue(), fmt.Sprintf("%s should be ready", name))
				}
			})
		}).
		Run(t, "Basic Test")
}

func validatorDaemonSetReady(wcClient client.Client, daemonSetName string) bool {
	logger.Log("Checking if %s DaemonSet is in running state", daemonSetName)
	daemonSet := appsv1.DaemonSet{}
	err := wcClient.Get(state.GetContext(), types.NamespacedName{Name: daemonSetName, Namespace: namespace}, &daemonSet)
	if err != nil {
		logger.Log("%s DaemonSet is not ready", daemonSetName)
		return false
	}

	return daemonSet.Status.DesiredNumberScheduled == daemonSet.Status.NumberReady && daemonSet.Status.DesiredNumberScheduled > 0
}

func deploymentReady(wcClient client.Client, deploymentName string) bool {
	logger.Log("Checking if %s Deployment is in running state", deploymentName)
	deployment := appsv1.Deployment{}
	err := wcClient.Get(state.GetContext(), types.NamespacedName{Name: deploymentName, Namespace: namespace}, &deployment)
	if err != nil {
		logger.Log("%s Deployment is not ready", deploymentName)
		return false
	}

	// A deployment is considered ready when all of its replicas are ready
	return deployment.Status.ReadyReplicas == deployment.Status.Replicas && deployment.Status.Replicas > 0
}
