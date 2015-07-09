package cfcurl

import (
	"errors"

	"github.com/cloudfoundry/cli/plugin"
)

// Curl calls cf curl  and return the resulting json. This method will fail if
// the api is depricated
func Curl(cli plugin.CliConnection, path string) (map[string]interface{}, error) {
	_, err := cli.CliCommandWithoutTerminalOutput("curl", path)
	return nil, err
}

// CurlDepricated calls cf curl and return the resulting json, even if the api is depricated
func CurlDepricated(cli plugin.CliConnection, path string) (map[string]interface{}, error) {
	return nil, errors.New("Not implemented")
}
