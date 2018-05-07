package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

// The actual Provider test itself
func TestMessageProvider_Success(t *testing.T) {
	pact := createPact()

	// Map test descriptions to message producer (handlers)
	// TODO: convert these all to types to ease readability
	functionMappings := dsl.MessageProviders{
		"some test case": func(m dsl.Message) (interface{}, error) {
			fmt.Println("Calling 'text' function that would produce a message")
			res := map[string]interface{}{
				"content": map[string]interface{}{
					"access": []map[string]string{
						{
							"role": "admin",
						},
						{
							"role": "admin",
						},
						{
							"role": "admin",
						},
					},
					"id":   27,
					"name": "Baz",
				},
			}
			return res, nil
		},
	}

	// Verify the Provider with local Pact Files
	pact.VerifyMessageProvider(t, types.VerifyMessageRequest{
		PactURLs: []string{filepath.ToSlash(fmt.Sprintf("%s/pactgomessageconsumer-pactgomessageprovider.json", pactDir))},
	}, functionMappings)
}

// Configuration / Test Data
var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)

// Setup the Pact client.
func createPact() dsl.Pact {
	// Create Pact connecting to local Daemon
	return dsl.Pact{
		Consumer:          "PactGoMessageConsumer",
		Provider:          "PactGoMessageProvider",
		LogDir:            logDir,
		LogLevel:          "WARN",
		PactFileWriteMode: "update",
	}
}