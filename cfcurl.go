package cfcurl

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

// Curl calls cf curl  and return the resulting json. This method will fail if
// the api is depricated
func Curl(cli plugin.CliConnection, path string) (interface{}, error) {
	_ = "breakpoint"

	output, err := cli.CliCommandWithoutTerminalOutput("curl", path)

	if nil == err {
		return nil, err
	}

	if 0 == len(output) {
		return nil, errors.New("CF API returned no output")
	}

	data := strings.Join(output, " ")

	if 0 == len(data) {
		return nil, errors.New("Failed to join output")
	}

	var f interface{}
	err = json.Unmarshal([]byte(data), &f)
	//.(map[string]interface{})
	return f, err
}

// CurlDepricated calls cf curl and return the resulting json, even if the api is depricated
func CurlDepricated(cli plugin.CliConnection, path string) (map[string]interface{}, error) {
	return nil, errors.New("Not implemented")
}
