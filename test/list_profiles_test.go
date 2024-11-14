package test

import (
	"context"
	"encoding/json"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/profiles"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/core-go"
)

func TestListProfiles(t *testing.T) {
	creds := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("EXCHANGE_CREDENTIALS")), creds); err != nil {
		t.Fatalf("unable to deserialize exchange credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		t.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(creds, httpClient)
	profilesSvc := profiles.NewProfilesService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := profilesSvc.ListProfiles(ctx, &profiles.ListProfilesRequest{Active: "true"})
	if err != nil {
		t.Fatalf("error listing profiles: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}

	if len(*response) == 0 {
		t.Fatal("expected at least one active profile in response")
	}

	firstProfile := (*response)[0]
	if firstProfile.Id == "" {
		t.Fatal("expected profile ID to be set")
	}
	if firstProfile.UserId == "" {
		t.Fatal("expected user ID to be set")
	}
	if firstProfile.Name == "" {
		t.Fatal("expected name to be set")
	}
	if !firstProfile.Active {
		t.Fatal("expected profile to be active")
	}
}
