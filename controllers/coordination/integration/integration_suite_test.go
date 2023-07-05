package integration_test

import (
	"testing"
	"time"

	"code.cloudfoundry.org/korifi/tests/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func TestIntegration(t *testing.T) {
	SetDefaultEventuallyTimeout(10 * time.Second)
	SetDefaultEventuallyPollingInterval(250 * time.Millisecond)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var (
	testEnv           *envtest.Environment
	k8sClient         client.Client
	controllersClient client.Client
)

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	testEnv = &envtest.Environment{}

	adminConfig, err := testEnv.Start()
	Expect(err).NotTo(HaveOccurred())

	k8sClient, err = client.New(adminConfig, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())

	controllersUser, err := testEnv.ControlPlane.AddUser(envtest.User{Name: "envtest-controller"}, adminConfig)
	Expect(err).NotTo(HaveOccurred())
	helpers.BindUserToControllersRole(k8sClient, "envtest-controller")

	controllersClient, err = client.New(
		controllersUser.Config(),
		client.Options{Scheme: scheme.Scheme},
	)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(testEnv.Stop()).To(Succeed())
})
