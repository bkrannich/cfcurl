package cfcurl_test

import (
	"bufio"
	"errors"
	"os"

	. "github.com/krujos/cfcurl"

	"github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cfcurl", func() {
	var fakeCliConnection *fakes.FakeCliConnection
	var v2apps []string

	Describe("an api that is not depricated", func() {
		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			file, err := os.Open("apps.json")
			defer file.Close()
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

			if 0 == len(v2apps) {
				Fail("you didn't read anything in")
			}
		})

		AfterEach(func() {
			v2apps = nil
		})

		Describe("cf cli results validation", func() {
			It("returns an error when there is no output", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputReturns(nil, nil)
				appsJSON, err := Curl(fakeCliConnection, "/v2/apps")
				Expect(err).ToNot(BeNil())
				Expect(appsJSON).To(BeNil())
			})

			It("returns an error with zero length output", func() {

				fakeCliConnection.CliCommandWithoutTerminalOutputReturns([]string{""}, nil)
				appsJSON, err := Curl(fakeCliConnection, "/v2/apps")
				Expect(err).ToNot(BeNil())
				Expect(appsJSON).To(BeNil())
			})

			It("should call the path specified", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputReturns(v2apps, nil)
				Curl(fakeCliConnection, "/v2/an_unpredictable_path")
				args := fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)
				Expect("curl").To(Equal(args[0]))
				Expect("/v2/an_unpredictable_path").To(Equal(args[1]))
			})

			It("returns an error when the cli fails", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputReturns(nil, errors.New("Something bad"))
				appsJSON, err := Curl(fakeCliConnection, "/v2/an_unpredictable_path")
				Expect(appsJSON).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})
		It("should return the output for apps", func() {
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(v2apps, nil)
			appsJSON, err := Curl(fakeCliConnection, "/v2/apps")
			Expect(err).To(BeNil())
			Expect(appsJSON).ToNot(BeNil())
		})
	})
})
