package client

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/aquasecurity/trivy/pkg/log"
	"github.com/pact-foundation/pact-go/dsl"
	"gorm.io/gorm"
)

var pact dsl.Pact
var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)

type Artist struct {
	gorm.Model `swaggerignore:"true"` // using property gorm.Model automatically will be included: id, createdAt, updatedAt, deletedAt.
	Name       string                 `json:"name" validate:"required"`
	Album      string                 `json:"album"`
	Released   string                 `json:"released"`
	Country    string                 `json:"country"`
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer:                 "Sample consumer",
		Provider:                 "Sample provider",
		LogDir:                   logDir,
		PactDir:                  pactDir,
		DisableToolValidityCheck: true,
	}
}

func TestMain(m *testing.M) {
	pact = createPact()

	pact.Setup(true)

	code := m.Run()

	err := pact.WritePact()
	if err != nil {
		log.Infof("Failed to write your contract")
		return
	}

	pact.Teardown()
	os.Exit(code)
}

func getArtists() (err error) {
	url := fmt.Sprintf("http://localhost:%d/artist/specific?id=9", pact.Server.Port)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(req)
	//fmt.Println(req)

	return
}

func TestTheWholeBody_GET(t *testing.T) {
	pact.AddInteraction().
		Given("Match the response body").
		UponReceiving("A get request").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.Like("/artist/specific"),
			Query: dsl.MapMatcher{
				"id": dsl.Term("9", "[0-9]+"),
			},
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json(; charset=utf-8)?`),
			},
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body:   Artist{Name: "Racionais MC's", Album: "Sobrevivendo no Inferno", Released: "1997-12-20", Country: "Brazil"},
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json`),
			},
		})

	err := pact.Verify(getArtists)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)

	}

}

func TestInvalidRequest_GET(t *testing.T) {
	pact.
		AddInteraction().
		Given("Error message").
		UponReceiving("A GET invalid request").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.Like("/artist/specific"),
			Query: dsl.MapMatcher{
				"id": dsl.Term("9999", "[0-9]+"),
			},
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json`),
			},
		}).
		WillRespondWith(dsl.Response{
			Status: 404,
			Body: map[string]string{
				"message": "Requested artist is not found",
			},
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json`),
			},
		})

	err := pact.Verify(getArtists)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}
}
