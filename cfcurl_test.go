package cfcurl_test

import (
	"bufio"
	"os"

	. "github.com/krujos/cfcurl"

	"github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cfcurl", func() {
	var fakeCliConnection *fakes.FakeCliConnection
	var v2apps []string
	var file *os.File

	Describe("an api that is not depricated", func() {
		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			file, err := os.Open("apps.json")
			if err != nil {
				Fail("Could not open apps.json")
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				v2apps = append(v2apps, scanner.Text())
			}
			if scanner.Err() != nil {
				Fail("Failed to read lines from file")
			}
		})

		AfterEach(func() {
			file.Close()
		})

		It("Should return a string with the apps output for apps", func() {
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(
				v2apps, nil)
			appsJSON, err := Curl("/v2/apps")
			Expect(err).To(BeNil())
			Expect(appsJSON).ToNot(BeNil())
		})

	})
})