package provider

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/lucasjct/database"
	"github.com/lucasjct/routes"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func startProvider() {
	routes.HandleRequest()
	database.DataBaseConnect()
}

func TestServerPactVerification(t *testing.T) {

	go startProvider()

	time.Sleep(2 * time.Second)

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../../client/pacts", dir)
	var logDir = fmt.Sprintf("%s/../../client/log", dir)

	pact := &dsl.Pact{
		Provider:                 "Sample Provider",
		LogDir:                   logDir,
		PactDir:                  pactDir,
		DisableToolValidityCheck: true,
		LogLevel:                 "DEBUG",
	}
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"../client/pacts/sample_consumer-sample_provider.json"},
	})

	if err != nil {
		t.Fatal(err)
	}

}
