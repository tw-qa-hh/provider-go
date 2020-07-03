package main

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Provider:                 "provider-go",
		DisableToolValidityCheck: true,
		LogLevel:                 "DEBUG",
	}

	//consumer := types.ConsumerVersionSelector{
	//	Latest: true,
	//}

	// Start provider API in the background
	go main()

	pact.VerifyProvider(t, types.VerifyRequest{
		Provider:                   "provider-go",
		ProviderBaseURL:            "http://localhost:8080",
		BrokerURL:                  "https://qa-ham-pact-broker.herokuapp.com/",
		//ConsumerVersionSelectors:   []types.ConsumerVersionSelector{consumer},
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		FailIfNoPactsFound:         false,
		EnablePending:              true,
	})
}
