package e2e

import (
	"fmt"

	"github.com/cloud-native-application/rudrx/e2e"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var (
	envName         = "env-application"
	applicationName = "app-ls-basic"
)

var _ = ginkgo.Describe("Application", func() {
	e2e.RefreshContext("refresh")
	e2e.EnvInitContext("env init", envName)
	e2e.EnvShowContext("env show", envName)
	e2e.EnvSwitchContext("env switch", envName)
	e2e.WorkloadRunContext("run", fmt.Sprintf("vela containerized:run %s -p 80 --image nginx:1.9.4", applicationName))

	ginkgo.Context("ls", func() {
		ginkgo.It("should list all applications", func() {
			output, err := e2e.Exec("vela app:ls")
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(output).To(gomega.ContainSubstring("NAME"))
			gomega.Expect(output).To(gomega.ContainSubstring(applicationName))
		})
	})

	e2e.WorkloadDeleteContext("delete", applicationName)
})
