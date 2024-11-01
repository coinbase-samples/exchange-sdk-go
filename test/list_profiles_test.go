package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/profiles"
)

func TestListProfiles(t *testing.T) {
	credentials, err := credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
	if err != nil {
		panic(fmt.Sprintf("unable to read exchange credentials: %v", err))
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)
	profilesSvc := profiles.NewProfilesService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := profilesSvc.ListProfiles(ctx, &profiles.ListProfilesRequest{})
	if err != nil {
		t.Fatalf("error listing profiles: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}
	if len(response.Profiles) == 0 {
		t.Fatal("expected at least one active profile in response")
	}

	firstProfile := response.Profiles[0]
	if firstProfile.Id == "" {
		t.Fatal("expected profile ID to be set")
	}
	if firstProfile.UserId == "" {
		t.Fatal("expected user ID to be set")
	}
	if firstProfile.Name == "" {
		t.Fatal("expected name to be set")
	}
}
